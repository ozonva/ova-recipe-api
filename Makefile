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
	go get google.golang.org/protobuf@v1.27.1
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go get github.com/jackc/pgx@v3.6.2
	go get github.com/jmoiron/sqlx@v1.3.4
	go get github.com/lib/pq@v1.10.2
	go get github.com/envoyproxy/protoc-gen-validate@v0.6.1
	go install github.com/envoyproxy/protoc-gen-validate

.PHONY: build
build:
	go build -o bin/ova-recipe-api ./cmd/ova-recipe-api.go

.PHONY: test-internal
test-internal:
	rm -rf ./internal/*/*_mock.go
	mockgen -destination=./internal/repo/repo_mock.go -package=repo ova-recipe-api/internal/repo RecipeRepo
	mockgen -destination=./internal/flusher/flusher_mock.go -package=flusher ova-recipe-api/internal/flusher Flusher
	mockgen -destination=./internal/ticker/ticker_mock.go -package=ticker ova-recipe-api/internal/ticker Ticker
	go test -race ./internal/...
	ginkgo -race ./internal/...

.PHONY: generate-proto
generate-proto:
	protoc -I vendor.protogen \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	--validate_out lang=go:pkg/api \
	api/api.proto


.PHONY: generate-vendor-proto
generate-vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api
	cp api/api.proto vendor.protogen/api/api.proto
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
		mkdir -p vendor.protogen/github.com/envoyproxy &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
	fi

.PHONY: all
all: deps generate-vendor-proto generate-proto build