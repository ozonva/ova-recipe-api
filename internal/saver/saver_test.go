package saver

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-recipe-api/internal/flusher"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/ticker"
	"sync"
	"sync/atomic"
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
		Describe("Close function", func() {
			var tickerCh chan time.Time
			BeforeEach(func() {
				tickerCh = make(chan time.Time)
				mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
				mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
			})
			When("call Close w/o Save", func() {
				It("returns nil", func() {
					mockFlusher.EXPECT().Flush(nil).Return(nil).Times(0)
					s, err := New(mockFlusher, 1, mockTicker)
					Expect(err).To(BeNil())
					s.Close()
					s.Close()
				})
			})
		})
		Describe("Save function", func() {
			When("call Save after Close", func() {
				BeforeEach(func() {
					tickerCh := make(chan time.Time)
					mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
					mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
				})
				It("returns error", func() {
					mockFlusher.EXPECT().Flush(nil).Return(nil).Times(0)
					s, _ := New(mockFlusher, 1, mockTicker)
					s.Close()
					Expect(s.Save(recipe.Recipe{})).To(Equal(SaveAfterCloseError))
				})
			})
			When("call Save more than capacity times", func() {
				BeforeEach(func() {
					tickerCh := make(chan time.Time)
					mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
					mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
				})
				It("returns error", func() {
					mockFlusher.EXPECT().Flush([]recipe.Recipe{{}}).Return(nil).Times(1)
					s, _ := New(mockFlusher, 1, mockTicker)
					Expect(s.Save(recipe.Recipe{})).To(BeNil())
					Expect(s.Save(recipe.Recipe{})).To(Equal(NotEnoughCapacityError))
					s.Close()
				})
			})
			When("call Save more than capacity times with tick between", func() {
				var tickerCh chan time.Time
				BeforeEach(func() {
					tickerCh = make(chan time.Time)
					mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
					mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
				})
				It("returns error", func() {
					mockFlusher.EXPECT().Flush([]recipe.Recipe{{}}).Return(nil).Times(2)
					s, _ := New(mockFlusher, 1, mockTicker)
					Expect(s.Save(recipe.Recipe{})).To(BeNil())
					tickerCh <- time.Time{}
					Expect(s.Save(recipe.Recipe{})).To(BeNil())
					s.Close()
				})
			})
			When("Save calls from a lot of coroutines", func() {
				BeforeEach(func() {
					tickerCh := make(chan time.Time)
					mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
					mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
				})
				It("should be ok", func() {
					capacity := 1000
					mockFlusher.EXPECT().Flush(gomock.Any()).DoAndReturn(func(r []recipe.Recipe) {
						Expect(len(r)).To(Equal(capacity))
					}).Return(nil).Times(1)
					coroutinesCount := 10
					var startWG sync.WaitGroup
					startWG.Add(coroutinesCount)
					var readyWG sync.WaitGroup
					readyWG.Add(coroutinesCount)

					s, _ := New(mockFlusher, uint(capacity), mockTicker)

					goroFunc := func(startIdx uint64, count uint64) {
						defer readyWG.Done()
						startWG.Done()
						startWG.Wait()
						for idx := startIdx; idx < startIdx+count; idx++ {
							Expect(s.Save(recipe.New(idx, idx, "", "", []string{}))).To(BeNil())
						}
					}
					for idx := 0; idx < coroutinesCount; idx += 1 {
						go goroFunc(uint64(idx*100), 100)
					}

					readyWG.Wait()

					s.Close()
				})
			})
			When("Save calls from a lot of coroutines with ticks", func() {
				var tickerCh chan time.Time
				BeforeEach(func() {
					tickerCh = make(chan time.Time)
					mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
					mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
				})
				It("should be ok", func() {
					totalFlushedRecipes := uint64(0)
					mockFlusher.EXPECT().Flush(gomock.Any()).DoAndReturn(func(r []recipe.Recipe) {
						totalFlushedRecipes += uint64(len(r))
					}).Return(nil).AnyTimes()

					coroutinesCount := 11
					var startWG sync.WaitGroup
					startWG.Add(coroutinesCount)
					var readyWG sync.WaitGroup
					readyWG.Add(coroutinesCount)

					capacity := 1000
					s, _ := New(mockFlusher, uint(capacity), mockTicker)

					goroSaveFunc := func(startIdx uint64, count uint64) {
						defer readyWG.Done()
						startWG.Done()
						startWG.Wait()
						for idx := startIdx; idx < startIdx+count; idx++ {
							Expect(s.Save(recipe.New(idx, idx, "", "", []string{}))).To(BeNil())
						}
					}

					for idx := 0; idx < coroutinesCount-1; idx += 1 {
						go goroSaveFunc(uint64(idx*100), 100)
					}

					go func() {
						defer readyWG.Done()
						startWG.Done()
						startWG.Wait()
						for idx := 0; idx < 10; idx += 1 {
							tickerCh <- time.Time{}
							time.Sleep(100 * time.Millisecond)
						}
					}()

					readyWG.Wait()
					s.Close()
					Expect(totalFlushedRecipes).To(Equal(uint64(capacity)))
				})
			})
			When("Save calls from a lot of coroutines with ticks and small capacity", func() {
				var tickerCh chan time.Time
				BeforeEach(func() {
					tickerCh = make(chan time.Time)
					mockTicker.EXPECT().Chanel().Return(tickerCh).Times(1)
					mockTicker.EXPECT().Stop().Do(func() { close(tickerCh) }).Return().Times(1)
				})
				It("should be ok", func() {
					totalFlushedRecipes := uint64(0)
					mockFlusher.EXPECT().Flush(gomock.Any()).DoAndReturn(func(r []recipe.Recipe) {
						totalFlushedRecipes += uint64(len(r))
					}).Return(nil).AnyTimes()

					coroutinesCount := 11
					var startWG sync.WaitGroup
					startWG.Add(coroutinesCount)
					var readyWG sync.WaitGroup
					readyWG.Add(coroutinesCount)

					s, _ := New(mockFlusher, 1000, mockTicker)

					totalRejectedRecipes := uint64(0)
					totalAttempts := uint64(0)
					goroSaveFunc := func(startIdx uint64, count uint64) {
						defer readyWG.Done()
						startWG.Done()
						startWG.Wait()
						for idx := startIdx; idx < startIdx+count; idx++ {
							atomic.AddUint64(&totalAttempts, 1)
							if err := s.Save(recipe.New(idx, idx, "", "", []string{})); err != nil {
								atomic.AddUint64(&totalRejectedRecipes, 1)
							}
						}
					}

					for idx := 0; idx < coroutinesCount-1; idx += 1 {
						go goroSaveFunc(uint64(idx*100), 100)
					}

					go func() {
						defer readyWG.Done()
						startWG.Done()
						startWG.Wait()
						for idx := 0; idx < 10; idx += 1 {
							tickerCh <- time.Time{}
							time.Sleep(100 * time.Millisecond)
						}
					}()

					readyWG.Wait()
					s.Close()
					Expect(totalFlushedRecipes + totalRejectedRecipes).To(Equal(totalAttempts))
				})
			})
		})
	})
})
