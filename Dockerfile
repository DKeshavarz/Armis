# Stage 1: Build
FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \ 
    GOFLAGS=-mod=vendor

WORKDIR /app

COPY go.mod go.sum ./

COPY vendor/ ./vendor/

COPY . .

RUN go build -o Armis ./cmd/main.go

# Stage 2: Run (Alpine base for CLI)
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/Armis .

EXPOSE 8080

ENTRYPOINT ["./Armis"]
