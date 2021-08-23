run:
	go run ./cmd/ova-recipe-api.go

build:
	go build -o bin/ova-recipe-api ./cmd/ova-recipe-api.go

test-internal:
	rm -rf ./internal/*/*_mock.go
	mockgen -destination=./internal/recipe/action_mock.go -package=recipe ova-recipe-api/internal/recipe Action
	mockgen -destination=./internal/repo/repo_mock.go -package=repo ova-recipe-api/internal/repo RecipeRepo
	mockgen -destination=./internal/flusher/flusher_mock.go -package=flusher ova-recipe-api/internal/flusher Flusher
	go test ./internal/...
	ginkgo -race ./internal/...
