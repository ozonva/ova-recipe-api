package saver

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-recipe-api/internal/flusher"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/ticker"
	"time"
)

func runGoroutines(funcs []func()) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		for idx := range funcs {
			funcs[idx]()
			time.Sleep(100 * time.Millisecond)
			ch <- struct{}{}
		}
	}()
	return ch
}

var _ = Describe("Saver", func() {
	var (
		mockCtrl    *gomock.Controller
		mockFlusher *flusher.MockFlusher
		mockTicker  *ticker.MockTicker
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = flusher.NewMockFlusher(mockCtrl)
		mockTicker = ticker.NewMockTicker(mockCtrl)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Saver test", func() {
		When("saver does not run", func() {
			Context("saver capacity == 0", func() {
				saver := saver{flusher: mockFlusher, recipes: make([]recipe.Recipe, 0), cancelFn: func() {}}
				It("Save returns error at first Save call", func() {
					err := saver.Save(recipe.New(1, 1, "", "", []recipe.Action{}))
					Expect(err).To(Equal(NotEnoughCapacityError))
				})
			})
			Context("saver capacity > 0", func() {
				capacity := uint64(2)
				saver := saver{flusher: mockFlusher, recipes: make([]recipe.Recipe, 0, capacity), cancelFn: func() {}}
				It("returns error after 3rd Save call", func() {
					for idx := uint64(0); idx < capacity; idx++ {
						err := saver.Save(recipe.New(idx, idx, "", "", []recipe.Action{}))
						Expect(err).To(BeNil())
					}
					err := saver.Save(recipe.New(1, 1, "", "", []recipe.Action{}))
					Expect(err).To(Equal(NotEnoughCapacityError))
					Expect(saver.recipes).To(Equal([]recipe.Recipe{
						recipe.New(0, 0, "", "", []recipe.Action{}),
						recipe.New(1, 1, "", "", []recipe.Action{}),
					}))
				})
			})
		})
		When("server runs", func() {
			Context("Empty saver, tick event than cancel", func() {
				It("test", func() {
					tickerCh := make(chan time.Time)
					saver := saver{flusher: mockFlusher, recipes: make([]recipe.Recipe, 5), cancelFn: func() {}}
					mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
					mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
					gomock.InOrder(
						mockFlusher.EXPECT().Flush(make([]recipe.Recipe, 5)).Return(nil).Times(1),
						mockFlusher.EXPECT().Flush(make([]recipe.Recipe, 0)).Return(nil).Times(1),
					)
					ch := runGoroutines(
						[]func(){
							func() { saver.Run(mockTicker) },
							func() { tickerCh <- time.Time{} },
							func() { saver.Close() },
						})
					<-ch // saver started
					<-ch // set ticker closed
					<-ch // saver closed
				})
			})
		})
	})
})
