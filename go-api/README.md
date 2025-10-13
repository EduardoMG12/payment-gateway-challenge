# 🚀 PayGateway Go API

<div align="center">

**High-performance HTTP API Gateway built with Go**

*RESTful API server providing secure transaction processing and account management*

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-Web_Framework-00ADD8?style=flat-square&logo=go&logoColor=white)](https://gin-gonic.com/)
[![Swagger](https://img.shields.io/badge/Swagger-API_Docs-85EA2D?style=flat-square&logo=swagger&logoColor=black)](https://swagger.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org/)

</div>

## 📋 Overview

The **PayGateway Go API** serves as the **HTTP gateway** and **orchestration layer** for the PayGateway system. It provides RESTful endpoints for account management, card operations, and transaction processing while maintaining loose coupling through asynchronous message publishing.

---

## 📋 Table of Contents

<details>
<summary><strong>🎯 Core Responsibilities</strong></summary>

### Primary Functions

#### 🏦 **Account & Card Management**
- **Account Creation** - User registration and profile management
- **Account Retrieval** - Secure account data access
- **Virtual Card Generation** - Secure tokenization and card creation
- **Card Management** - List and manage user payment methods
- **Balance Tracking** - Real-time account balance calculations

#### 🔐 **Security & Tokenization**
- **Card Tokenization** - Convert sensitive card data to secure tokens
- **Hash Generation** - Secure internal representation of payment methods
- **Input Validation** - Comprehensive request validation and sanitization
- **CORS Management** - Cross-origin resource sharing configuration
- **Rate Limiting** - API rate limiting and abuse prevention

#### 📡 **Asynchronous Communication**
- **Message Publishing** - RabbitMQ integration for decoupled processing
- **Event Sourcing** - Transaction event publishing for audit trails
- **Queue Management** - Reliable message delivery and error handling
- **Status Tracking** - Transaction status polling and updates

#### 🔍 **Transaction Orchestration**
- **Transaction Initiation** - Validate and initiate payment requests
- **Business Rule Validation** - Pre-processing validation logic
- **Status Polling** - Real-time transaction status queries
- **Error Handling** - Comprehensive error responses and logging

</details>

<details>
<summary><strong>🛠️ Technology Stack</strong></summary>

### Core Framework
- **Go 1.21+** - High-performance compiled language
- **Gin Web Framework** - Fast HTTP web framework
- **GORM** - Feature-rich ORM for database operations
- **Validator** - Struct and field validation

### Database & Storage
- **PostgreSQL** - Primary relational database
- **SQLX** - SQL builder and query interface
- **Database Migrations** - Schema versioning and management
- **Connection Pooling** - Optimized database connections

### Message Queue Integration
- **RabbitMQ** - AMQP message broker integration
- **Streadway AMQP** - Go RabbitMQ client library
- **Message Serialization** - JSON message formatting
- **Dead Letter Queues** - Failed message handling

### Documentation & Testing
- **Swaggo** - Automatic Swagger documentation generation
- **Testify** - Testing framework and assertions
- **Gomock** - Mock generation for testing
- **Test Coverage** - Comprehensive test coverage reports

### Configuration & Deployment
- **Viper** - Configuration management
- **Environment Variables** - 12-factor app configuration
- **Docker Support** - Containerized deployment
- **Health Checks** - Application health monitoring

</details>

<details>
<summary><strong>🚀 Quick Start</strong></summary>

### Prerequisites

- **Go 1.21+** installed
- **PostgreSQL** database running
- **RabbitMQ** message broker running
- **Redis** cache (optional but recommended)

### Development Setup

```bash
# Navigate to Go API directory
cd go-api

# Download dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your configuration

# Run database migrations
export DATABASE_URL="postgres://user:pass@localhost:5432/paygateway_db?sslmode=disable"
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -path migrations -database $DATABASE_URL up

# Start the development server
go run cmd/main.go
```

The API will be available at **http://localhost:8080**

### Docker Development

```bash
# Build and run with Docker
docker build -t paygateway-go-api .
docker run -p 8080:8080 paygateway-go-api

# Or use Docker Compose
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up go-api
```

### Environment Configuration

```env
# Database Configuration
DATABASE_URL=postgres://user:pass@localhost:5432/paygateway_db?sslmode=disable
DB_HOST=localhost
DB_PORT=5432
DB_USER=paygateway_user
DB_PASSWORD=paygateway_pass
DB_NAME=paygateway_db

# RabbitMQ Configuration
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASS=guest

# Redis Configuration (optional)
REDIS_URL=redis://localhost:6379
REDIS_HOST=localhost
REDIS_PORT=6379

# API Configuration
API_PORT=8080
API_HOST=0.0.0.0
GIN_MODE=debug

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:8081
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=Content-Type,Authorization
```

</details>

<details>
<summary><strong>📁 Project Structure</strong></summary>

```
go-api/
├── 📁 cmd/                      # Application entrypoints
│   └── main.go                 # Main application entry
│
├── 📁 internal/                # Private application code
│   ├── 📁 account/            # Account domain
│   │   ├── handler.go         # HTTP handlers
│   │   ├── service.go         # Business logic
│   │   ├── module.go          # Dependency injection
│   │   └── 📁 dto/           # Data transfer objects
│   │
│   ├── 📁 card/              # Card domain
│   │   ├── handler.go         # HTTP handlers
│   │   ├── service.go         # Business logic
│   │   ├── module.go          # Dependency injection
│   │   └── 📁 dto/           # Data transfer objects
│   │
│   ├── 📁 transaction/        # Transaction domain
│   │   ├── handler.go         # HTTP handlers
│   │   ├── service.go         # Business logic
│   │   ├── module.go          # Dependency injection
│   │   └── 📁 dto/           # Data transfer objects
│   │
│   ├── 📁 config/            # Configuration
│   │   ├── config.go          # Application configuration
│   │   ├── db_connection.go   # Database setup
│   │   ├── rabbitmq_connection.go # Message queue setup
│   │   ├── redis_connection.go # Cache setup
│   │   └── cors.go            # CORS configuration
│   │
│   ├── 📁 models/            # Database models
│   │   ├── account.go         # Account entity
│   │   ├── card.go            # Card entity
│   │   └── transaction.go     # Transaction entity
│   │
│   ├── 📁 repository/        # Data access layer
│   │   ├── account.go         # Account repository
│   │   ├── card.go            # Card repository
│   │   └── transaction.go     # Transaction repository
│   │
│   ├── 📁 router/            # HTTP routing
│   │   └── router.go          # Route definitions
│   │
│   ├── 📁 connection/        # External connections
│   │   ├── db.go             # Database connection
│   │   ├── rabbitmq.go       # RabbitMQ connection
│   │   └── redis.go          # Redis connection
│   │
│   ├── 📁 utils/             # Utility functions
│   │   ├── hash.go           # Hashing utilities
│   │   ├── validation.go     # Validation helpers
│   │   └── response.go       # Response helpers
│   │
│   └── 📁 i18n/             # Internationalization
│       └── errors.go          # Error messages
│
├── 📁 migrations/              # Database migrations
│   ├── 001_create_tables.sql  # Initial schema
│   ├── 002_gen_random_uuid.sql # UUID support
│   ├── 003_idempotency_key.sql # Idempotency support
│   └── ... (more migrations)
│
├── 📁 docs/                   # API documentation
│   ├── docs.go               # Swagger docs
│   ├── swagger.json          # Generated JSON docs
│   └── swagger.yaml          # Generated YAML docs
│
├── 📁 tests/                  # Test suites
│   ├── 📁 integration/       # Integration tests
│   └── 📁 e2e/              # End-to-end tests
│
├── 📄 go.mod                  # Go module definition
├── 📄 go.sum                  # Dependency checksums
├── 📄 Dockerfile             # Docker configuration
├── 📄 Dockerfile.dev         # Development Docker config
└── 📄 .env.example           # Environment template
```

</details>

<details>
<summary><strong>🔗 API Documentation</strong></summary>

### Swagger Documentation

**Interactive API documentation is available at:**

🌐 **Development:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Authentication

Currently, the API uses a simplified authentication model for development purposes. In production, implement:

- **JWT Tokens** for stateless authentication
- **API Keys** for service-to-service communication
- **Rate Limiting** per user/IP
- **OAuth 2.0** for third-party integrations

### Core Endpoints

#### 👤 **Account Management**

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `POST` | `/accounts` | Create new account | `{"username": "string"}` |
| `GET` | `/accounts` | List all accounts | - |
| `GET` | `/accounts/{id}` | Get account by ID | - |
| `GET` | `/accounts/{id}/balance` | Get account balance | - |

#### 💳 **Card Operations**

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `POST` | `/cards` | Create new virtual card | `{"account_id": "uuid"}` |
| `GET` | `/cards/{accountId}` | List account cards | - |
| `GET` | `/cards/{id}` | Get card details | - |
| `DELETE` | `/cards/{id}` | Deactivate card | - |

#### 💰 **Transaction Processing**

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `POST` | `/transactions` | Process transaction | `{"type": "PURCHASE", "amount_cents": 1000, "card_token": "string"}` |
| `GET` | `/transactions/{accountId}` | Get transaction history | - |
| `GET` | `/transactions/{id}` | Get transaction details | - |
| `POST` | `/transactions/{id}/refund` | Process refund | - |

#### 🔍 **System Endpoints**

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/swagger/*` | API documentation |
| `GET` | `/metrics` | Prometheus metrics |

### Request/Response Examples

#### Create Account
```bash
# Request
curl -X POST http://localhost:8080/accounts \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe"}'

# Response
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "username": "john_doe",
  "created_at": "2023-10-12T10:30:00Z",
  "balance": 0
}
```

#### Process Transaction
```bash
# Request
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "type": "PURCHASE",
    "amount_cents": 5000,
    "card_token": "tok_1234567890abcdef",
    "idempotency_key": "unique-key-123"
  }'

# Response
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "type": "PURCHASE",
  "amount_cents": 5000,
  "status": "PROCESSING",
  "created_at": "2023-10-12T10:35:00Z"
}
```

</details>

<details>
<summary><strong>🔧 Development</strong></summary>

### Available Commands

```bash
# Development
go run cmd/main.go              # Start development server
go build -o bin/api cmd/main.go # Build binary
go test ./...                   # Run all tests
go test -cover ./...           # Run tests with coverage

# Code Quality
go fmt ./...                    # Format code
go vet ./...                    # Examine code for issues
golint ./...                   # Lint code (install: go install golang.org/x/lint/golint@latest)

# Dependencies
go mod tidy                     # Clean up dependencies
go mod vendor                   # Vendor dependencies
go mod download                 # Download dependencies

# Documentation
swag init -g cmd/main.go        # Generate Swagger docs
```

### Testing Strategy

#### Unit Tests
```bash
# Run unit tests
go test ./internal/account/...
go test ./internal/card/...
go test ./internal/transaction/...

# Test with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

#### Integration Tests
```bash
# Run integration tests (requires test database)
go test ./tests/integration/...

# Run with test database
export TEST_DATABASE_URL="postgres://test:test@localhost:5433/paygateway_test"
go test ./tests/integration/...
```

### Database Operations

#### Migrations
```bash
# Create new migration
migrate create -ext sql -dir migrations -seq migration_name

# Run migrations
migrate -path migrations -database $DATABASE_URL up

# Rollback migrations
migrate -path migrations -database $DATABASE_URL down 1

# Check migration status
migrate -path migrations -database $DATABASE_URL version
```

#### Development Database
```bash
# Start local PostgreSQL with Docker
docker run --name paygateway-postgres \
  -e POSTGRES_DB=paygateway_db \
  -e POSTGRES_USER=paygateway_user \
  -e POSTGRES_PASSWORD=paygateway_pass \
  -p 5432:5432 \
  -d postgres:15
```

### Performance Optimization

#### Database Optimization
- **Connection Pooling** - Configured connection limits
- **Query Optimization** - Indexed queries and efficient JOINs  
- **Prepared Statements** - Compiled query caching
- **Batch Operations** - Bulk insert/update operations

#### Application Performance
- **Goroutine Pooling** - Limited concurrent request handling
- **Memory Management** - Optimized struct layouts
- **Caching Strategy** - Redis integration for hot data
- **Request Compression** - Gzip compression for responses

</details>

<details>
<summary><strong>📊 Monitoring & Observability</strong></summary>

### Health Checks

The API provides comprehensive health checking:

```bash
# Basic health check
curl http://localhost:8080/health

# Response
{
  "status": "healthy",
  "timestamp": "2023-10-12T10:30:00Z",
  "services": {
    "database": "healthy",
    "rabbitmq": "healthy", 
    "redis": "healthy"
  }
}
```

### Logging

#### Structured Logging
- **JSON Format** - Machine-readable log format
- **Log Levels** - DEBUG, INFO, WARN, ERROR, FATAL
- **Request Tracing** - Unique request ID tracking
- **Performance Metrics** - Response time logging

#### Log Examples
```json
{
  "level": "info",
  "timestamp": "2023-10-12T10:30:00Z",
  "request_id": "req-123456",
  "method": "POST",
  "path": "/transactions",
  "status": 201,
  "duration": "45ms",
  "message": "Transaction created successfully"
}
```

### Metrics Collection

#### Prometheus Metrics
- **HTTP Request Metrics** - Request count, duration, status codes
- **Database Metrics** - Connection pool stats, query duration
- **Business Metrics** - Transaction volume, success rates
- **System Metrics** - Memory usage, goroutine count

```bash
# Access metrics endpoint
curl http://localhost:8080/metrics
```

</details>

<details>
<summary><strong>🚀 Deployment</strong></summary>

### Production Build

```bash
# Build for production
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api cmd/main.go

# Build with version info
go build -ldflags "-X main.version=$(git describe --tags --always)" -o bin/api cmd/main.go
```

### Docker Deployment

#### Production Dockerfile
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
CMD ["./api"]
```

#### Docker Compose
```yaml
services:
  go-api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://user:pass@postgres:5432/db
      - RABBITMQ_URL=amqp://rabbitmq:5672/
    depends_on:
      - postgres
      - rabbitmq
```

### Environment Configuration

#### Production Environment Variables
```env
# Required
DATABASE_URL=postgres://user:pass@host:5432/db?sslmode=require
RABBITMQ_URL=amqp://user:pass@host:5672/

# Optional
REDIS_URL=redis://host:6379
API_PORT=8080
GIN_MODE=release
LOG_LEVEL=info
```

</details>

---

## 🧪 Testing

### Test Coverage

Current test coverage targets:
- **Unit Tests**: >80% coverage
- **Integration Tests**: Critical paths covered
- **E2E Tests**: Full user workflows

### Running Tests

```bash
# All tests
make test

# Unit tests only
make test-unit

# Integration tests
make test-integration

# With coverage report
make test-coverage
```

## 📈 Performance

### Benchmarks

- **Request Handling**: >1000 requests/second
- **Database Operations**: <100ms average
- **Memory Usage**: <100MB under normal load
- **Start Time**: <5 seconds

### Optimization Strategies

- Efficient JSON serialization
- Connection pooling
- Query optimization
- Caching strategies

## 🤝 Contributing

When contributing to the Go API:

1. Follow Go conventions and best practices
2. Write comprehensive tests
3. Update Swagger documentation
4. Ensure backward compatibility
5. Add proper error handling

## 📄 License

This Go API is part of the PayGateway project and follows the same [MIT License](../LICENSE).

---

<div align="center">

**Part of the PayGateway ecosystem**

[🏠 Main Project](../README.md) • [🌐 Frontend](../frontend/README.md) • [⚡ Rust Processor](../rust-processor/README.md)

</div>