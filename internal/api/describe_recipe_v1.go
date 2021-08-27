package api

import (
	"context"
	"github.com/rs/zerolog/log"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) DescribeRecipeV1(ctx context.Context, req *recipeApi.DescribeRecipeRequestV1) (*recipeApi.DescribeRecipeResponseV1, error) {
	log.Info().Msgf("Receive new describe request: %s", req.String())
	return &recipeApi.DescribeRecipeResponseV1{}, nil
}
