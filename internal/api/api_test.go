package api

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/repo"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
	"sync"
)

var _ = Describe("Api", func() {
	const bufSize = 1024 * 1024
	var (
		bufListener *bufconn.Listener
		mockCtrl    *gomock.Controller
		mockRepo    *repo.MockRecipeRepo
		mockMetrics *MockMetrics
		grpcServer  *grpc.Server
		ctx         context.Context
		startWG     sync.WaitGroup
		connect     *grpc.ClientConn
		client      recipeApi.OvaRecipeApiClient
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRecipeRepo(mockCtrl)
		mockMetrics = NewMockMetrics(mockCtrl)
		grpcServer = grpc.NewServer()
		recipeApi.RegisterOvaRecipeApiServer(grpcServer, &GRPCServer{recipeRepo: mockRepo, metrics: mockMetrics})
		bufListener = bufconn.Listen(bufSize)
		startWG.Add(1)
		go func() {
			startWG.Done()
			Expect(grpcServer.Serve(bufListener)).To(BeNil())
		}()
		startWG.Wait()
		ctx = context.Background()
		conn, err := grpc.DialContext(
			ctx, "bufnet",
			grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
				return bufListener.Dial()
			}), grpc.WithInsecure())
		Expect(err).To(BeNil())
		connect = conn
		client = recipeApi.NewOvaRecipeApiClient(conn)
	})
	AfterEach(func() {
		mockCtrl.Finish()
		connect.Close()
		grpcServer.Stop()
	})
	Describe("Do Create queries", func() {
		When("Can create new recipe", func() {
			expectedRecipe := recipe.New(0, 1, "test name", "test description", []string{"testOne", "testTwo"})
			expectedRecipeId := uint64(1)
			BeforeEach(func() {
				mockMetrics.EXPECT().incSuccessCreateRecipeCounter().Times(1)
				mockRepo.EXPECT().AddRecipe(gomock.Any(), expectedRecipe).Return(expectedRecipeId, nil).Times(1)
			})
			It("should return new recipe id", func() {
				req := recipeApi.CreateRecipeRequestV1{
					UserId:      expectedRecipe.UserId(),
					Name:        expectedRecipe.Name(),
					Description: expectedRecipe.Description(),
					Actions:     expectedRecipe.Actions(),
				}
				newRecipeResponse, err := client.CreateRecipeV1(ctx, &req)
				Expect(err).To(BeNil())
				Expect(newRecipeResponse.RecipeId).To(Equal(expectedRecipeId))
			})
		})
		When("Can not create new recipe", func() {
			expectedRecipe := recipe.New(0, 1, "test name", "test description", []string{"testOne", "testTwo"})
			expectedError := fmt.Errorf("Can not create new recipe. ")
			BeforeEach(func() {
				mockMetrics.EXPECT().incSuccessCreateRecipeCounter().Times(0)
				mockRepo.EXPECT().AddRecipe(gomock.Any(), expectedRecipe).Return(
					uint64(0), expectedError).Times(1)
			})
			It("should return error", func() {
				req := recipeApi.CreateRecipeRequestV1{
					UserId:      expectedRecipe.UserId(),
					Name:        expectedRecipe.Name(),
					Description: expectedRecipe.Description(),
					Actions:     expectedRecipe.Actions(),
				}
				newRecipeResponse, err := client.CreateRecipeV1(ctx, &req)
				Expect(err.Error()).To(ContainSubstring(expectedError.Error()))
				Expect(newRecipeResponse).To(BeNil())
			})
		})
		When("Invalid userId", func() {
			BeforeEach(func() {
				mockMetrics.EXPECT().incSuccessCreateRecipeCounter().Times(0)
				mockRepo.EXPECT().AddRecipe(gomock.Any(), gomock.Any()).Times(0)
			})
			It("should return error", func() {
				req := recipeApi.CreateRecipeRequestV1{
					UserId:      0, // invalid id
					Name:        "test name",
					Description: "test description",
					Actions:     []string{"testOne", "testTwo"},
				}
				newRecipeResponse, err := client.CreateRecipeV1(ctx, &req)
				Expect(err.Error()).To(ContainSubstring("invalid CreateRecipeRequestV1.UserId: value must be greater than 0"))
				Expect(newRecipeResponse).To(BeNil())
			})
		})
	})
	Describe("Do multi create query", func() {
		When("recipes count > batch size", func() {
			expectedRecipes := []recipe.Recipe{
				recipe.New(0, 1, "test name", "test description", []string{"testOne", "testTwo"}),
				recipe.New(0, 2, "test name", "test description", []string{"testOne", "testTwo"}),
				recipe.New(0, 3, "test name", "test description", []string{"testOne", "testTwo"}),
			}
			BeforeEach(func() {
				gomock.InOrder(
					mockRepo.EXPECT().AddRecipes(gomock.Any(), expectedRecipes[:2]).Return(nil).Times(1),
					mockRepo.EXPECT().AddRecipes(gomock.Any(), expectedRecipes[2:]).Return(nil).Times(1),
				)
			})
			It("should return new recipe id", func() {
				recipes := make([]*recipeApi.CreateRecipeV1, 0, len(expectedRecipes))
				for _, expectedRecipe := range expectedRecipes {
					recipes = append(recipes, &recipeApi.CreateRecipeV1{
						UserId:      expectedRecipe.UserId(),
						Name:        expectedRecipe.Name(),
						Description: expectedRecipe.Description(),
						Actions:     expectedRecipe.Actions(),
					})
				}
				req := recipeApi.MultiCreateRecipeRequestV1{
					Recipes: recipes,
				}
				_, err := client.MultiCreateRecipeV1(ctx, &req)
				Expect(err).To(BeNil())
			})
		})
	})
})
