FROM golang:1.17.0-buster AS builder

WORKDIR /app
RUN apt-get update && apt-get install -y protobuf-compiler
COPY . /app/
RUN make all

FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=builder /app/bin/ova-recipe-api .
COPY --from=builder /app/.env .
CMD ["/app/ova-recipe-api"]