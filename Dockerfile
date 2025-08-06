# Stage 1: Build
FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o Armis ./cmd/main.go


# Stage 2: Run (Alpine base for CLI)
FROM alpine:latest

WORKDIR /app

COPY .env .env
COPY --from=builder /app/Armis .

EXPOSE 8080

ENTRYPOINT ["./Armis"]
