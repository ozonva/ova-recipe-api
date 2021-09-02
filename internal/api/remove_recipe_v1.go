package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) RemoveRecipeV1(ctx context.Context, req *recipeApi.RemoveRecipeRequestV1) (*recipeApi.RemoveRecipesResponseV1, error) {
	if sendError := s.sendKafkaDeleteEvent(); sendError != nil {
		log.Error().Msgf("Can not send delete event to kafka, error: %s", sendError)
	}
	log.Info().Msgf("Receive new remove request: %s", req.String())
	if validationErr := req.Validate(); validationErr != nil {
		log.Error().Msgf("Invalid remove recipe request, error: %s", validationErr)
		return nil, status.Error(codes.InvalidArgument, validationErr.Error())
	}
	if err := s.recipeRepo.RemoveRecipe(ctx, req.RecipeId); err != nil {
		log.Error().Msgf("Can not remove recipe, error: %s", err)
		return nil, err
	}
	s.metrics.incSuccessRemoveRecipeCounter()
	return &recipeApi.RemoveRecipesResponseV1{RecipeId: req.RecipeId}, nil
}
