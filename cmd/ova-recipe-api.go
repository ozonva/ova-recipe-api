package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"ova-recipe-api/internal/api"
	"ova-recipe-api/internal/repo"
	recipeApi "ova-recipe-api/pkg/api/github.com/ozonva/ova-recipe-api/pkg/api"
)

func main() {
	fmt.Println("Hi, i am ova-recipe-api!")

	db, openDbErr := repo.OpenDb("postgres://admin:12345678@db:5432/recipe_api?sslmode=disable")
	if openDbErr != nil {
		log.Fatal().Msgf("Can not open db, %s", openDbErr)
	}
	recipeRepo, newRepoErr := repo.New(db)
	if newRepoErr != nil {
		log.Fatal().Msgf("Can not create recipeRepo, %s", newRepoErr)
	}

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Msgf("Failed to listen server %s", err)
	}
	service := grpc.NewServer()
	recipeApi.RegisterOvaRecipeApiServer(service, api.NewOvaRecipeApiServer(recipeRepo))
	if err = service.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %s", err)
	}
}
