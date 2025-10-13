FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# Install swag for generating swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

# Generate swagger documentation
RUN swag init -g cmd/main.go

RUN go build -o go-api ./cmd/main.go

FROM alpine:3.20
WORKDIR /app

# Install tools for health checks
RUN apk add --no-cache wget curl netcat-openbsd

COPY --from=builder /app/go-api /app/

RUN adduser -D appuser
USER appuser

CMD ["./go-api"]