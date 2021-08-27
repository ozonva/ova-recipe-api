package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"ova-recipe-api/internal/api"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func main() {
	fmt.Println("Hi, i am ova-recipe-api!")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Msgf("Failed to listen server %s", err)
	}
	service := grpc.NewServer()
	recipeApi.RegisterOvaRecipeApiServer(service, api.NewOvaRecipeApiServer())
	if err = service.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %s", err)
	}
}
