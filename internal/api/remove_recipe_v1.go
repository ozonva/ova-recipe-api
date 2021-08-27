package api

import (
	"context"
	"github.com/rs/zerolog/log"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) RemoveRecipeV1(ctx context.Context, req *recipeApi.RemoveRecipeRequestV1) (*recipeApi.RemoveRecipesResponseV1, error) {
	log.Info().Msgf("Receive new remove request: %s", req.String())
	return &recipeApi.RemoveRecipesResponseV1{}, nil
}
