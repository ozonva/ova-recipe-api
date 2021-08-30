package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) RemoveRecipeV1(ctx context.Context, req *recipeApi.RemoveRecipeRequestV1) (*recipeApi.RemoveRecipesResponseV1, error) {
	log.Info().Msgf("Receive new remove request: %s", req.String())
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.recipeRepo.RemoveRecipe(ctx, req.RecipeId); err != nil {
		log.Error().Msgf("Can not remove recipe, error: %s", err)
		return nil, err
	}
	return &recipeApi.RemoveRecipesResponseV1{RecipeId: req.RecipeId}, nil
}
