package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ova-recipe-api/internal/recipe"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func (s *GRPCServer) CreateRecipeV1(ctx context.Context, req *recipeApi.CreateRecipeRequestV1) (*recipeApi.CreateRecipeResponseV1, error) {
	if sendError := s.sendKafkaCreateEvent(); sendError != nil {
		log.Error().Msgf("Can not send create event to kafka, error: %s", sendError)
	}
	log.Info().Msgf("Receive new create request: %s", req.String())
	if validationErr := req.Validate(); validationErr != nil {
		log.Error().Msgf("Invalid create recipe request, error: %s", validationErr)
		return nil, status.Error(codes.InvalidArgument, validationErr.Error())
	}
	newRecipeId, err := s.recipeRepo.AddRecipe(ctx, recipe.New(0, req.UserId, req.Name, req.Description, req.Actions))
	if err != nil {
		log.Error().Msgf("Can not create new recipe, error: %s", err)
		return nil, err
	}
	s.metrics.incSuccessCreateRecipeCounter()
	return &recipeApi.CreateRecipeResponseV1{RecipeId: newRecipeId}, nil
}
