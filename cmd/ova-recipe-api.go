package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"ova-recipe-api/internal/api"
	"ova-recipe-api/internal/kafka_client"
	"ova-recipe-api/internal/repo"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func runPrometheusMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal().Msgf("Failed to start listen to metric requests, error %s", err)
	}
}

func main() {
	fmt.Println("Hi, i am ova-recipe-api!")

	if loadEnvErr := godotenv.Load(); loadEnvErr != nil {
		log.Fatal().Msgf("Can not load .env file, error: %s", loadEnvErr)
	}

	ctx := context.Background()

	go runPrometheusMetrics()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, openDbErr := repo.OpenDb(dsn)
	if openDbErr != nil {
		log.Fatal().Msgf("Can not open db, %s", openDbErr)
	}
	recipeRepo, newRepoErr := repo.New(db)
	if newRepoErr != nil {
		log.Fatal().Msgf("Can not create recipeRepo, %s", newRepoErr)
	}

	kafkaClient := kafka_client.New()
	kafkaConnErr := kafkaClient.Connect(
		ctx,
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		"CUDEvents",
		0)
	if kafkaConnErr != nil {
		log.Fatal().Msgf("Can not connect to kafka, %s", newRepoErr)
	}

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Msgf("Failed to listen server %s", err)
	}
	service := grpc.NewServer()
	recipeApi.RegisterOvaRecipeApiServer(service, api.NewOvaRecipeApiServer(recipeRepo, kafkaClient))
	if err = service.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %s", err)
	}
}
