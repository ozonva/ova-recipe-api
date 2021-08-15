run:
	go run ./cmd/ova-recipe-api.go

build:
	go build -o bin/ova-recipe-api ./cmd/ova-recipe-api.go

test-internal:
	rm -rf ./internal/recipe/*_mock.go
	mockgen -destination=./internal/recipe/action_mock.go -package=recipe ova-recipe-api/internal/recipe Action
	go test ./internal/...
