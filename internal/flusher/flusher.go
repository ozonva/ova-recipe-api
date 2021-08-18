package flusher

import (
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/repo"
	"ova-recipe-api/internal/utils"
)

type Flusher interface {
	Flush(recipes []recipe.Recipe) []recipe.Recipe
}

func NewFlusher(chunkSize uint, recipeRepo repo.RecipeRepo) Flusher {
	return &flusher{chunkSize: chunkSize, recipeRepo: recipeRepo}
}

type flusher struct {
	chunkSize  uint
	recipeRepo repo.RecipeRepo
}

func (f *flusher) Flush(recipes []recipe.Recipe) []recipe.Recipe {
	var result []recipe.Recipe
	for _, recipesChunk := range utils.SplitRecipeSlice(recipes, f.chunkSize) {
		if err := f.recipeRepo.AddRecipes(recipesChunk); err != nil {
			result = append(result, recipesChunk...)
		}
	}
	return result
}
