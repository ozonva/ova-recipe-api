LOCAL_BIN:=$(CURDIR)/bin

.PHONY: run, build, test-internal
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

.PHONY: run
run:
	go run ./cmd/ova-recipe-api.go

.PHONY: deps
deps:
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/gomega
	go get -u github.com/golang/mock
	go get -u github.com/rs/zerolog/log
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

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
