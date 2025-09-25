FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-api ./cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/go-api /app/

RUN adduser -D appuser
USER appuser

CMD ["./go-api"]