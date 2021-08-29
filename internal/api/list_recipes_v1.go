package api

import (
	"context"
	"github.com/rs/zerolog/log"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) ListRecipesV1(ctx context.Context, req *recipeApi.ListRecipesRequestV1) (*recipeApi.ListRecipesResponseV1, error) {
	log.Info().Msgf("Receive new list request: %s", req.String())
	dbRecipes, err := s.recipeRepo.ListRecipes(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Error().Msgf("Can not get list recipes, error: %s", err)
		return nil, err
	}
	resultRecipes := make([]*recipeApi.RecipeV1, 0, len(dbRecipes))
	for idx := range dbRecipes {
		apiRecipe := recipeApi.RecipeV1{
			RecipeId:    dbRecipes[idx].Id(),
			UserId:      dbRecipes[idx].UserId(),
			Name:        dbRecipes[idx].Name(),
			Description: dbRecipes[idx].Description(),
			Actions:     dbRecipes[idx].Actions(),
		}
		resultRecipes = append(resultRecipes, &apiRecipe)
	}
	return &recipeApi.ListRecipesResponseV1{Recipes: resultRecipes}, nil
}
