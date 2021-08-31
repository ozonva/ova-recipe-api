package flusher_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-recipe-api/internal/flusher"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/repo"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *repo.MockRecipeRepo
		testFlusher flusher.Flusher
		recipes     = []recipe.Recipe{
			recipe.New(1, 1, "", "", []string{}),
			recipe.New(2, 2, "", "", []string{}),
			recipe.New(3, 3, "", "", []string{}),
		}
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = repo.NewMockRecipeRepo(mockCtrl)
		testFlusher = flusher.NewFlusher(uint(len(recipes)-1), mockRepo)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Writing data to DB", func() {
		When("can write all data", func() {
			AssertFlushReturnNil := func(inRecipes []recipe.Recipe) {
				Expect(testFlusher.Flush(inRecipes)).To(BeNil())
			}
			Context("recipes count less than chunkSize", func() {
				oneRecipe := recipes[:1]
				BeforeEach(func() {
					mockRepo.EXPECT().AddRecipes(oneRecipe).Return(nil).Times(1)
				})
				It("Flush should return nil", func() {
					AssertFlushReturnNil(oneRecipe)
				})
			})
			Context("recipes count more than chunkSize", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddRecipes(recipes[:2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddRecipes(recipes[2:]).Return(nil).Times(1),
					)
				})
				It("Flush should return nil", func() {
					AssertFlushReturnNil(recipes)
				})
			})
		})
		When("can not write", func() {
			err := fmt.Errorf("can not wrire data")
			AssertFlushReturnSomeRecipes := func(outRecipes []recipe.Recipe) {
				Expect(testFlusher.Flush(recipes)).To(Equal(outRecipes))
			}
			Context("all data", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddRecipes(recipes[:2]).Return(err).Times(1),
						mockRepo.EXPECT().AddRecipes(recipes[2:]).Return(err).Times(1),
					)
				})
				It("Flush should return all data", func() {
					AssertFlushReturnSomeRecipes(recipes)
				})
			})
			Context("some data", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddRecipes(recipes[:2]).Return(err).Times(1),
						mockRepo.EXPECT().AddRecipes(recipes[2:]).Return(nil).Times(1),
					)
				})
				It("Flush should return only rejected data", func() {
					AssertFlushReturnSomeRecipes(recipes[:2])
				})
			})
		})
	})
})
