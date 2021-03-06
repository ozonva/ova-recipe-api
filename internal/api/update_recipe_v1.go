package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ova-recipe-api/internal/recipe"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) UpdateRecipeV1(ctx context.Context, req *recipeApi.UpdateRecipeRequestV1) (resp *recipeApi.UpdateRecipeResponseV1, err error) {
	defer func() {
		if err != nil {
			s.metrics.incFailUpdateRecipeCounter()
		} else {
			s.metrics.incSuccessUpdateRecipeCounter()
		}
	}()
	if sendError := s.sendKafkaUpdateEvent(); sendError != nil {
		log.Error().Msgf("Can not send update event to kafka, error: %s", sendError)
	}
	log.Info().Msgf("Receive new update request: %s", req.String())
	if validationErr := req.Validate(); validationErr != nil {
		log.Error().Msgf("Invalid update recipe request, error: %s", validationErr)
		return nil, status.Error(codes.InvalidArgument, validationErr.Error())
	}
	newRecipe := recipe.New(req.Recipe.RecipeId, req.Recipe.UserId, req.Recipe.Name, req.Recipe.Description, req.Recipe.Actions)
	if err := s.recipeRepo.UpdateRecipe(ctx, newRecipe); err != nil {
		log.Error().Msgf("Can not update recipe, error: %s", err)
		return nil, err
	}
	return &recipeApi.UpdateRecipeResponseV1{}, nil
}
