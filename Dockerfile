# Etapa de build
FROM golang:alpine AS builder

WORKDIR /app
COPY telegram/ /app
RUN go mod tidy && go clean -modcache
RUN go build -ldflags="-s -w" -o app

# Imagem final minimalista
FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
