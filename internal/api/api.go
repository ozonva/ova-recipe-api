package api

import (
	"ova-recipe-api/internal/repo"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

type GRPCServer struct {
	recipeApi.OvaRecipeApiServer
	recipeRepo repo.RecipeRepo
	metrics    Metrics
}

func NewOvaRecipeApiServer(recipeRepo repo.RecipeRepo) recipeApi.OvaRecipeApiServer {
	return &GRPCServer{recipeRepo: recipeRepo, metrics: newApiMetrics()}
}
