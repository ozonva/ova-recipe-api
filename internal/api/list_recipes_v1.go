package api

import (
	"context"
	"github.com/rs/zerolog/log"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) ListRecipesV1(ctx context.Context, req *recipeApi.ListRecipesRequestV1) (*recipeApi.ListRecipesResponseV1, error) {
	log.Info().Msgf("Receive new list request: %s", req.String())
	return &recipeApi.ListRecipesResponseV1{}, nil
}
