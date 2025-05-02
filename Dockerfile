FROM golang:1.23 AS builder

WORKDIR /app
COPY telegram/ /app
RUN go mod tidy
RUN go build -o app

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]