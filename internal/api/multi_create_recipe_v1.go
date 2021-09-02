package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"ova-recipe-api/internal/recipe"
	"ova-recipe-api/internal/utils"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

const batchSize = uint(2)

func (s *GRPCServer) MultiCreateRecipeV1(ctx context.Context, req *recipeApi.MultiCreateRecipeRequestV1) (*emptypb.Empty, error) {
	if sendError := s.sendKafkaCreateEvent(); sendError != nil {
		log.Error().Msgf("Can not send create event to kafka, error: %s", sendError)
	}
	log.Info().Msgf("Receive new multi create request: %s", req.String())
	if validationErr := req.Validate(); validationErr != nil {
		log.Error().Msgf("Invalid multi create recipe request, error: %s", validationErr)
		return nil, status.Error(codes.InvalidArgument, validationErr.Error())
	}

	newRecipes := make([]recipe.Recipe, 0, len(req.Recipes))
	for _, r := range req.Recipes {
		newRecipes = append(newRecipes, recipe.New(0, r.UserId, r.Name, r.Description, r.Actions))
	}
	for _, recipesSlice := range utils.SplitRecipeSlice(newRecipes, batchSize) {
		if addErr := s.recipeRepo.AddRecipes(ctx, recipesSlice); addErr != nil {
			log.Error().Msgf("Can not add new recipes, error: %s", addErr)
			return nil, addErr
		}
	}

	return &emptypb.Empty{}, nil
}
