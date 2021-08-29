package api

import (
	"context"
	"github.com/rs/zerolog/log"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) DescribeRecipeV1(ctx context.Context, req *recipeApi.DescribeRecipeRequestV1) (*recipeApi.DescribeRecipeResponseV1, error) {
	log.Info().Msgf("Receive new describe request: %s", req.String())
	dbRecipe, err := s.recipeRepo.DescribeRecipe(ctx, req.RecipeId)
	if err != nil {
		log.Error().Msgf("Can not get recipe, error: %s", err)
		return nil, err
	}
	apiRecipe := recipeApi.RecipeV1{
		RecipeId:    dbRecipe.Id(),
		UserId:      dbRecipe.UserId(),
		Name:        dbRecipe.Name(),
		Description: dbRecipe.Description(),
		Actions:     dbRecipe.Actions(),
	}
	return &recipeApi.DescribeRecipeResponseV1{Recipe: &apiRecipe}, nil
}
