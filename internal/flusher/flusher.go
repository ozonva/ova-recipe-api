package flusher

import (
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/repo"
)

type Flusher interface {
	Flush(recipes []recipe.Recipe) []recipe.Recipe
}

func NewFlusher(chunkSize uint64, recipeRepo repo.RecipeRepo) Flusher {
	return &flusher{chunkSize: chunkSize, recipeRepo: recipeRepo}
}

type flusher struct {
	chunkSize  uint64
	recipeRepo repo.RecipeRepo
}

func (f *flusher) Flush(recipes []recipe.Recipe) []recipe.Recipe {
	return make([]recipe.Recipe, 0)
}
