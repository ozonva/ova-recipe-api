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
			recipe.New(1, 1, "", "", []recipe.Action{}),
			recipe.New(2, 2, "", "", []recipe.Action{}),
			recipe.New(3, 3, "", "", []recipe.Action{}),
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
			Context("recipes count less than chunkSize", func() {
				It("Flush should return nil", func() {
					oneRecipe := recipes[:1]
					mockRepo.EXPECT().AddRecipes(oneRecipe).Return(nil).Times(1)
					Expect(testFlusher.Flush(oneRecipe)).To(BeNil())
				})
			})
			Context("recipes count more than chunkSize", func() {
				It("Flush should return nil", func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddRecipes(recipes[:2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddRecipes(recipes[2:]).Return(nil).Times(1),
					)
					Expect(testFlusher.Flush(recipes)).To(BeNil())
				})
			})
		})
		When("can not write", func() {
			err := fmt.Errorf("can not wrire data")
			Context("all data", func() {
				It("Flush should return all data", func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddRecipes(recipes[:2]).Return(err).Times(1),
						mockRepo.EXPECT().AddRecipes(recipes[2:]).Return(err).Times(1),
					)
					Expect(testFlusher.Flush(recipes)).To(Equal(recipes))
				})
			})
			Context("some data", func() {
				It("Flush should return only rejected data", func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddRecipes(recipes[:2]).Return(err).Times(1),
						mockRepo.EXPECT().AddRecipes(recipes[2:]).Return(nil).Times(1),
					)
					Expect(testFlusher.Flush(recipes)).To(Equal(recipes[:2]))
				})
			})
		})
	})
})
