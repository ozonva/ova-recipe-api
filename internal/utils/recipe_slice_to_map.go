package utils

import (
	"fmt"
	"ova-recipe-api/internal/recipe"
)

func RecipeSliceToMap(recipes []recipe.Recipe) (map[uint64]recipe.Recipe, error) {
	result := make(map[uint64]recipe.Recipe, len(recipes))
	for idx := range recipes {
		recipeId := recipes[idx].Id()
		if _, ok := result[recipeId]; ok {
			return nil, fmt.Errorf("duplicate recipe id '%d'", recipeId)
		}
		result[recipeId] = recipes[idx]
	}
	return result, nil
}
