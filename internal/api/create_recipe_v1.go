package api

import (
	"context"
	"github.com/rs/zerolog/log"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) CreateRecipeV1(ctx context.Context, req *recipeApi.CreateRecipeRequestV1) (*recipeApi.CreateRecipeResponseV1, error) {
	log.Info().Msgf("Receive new create request: %s", req.String())
	return &recipeApi.CreateRecipeResponseV1{}, nil
}
