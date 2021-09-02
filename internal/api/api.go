package api

import (
	"ova-recipe-api/internal/kafka_client"
	"ova-recipe-api/internal/repo"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
	"strconv"
)

type GRPCServer struct {
	recipeApi.OvaRecipeApiServer
	recipeRepo  repo.RecipeRepo
	kafkaClient kafka_client.Client
	metrics     Metrics
}

type cudEvent uint8

const (
	createEvent cudEvent = iota
	updateEvent
	deleteEvent
)

func (s *GRPCServer) sendKafkaCUDEvent(event cudEvent) error {
	return s.kafkaClient.SendMessage([]byte(strconv.Itoa(int(event))))
}

func (s *GRPCServer) sendKafkaCreateEvent() error {
	return s.sendKafkaCUDEvent(createEvent)
}

func (s *GRPCServer) sendKafkaUpdateEvent() error {
	return s.sendKafkaCUDEvent(updateEvent)
}

func (s *GRPCServer) sendKafkaDeleteEvent() error {
	return s.sendKafkaCUDEvent(deleteEvent)
}

func NewOvaRecipeApiServer(recipeRepo repo.RecipeRepo, kafkaClient kafka_client.Client) recipeApi.OvaRecipeApiServer {
	return &GRPCServer{recipeRepo: recipeRepo, kafkaClient: kafkaClient ,metrics: newApiMetrics()}
}
