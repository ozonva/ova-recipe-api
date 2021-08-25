package saver

import (
	"context"
	"ova-recipe-api/internal/flusher"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/ticker"
	"sync"
)

type Error string

func (e Error) Error() string { return string(e) }

const (
	ZeroCapacityError      = Error("Capacity should be more than 0. ")
	SaveAfterCloseError    = Error("Save error, saver already closed. ")
	NotEnoughCapacityError = Error("Cannot save new recipe, not enough capacity. ")
)

type Saver interface {
	Save(recipe recipe.Recipe) error
	Close()
}

func New(flusher flusher.Flusher, capacity uint, ticker ticker.Ticker) (Saver, error) {
	if capacity == 0 {
		return nil, ZeroCapacityError
	}
	startCloseCtx, startCloseOp := context.WithCancel(context.Background())
	finishCloseCtx, finishCancelOp := context.WithCancel(context.Background())

	s := &saver{
		flusher:        flusher,
		recipesBuf:     make([]recipe.Recipe, 0, capacity),
		startCloseOp:   startCloseOp,
		finishCloseCtx: finishCloseCtx,
	}
	s.run(ticker, startCloseCtx, finishCancelOp)
	return s, nil
}

type saver struct {
	flusher        flusher.Flusher
	recipesGuard   sync.Mutex
	recipesBuf     []recipe.Recipe
	startCloseOp   context.CancelFunc
	finishCloseCtx context.Context
}

func (s *saver) run(ticker ticker.Ticker, startCloseCtx context.Context, finishCloseOp context.CancelFunc) {
	go func() {
		tickerCh := ticker.Chanel()
		defer ticker.Stop()
		for {
			select {
			case <-startCloseCtx.Done():
				s.flushWithClose()
				finishCloseOp()
				return
			case <-tickerCh:
				s.flush(s.cloneRecipes())
			}
		}
	}()
}

func (s *saver) cloneRecipes() []recipe.Recipe {
	s.recipesGuard.Lock()
	defer s.recipesGuard.Unlock()
	if len(s.recipesBuf) > 0 {
		recipes := s.recipesBuf
		s.recipesBuf = make([]recipe.Recipe, 0, cap(s.recipesBuf))
		return recipes
	}
	return nil
}

func (s *saver) flush(recipes []recipe.Recipe) {
	if recipes != nil && len(recipes) > 0 {
		s.flusher.Flush(recipes)
	}
}

func (s *saver) flushWithClose() {
	s.recipesGuard.Lock()
	defer s.recipesGuard.Unlock()
	s.flush(s.recipesBuf)
	s.recipesBuf = nil
}

func (s *saver) Save(recipe recipe.Recipe) error {
	s.recipesGuard.Lock()
	defer s.recipesGuard.Unlock()
	if s.recipesBuf == nil {
		return SaveAfterCloseError
	}
	if len(s.recipesBuf) == cap(s.recipesBuf) {
		return NotEnoughCapacityError
	}
	s.recipesBuf = append(s.recipesBuf, recipe)
	return nil
}

// Close stops saver and wait flushing last data, it is blocking call.
// After this call any Save calls will return error SaveAfterCloseError
func (s *saver) Close() {
	s.startCloseOp()
	<-s.finishCloseCtx.Done()
}
