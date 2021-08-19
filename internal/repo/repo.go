package repo

import "ova-recipe-api/internal/recipe"

type RecipeRepo interface {
	AddRecipes(recipes []recipe.Recipe) error
	ListRecipes(limit, offset uint64) ([]recipe.Recipe, error)
	DescribeRecipe(recipeId uint64) (*recipe.Recipe, error)
}
