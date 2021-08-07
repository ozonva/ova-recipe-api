run:
	go run ./cmd/ova-recipe-api.go

build:
	go build -o bin/ova-recipe-api ./cmd/ova-recipe-api.go

test-internal:
	go test ./internal/...
