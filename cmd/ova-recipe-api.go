package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
	"io"
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

func initTracer() (opentracing.Tracer, io.Closer) {
	cfg := jaegercfg.Configuration{
		ServiceName: "ova-recipe-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Fatal().Msgf("Can not create tracer, %s", err)
	}
	return tracer, closer
}

func main() {
	fmt.Println("Hi, i am ova-recipe-api!")

	if loadEnvErr := godotenv.Load(); loadEnvErr != nil {
		log.Fatal().Msgf("Can not load .env file, error: %s", loadEnvErr)
	}

	go runPrometheusMetrics()

	tracer, closer := initTracer()
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

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
	kafkaDsn := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))
	if kafkaConnErr := kafkaClient.Connect(context.Background(), kafkaDsn, "CUDEvents", 0); kafkaConnErr != nil {
		log.Fatal().Msgf("Can not connect to kafka, %s", kafkaConnErr)
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
