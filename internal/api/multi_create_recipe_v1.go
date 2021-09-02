package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	opentracingLog "github.com/opentracing/opentracing-go/log"
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

	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "MultiCreateRecipeV1")
	defer parentSpan.Finish()

	for _, recipesSlice := range utils.SplitRecipeSlice(newRecipes, batchSize) {
		if insertErr := s.batchInsert(ctx, parentSpan, recipesSlice); insertErr != nil {
			return nil, insertErr
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) batchInsert(ctx context.Context, parentSpan opentracing.Span, recipesSlice []recipe.Recipe) error {
	childSpan := opentracing.StartSpan("MultiCreateRecipeV1Batch", opentracing.ChildOf(parentSpan.Context()))
	childSpan.LogFields(
		opentracingLog.Int("Recipes count", len(recipesSlice)),
	)
	defer childSpan.Finish()
	if addErr := s.recipeRepo.AddRecipes(ctx, recipesSlice); addErr != nil {
		log.Error().Msgf("Can not add new recipes, error: %s", addErr)
		return addErr
	}
	return nil
}
