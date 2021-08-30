.PHONY: run, build, test-internal
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

.PHONY: run
run:
	go run ./cmd/ova-recipe-api.go

.PHONY: deps
deps:
	go get github.com/onsi/ginkgo@v1.16.4
	go get github.com/onsi/gomega@v1.16.0
	go get github.com/golang/mock@v1.6.0
	go get github.com/rs/zerolog/log@v1.23.0
	go get google.golang.org/grpc@v1.40.0
	go get -d github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go get github.com/golang/protobuf/proto@v1.5.2
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

.PHONY: build
build:
	go build -o bin/ova-recipe-api ./cmd/ova-recipe-api.go

.PHONY: test-internal
test-internal:
	rm -rf ./internal/*/*_mock.go
	mockgen -destination=./internal/recipe/action_mock.go -package=recipe ova-recipe-api/internal/recipe Action
	mockgen -destination=./internal/repo/repo_mock.go -package=repo ova-recipe-api/internal/repo RecipeRepo
	mockgen -destination=./internal/flusher/flusher_mock.go -package=flusher ova-recipe-api/internal/flusher Flusher
	mockgen -destination=./internal/ticker/ticker_mock.go -package=ticker ova-recipe-api/internal/ticker Ticker
	go test -race ./internal/...
	ginkgo -race ./internal/...

.PHONY: generate
generate:
	protoc \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	api/api.proto
