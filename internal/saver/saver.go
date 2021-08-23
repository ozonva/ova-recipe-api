package saver

import (
	"context"
	"ova-recipe-api/internal/flusher"
	"ova-recipe-api/internal/recipe"
	"sync"
	"time"
)

type Error string

func (e Error) Error() string { return string(e) }

const NotEnoughCapacityError = Error("Cannot save new recipe, not enough capacity. ")

type Saver interface {
	Run(saveInterval time.Duration)
	Save(recipe recipe.Recipe) error
	Close()
}

func New(flusher flusher.Flusher, capacity uint) Saver {
	return &saver{flusher: flusher, recipes: make([]recipe.Recipe, 0, capacity)}
}

type saver struct {
	flusher      flusher.Flusher
	recipesGuard sync.Mutex
	recipes      []recipe.Recipe
	cancelFn     context.CancelFunc
}

func (s *saver) Run(saveInterval time.Duration) {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancelFn = cancel
	go func() {
		ticker := time.NewTicker(saveInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				s.flush()
				return
			case <-ticker.C:
				s.flush()
			}
		}
	}()
}

func (s *saver) cloneRecipes() []recipe.Recipe {
	s.recipesGuard.Lock()
	defer s.recipesGuard.Unlock()
	if 0 == len(s.recipes) {
		return make([]recipe.Recipe, 0)
	}
	clone := s.recipes
	s.recipes = make([]recipe.Recipe, 0, cap(s.recipes))
	return clone
}

func (s *saver) flush() {
	s.flusher.Flush(s.cloneRecipes())
}

func (s *saver) Save(recipe recipe.Recipe) error {
	s.recipesGuard.Lock()
	defer s.recipesGuard.Unlock()
	if cap(s.recipes) == len(s.recipes) {
		return NotEnoughCapacityError
	}
	s.recipes = append(s.recipes, recipe)
	return nil
}

func (s *saver) Close() {
	s.cancelFn()
}
