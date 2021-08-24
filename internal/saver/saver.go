package saver

import (
	"context"
	"ova-recipe-api/internal/flusher"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/ticker"
)

type Error string

func (e Error) Error() string { return string(e) }

const ZeroCapacityError = Error("Capacity should be more than 0. ")

type Saver interface {
	Save(recipe recipe.Recipe)
	Close()
}

func New(flusher flusher.Flusher, capacity uint, ticker ticker.Ticker) (Saver, error) {
	if capacity == 0 {
		return nil, ZeroCapacityError
	}
	finishContext, cancel := context.WithCancel(context.Background())
	s := &saver{
		flusher:       flusher,
		recipesBuf:    make([]recipe.Recipe, 0, capacity),
		recipesCh:     make(chan recipe.Recipe),
		finishContext: finishContext,
	}
	s.run(ticker, cancel)
	return s, nil
}

type saver struct {
	flusher       flusher.Flusher
	recipesBuf    []recipe.Recipe
	recipesCh     chan recipe.Recipe
	finishContext context.Context
}

func (s *saver) run(ticker ticker.Ticker, cancel context.CancelFunc) {
	go func() {
		tickerCh := ticker.Chanel()
		defer ticker.Stop()
		for {
			select {
			case r, isOpen := <-s.recipesCh:
				if !isOpen {
					s.flush()
					cancel()
					return
				}
				if len(s.recipesBuf) == cap(s.recipesBuf) {
					s.flush()
				}
				s.recipesBuf = append(s.recipesBuf, r)
			case <-tickerCh:
				s.flush()
			}
		}
	}()
}

func (s *saver) flush() {
	if len(s.recipesBuf) > 0 {
		s.flusher.Flush(s.recipesBuf)
		s.recipesBuf = s.recipesBuf[:0]
	}
}

func (s *saver) Save(recipe recipe.Recipe) {
	s.recipesCh <- recipe
}

// Close stops saver and wait flushing last data
// Could be called only once
// It is blocking call
func (s *saver) Close() {
	close(s.recipesCh)
	<-s.finishContext.Done()
}
