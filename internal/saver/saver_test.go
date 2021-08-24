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
		Describe("New function", func() {
			When("capacity == 0", func() {
				It("should return error", func() {
					_, err := New(mockFlusher, 0, mockTicker)
					Expect(err).To(Equal(ZeroCapacityError))
				})
			})
		})
		When("capacity > 0", func() {
			var tickerCh chan time.Time
			BeforeEach(func() {
				tickerCh = make(chan time.Time)
				mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
				mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
			})
			Context("one tick w/o values", func() {
				It("flush should not be called", func() {
					mockFlusher.EXPECT().Flush(nil).Return(nil).Times(0)
					s, err := New(mockFlusher, 1, mockTicker)
					Expect(err).To(BeNil())
					tickerCh <- time.Time{}
					s.Close()
				})
			})
			Context("one tick with values < cap", func() {
				It("flush should be called one time", func() {
					mockFlusher.EXPECT().Flush([]recipe.Recipe{{}, {}}).Return(nil).Times(1)
					s, _ := New(mockFlusher, 5, mockTicker)
					s.Save(recipe.Recipe{})
					s.Save(recipe.Recipe{})
					tickerCh <- time.Time{}
					s.Close()
				})
			})
			Context("w/o tick values > cap", func() {
				It("flush should be called several times", func() {
					gomock.InOrder(
						mockFlusher.EXPECT().Flush([]recipe.Recipe{{}, {}}).Return(nil).Times(1),
						mockFlusher.EXPECT().Flush([]recipe.Recipe{{}}).Return(nil).Times(1),
					)
					s, _ := New(mockFlusher, 2, mockTicker)
					s.Save(recipe.Recipe{})
					s.Save(recipe.Recipe{})
					s.Save(recipe.Recipe{})
					s.Close()
				})
			})
			Context("close w/o values", func() {
				It("should be ok", func() {
					s, _ := New(mockFlusher, 1, mockTicker)
					s.Close()
				})
			})
		})
	})
})
