package api_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"ova-recipe-api/internal/api"
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
		grpcServer  *grpc.Server
		ctx         context.Context
		startWG     sync.WaitGroup
		connect     *grpc.ClientConn
		client      recipeApi.OvaRecipeApiClient
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRecipeRepo(mockCtrl)
		grpcServer = grpc.NewServer()
		recipeApi.RegisterOvaRecipeApiServer(grpcServer, api.NewOvaRecipeApiServer(mockRepo))
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
				mockRepo.EXPECT().AddRecipe(gomock.Any(), expectedRecipe).Return(expectedRecipeId, nil).Times(1)
			})
			It("should return new recipe id", func() {
				req := recipeApi.CreateRecipeRequestV1{
					UserId: expectedRecipe.UserId(),
					Name: expectedRecipe.Name(),
					Description: expectedRecipe.Description(),
					Actions: expectedRecipe.Actions(),
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
				mockRepo.EXPECT().AddRecipe(gomock.Any(), expectedRecipe).Return(
					uint64(0), expectedError).Times(1)
			})
			It("should return error", func() {
				req := recipeApi.CreateRecipeRequestV1{
					UserId: expectedRecipe.UserId(),
					Name: expectedRecipe.Name(),
					Description: expectedRecipe.Description(),
					Actions: expectedRecipe.Actions(),
				}
				newRecipeResponse, err := client.CreateRecipeV1(ctx, &req)
				Expect(err.Error()).To(ContainSubstring(expectedError.Error()))
				Expect(newRecipeResponse).To(BeNil())
			})
		})
	})
})
