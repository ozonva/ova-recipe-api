package api

import recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"

type GRPCServer struct {
	recipeApi.OvaRecipeApiServer
}

func NewOvaRecipeApiServer() recipeApi.OvaRecipeApiServer {
	return &GRPCServer{}
}
