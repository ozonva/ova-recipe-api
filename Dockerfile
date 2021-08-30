FROM golang:1.17.0-buster AS builder

WORKDIR /app
RUN apt-get update && apt-get install -y protobuf-compiler
COPY . /app/
RUN make deps && make generate && make build
RUN go build -o ./bin/ova-recipe-api ./cmd/ova-recipe-api.go

FROM golang:1.17.0-buster
WORKDIR /app
COPY --from=builder /app/bin/ova-recipe-api .
CMD ["/app/ova-recipe-api"]