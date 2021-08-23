package saver

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-recipe-api/internal/flusher"
	"ova-recipe-api/internal/recipe"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl    *gomock.Controller
		mockFlusher *flusher.MockFlusher
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = flusher.NewMockFlusher (mockCtrl)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Save test", func() {
		When("saver does not run", func() {
			When("saver capacity == 0", func() {
				saver := saver{flusher: mockFlusher, recipes: make([]recipe.Recipe, 0)}
				It("returns error at first Save call", func() {
					err := saver.Save(recipe.New(1, 1, "", "", []recipe.Action{}))
					Expect(err).To(Equal(NotEnoughCapacityError))
				})
			})
			When("saver capacity == 2", func() {
				saver := saver{flusher: mockFlusher, recipes: make([]recipe.Recipe, 0, 2)}
				It("returns error after 3rd Save call", func() {
					for idx := 0; idx < 2; idx++ {
						err := saver.Save(recipe.New(1, 1, "", "", []recipe.Action{}))
						Expect(err).To(BeNil())
					}
					err := saver.Save(recipe.New(1, 1, "", "", []recipe.Action{}))
					Expect(err).To(Equal(NotEnoughCapacityError))
				})
			})
		})
	})
})
