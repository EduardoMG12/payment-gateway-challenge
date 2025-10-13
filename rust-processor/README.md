# ⚡ PayGateway Rust Processor

<div align="center">

**High-performance transaction processing engine built with Rust**

*Async event-driven processor for secure and fast financial transaction processing*

[![Rust](https://img.shields.io/badge/Rust-1.70+-000000?style=flat-square&logo=rust&logoColor=white)](https://www.rust-lang.org/)
[![Tokio](https://img.shields.io/badge/Tokio-Async_Runtime-000000?style=flat-square&logo=rust&logoColor=white)](https://tokio.rs/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![RabbitMQ](https://img.shields.io/badge/RabbitMQ-Message_Broker-FF6600?style=flat-square&logo=rabbitmq&logoColor=white)](https://www.rabbitmq.com/)

</div>

## 📋 Overview

The **PayGateway Rust Processor** is the **core transaction processing engine** that handles all financial operations with maximum performance, safety, and reliability. Built with Rust's zero-cost abstractions and memory safety guarantees, it processes transactions asynchronously while maintaining strict data consistency.

---

## 📋 Table of Contents

<details>
<summary><strong>🎯 Core Responsibilities</strong></summary>

### Primary Functions

#### 📬 **Message Queue Processing**
- **RabbitMQ Integration** - Consumes messages from multiple queues
- **Async Message Handling** - Non-blocking message processing
- **Queue Management** - Dead letter queue handling and retry mechanisms  
- **Message Acknowledgment** - Reliable message processing with proper ACKs
- **Concurrent Processing** - Multiple worker threads for high throughput

#### 🔒 **Business Logic & Validation**
- **Idempotency Protection** - Prevents duplicate transaction processing
- **Balance Validation** - Ensures sufficient funds before processing
- **Transaction Rules** - Applies complex business rules and constraints
- **Fraud Detection** - Basic anomaly detection patterns
- **Data Integrity** - Maintains strict data consistency across operations

#### 💾 **Ledger Management**
- **PostgreSQL Integration** - High-performance database operations
- **Transaction Recording** - Immutable transaction ledger
- **State Management** - Transaction status tracking and updates
- **Atomic Operations** - Database transactions with ACID properties
- **Audit Trail** - Complete transaction history with timestamps

#### 💰 **Balance Calculation Engine**
- **Real-time Balance Updates** - Instant balance recalculation
- **Redis Caching** - High-speed balance caching layer
- **Cache Invalidation** - Smart cache refresh strategies
- **Balance Aggregation** - Account and card-level balance tracking
- **Performance Optimization** - Sub-millisecond balance queries

</details>

<details>
<summary><strong>🛠️ Technology Stack</strong></summary>

### Core Language & Runtime
- **Rust 1.70+** - Systems programming with memory safety
- **Tokio** - Async/await runtime for high-performance I/O
- **Futures** - Async programming primitives
- **Rayon** - Data parallelism for CPU-intensive tasks

### Message Queue Integration
- **Lapin** - High-performance RabbitMQ client
- **AMQP** - Advanced Message Queuing Protocol
- **Serde** - Serialization/deserialization framework
- **JSON** - Message format for inter-service communication

### Database Operations
- **SQLx** - Compile-time verified SQL queries
- **PostgreSQL Driver** - Native async PostgreSQL integration
- **Migrations** - Database schema versioning
- **Connection Pooling** - Optimized database connection management

### Caching & Performance
- **Redis** - In-memory data structure store
- **Redis-rs** - High-performance Redis client
- **Connection Multiplexing** - Efficient Redis connection usage
- **Serialization** - Optimized data serialization for caching

### Development & Testing
- **Cargo** - Package manager and build system
- **Cargo Watch** - Auto-rebuilding during development
- **Tokio Test** - Async testing framework
- **Criterion** - Statistical benchmarking
- **Tracing** - Structured logging and observability

### Configuration & Deployment
- **Config** - Configuration management
- **Environment Variables** - 12-factor app configuration
- **Docker Support** - Multi-stage containerized builds
- **Health Checks** - Application monitoring endpoints

</details>

<details>
<summary><strong>🚀 Quick Start</strong></summary>

### Prerequisites

- **Rust 1.70+** with Cargo
- **PostgreSQL** database
- **RabbitMQ** message broker  
- **Redis** cache (recommended)

### Development Setup

```bash
# Navigate to Rust processor directory
cd rust-processor

# Install dependencies and build
cargo build

# Set up environment variables
cp .env.example .env
# Edit .env with your configuration

# Start infrastructure services
docker-compose -f ../docker-compose.yml -f ../docker-compose.dev.yml up postgres rabbitmq redis -d

# Run database migrations (from go-api directory)
cd ../go-api
sqlx migrate run --source migrations

# Start the processor
cd ../rust-processor
cargo run
```

The processor will connect to RabbitMQ and start listening for messages.

### Development with Auto-Reload

```bash
# Install cargo-watch for auto-reloading
cargo install cargo-watch

# Start with auto-reload on file changes
cargo watch -x run

# Run with specific features
cargo watch -x 'run --features="redis-cache,metrics"'
```

### Docker Development

```bash
# Build Docker image
docker build -t paygateway-rust-processor .

# Run with Docker
docker run --env-file .env paygateway-rust-processor

# Use Docker Compose (recommended)
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up rust-processor
```

### Environment Configuration

```env
# Database Configuration
DATABASE_URL=postgres://paygateway_user:paygateway_pass@localhost:5432/paygateway_db

# RabbitMQ Configuration  
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASS=guest

# Redis Configuration
REDIS_URL=redis://localhost:6379/
REDIS_HOST=localhost
REDIS_PORT=6379

# Application Configuration
RUST_LOG=info
PROCESSOR_WORKERS=4
PROCESSOR_PREFETCH=10
RETRY_ATTEMPTS=3
RETRY_DELAY_MS=1000

# Performance Tuning
DB_POOL_SIZE=10
REDIS_POOL_SIZE=8
QUEUE_BUFFER_SIZE=1000
```

</details>

<details>
<summary><strong>📁 Project Structure</strong></summary>

```
rust-processor/
├── 📁 src/                      # Source code
│   ├── main.rs                 # Application entry point
│   ├── application.rs          # Application setup and coordination
│   ├── config.rs               # Configuration management
│   │
│   ├── 📁 processors/          # Transaction processors
│   │   ├── mod.rs             # Module declarations
│   │   ├── transaction_processor.rs # Core transaction processing
│   │   ├── balance_processor.rs # Balance calculation processor
│   │   ├── refund_processor.rs # Refund transaction processor
│   │   └── retry_processor.rs  # Failed transaction retry logic
│   │
│   ├── 📁 services/            # Business services
│   │   ├── mod.rs             # Module declarations
│   │   ├── transaction_service.rs # Transaction business logic
│   │   ├── balance_service.rs  # Balance calculation service
│   │   ├── validation_service.rs # Business rule validation
│   │   └── notification_service.rs # Event notification service
│   │
│   ├── 📁 models/              # Data models
│   │   ├── mod.rs             # Module declarations
│   │   ├── transaction.rs      # Transaction entity
│   │   ├── account.rs          # Account entity
│   │   ├── card.rs            # Card entity
│   │   └── balance.rs          # Balance entity
│   │
│   ├── 📁 database/            # Database operations
│   │   ├── mod.rs             # Module declarations
│   │   ├── connection.rs       # Database connection management
│   │   ├── transaction_repo.rs # Transaction repository
│   │   ├── account_repo.rs     # Account repository
│   │   └── migrations.rs       # Migration helpers
│   │
│   ├── 📁 connections/         # External connections
│   │   ├── mod.rs             # Module declarations
│   │   ├── rabbitmq.rs        # RabbitMQ connection and setup
│   │   ├── redis.rs           # Redis connection and operations
│   │   └── postgres.rs        # PostgreSQL connection pool
│   │
│   └── 📁 utils/              # Utility modules
│       ├── mod.rs             # Module declarations
│       ├── errors.rs          # Custom error types
│       ├── logging.rs         # Logging configuration
│       ├── metrics.rs         # Performance metrics
│       └── helpers.rs         # Helper functions
│
├── 📁 tests/                   # Test suites
│   ├── integration_tests.rs   # Integration tests
│   ├── unit_tests.rs          # Unit tests
│   └── benchmarks.rs          # Performance benchmarks
│
├── 📁 migrations/              # Database migrations (if applicable)
│   └── ...
│
├── 📄 Cargo.toml              # Package configuration
├── 📄 Cargo.lock              # Dependency lock file
├── 📄 Dockerfile              # Production Docker config
├── 📄 Dockerfile.dev          # Development Docker config
├── 📄 .env.example            # Environment template
└── 📄 rust-toolchain.toml     # Rust version specification
```

</details>

<details>
<summary><strong>🔄 Message Processing</strong></summary>

### Queue Configuration

The processor listens to multiple RabbitMQ queues:

#### Primary Queues
- **`transactions_queue`** - Main transaction processing
- **`balance_calculation_queue`** - Balance update requests
- **`refund_queue`** - Refund processing
- **`retry_queue`** - Failed transaction retries

#### Dead Letter Queues
- **`transactions_dlq`** - Failed transactions
- **`balance_dlq`** - Failed balance calculations
- **`system_dlq`** - System-level failures

### Message Flow

```mermaid
graph TD
    A[Go API] -->|Publish Message| B[RabbitMQ Exchange]
    B --> C[Transaction Queue]
    B --> D[Balance Queue]
    B --> E[Refund Queue]
    
    C --> F[Transaction Processor]
    D --> G[Balance Processor]  
    E --> H[Refund Processor]
    
    F --> I[Validation Service]
    F --> J[Database Service]
    F --> K[Cache Service]
    
    I --> L{Valid?}
    L -->|Yes| M[Process Transaction]
    L -->|No| N[Reject & Log]
    
    M --> O[Update Database]
    M --> P[Update Cache]
    M --> Q[Send Notification]
```

### Message Types

#### Transaction Message
```rust
#[derive(Debug, Serialize, Deserialize)]
pub struct TransactionMessage {
    pub transaction_id: Uuid,
    pub account_id: Uuid,
    pub card_id: Option<Uuid>,
    pub amount_cents: i64,
    pub transaction_type: TransactionType,
    pub idempotency_key: String,
    pub timestamp: DateTime<Utc>,
    pub metadata: Option<serde_json::Value>,
}
```

#### Balance Update Message
```rust
#[derive(Debug, Serialize, Deserialize)]
pub struct BalanceUpdateMessage {
    pub account_id: Uuid,
    pub card_id: Option<Uuid>,
    pub force_recalculation: bool,
    pub timestamp: DateTime<Utc>,
}
```

### Error Handling & Retries

#### Retry Strategy
1. **Immediate Retry** - Network/connection errors
2. **Exponential Backoff** - Temporary failures
3. **Dead Letter Queue** - Permanent failures
4. **Manual Intervention** - Critical system errors

#### Error Categories
- **Validation Errors** - Business rule violations
- **System Errors** - Database/network failures  
- **Processing Errors** - Logic errors in processing
- **External Errors** - Third-party service failures

</details>

<details>
<summary><strong>🔧 Development</strong></summary>

### Available Commands

```bash
# Development
cargo run                       # Start the processor
cargo build                     # Build in debug mode
cargo build --release          # Build optimized for production
cargo check                     # Fast compilation check

# Testing
cargo test                      # Run all tests
cargo test --lib              # Run library tests only
cargo test integration_tests   # Run integration tests
cargo bench                    # Run benchmarks

# Code Quality
cargo fmt                      # Format code
cargo clippy                   # Lint code for improvements
cargo audit                    # Security audit of dependencies

# Documentation
cargo doc                      # Generate documentation
cargo doc --open              # Generate and open documentation

# Development Tools
cargo watch -x run             # Auto-reload on changes
cargo watch -x test           # Auto-test on changes
cargo watch -x check          # Auto-check on changes
```

### Testing Strategy

#### Unit Tests
```bash
# Run specific module tests
cargo test processors::transaction_processor
cargo test services::validation_service
cargo test database::transaction_repo

# Run with output
cargo test -- --nocapture

# Run single test
cargo test test_transaction_validation -- --exact
```

#### Integration Tests  
```bash
# Run integration tests (requires test infrastructure)
docker-compose -f docker-compose.test.yml up -d
cargo test integration_tests
docker-compose -f docker-compose.test.yml down
```

#### Benchmarks
```bash
# Run performance benchmarks
cargo bench

# Specific benchmark
cargo bench --bench transaction_processing

# Profile with flamegraph (requires cargo-flamegraph)
cargo flamegraph --bench transaction_processing
```

### Performance Optimization

#### Profiling Tools
```bash
# Install profiling tools
cargo install cargo-flamegraph
cargo install cargo-profdata

# Profile CPU usage
cargo flamegraph --bin rust-processor

# Memory profiling with heaptrack (Linux)
heaptrack cargo run --release
```

#### Optimization Strategies
- **Zero-Copy Deserialization** - Avoid unnecessary memory allocations
- **Connection Pooling** - Reuse database/Redis connections
- **Batch Processing** - Process multiple messages in batches
- **Parallel Processing** - Utilize multiple CPU cores
- **Cache Warming** - Pre-load frequently accessed data

### Configuration Management

#### Environment-based Configuration
```rust
#[derive(Debug, Clone)]
pub struct AppConfig {
    pub database_url: String,
    pub rabbitmq_url: String,
    pub redis_url: String,
    pub worker_count: usize,
    pub prefetch_count: u16,
    pub retry_attempts: u8,
}

impl AppConfig {
    pub fn from_env() -> Result<Self, ConfigError> {
        // Load from environment variables
    }
}
```

#### Feature Flags
```bash
# Build with specific features
cargo build --features "redis-cache"
cargo build --features "metrics,tracing"
cargo build --no-default-features --features "minimal"
```

</details>

<details>
<summary><strong>📊 Monitoring & Observability</strong></summary>

### Logging

#### Structured Logging with Tracing
```rust
use tracing::{info, warn, error, debug, instrument};

#[instrument(skip(self))]
pub async fn process_transaction(&self, message: TransactionMessage) -> Result<(), ProcessError> {
    info!(
        transaction_id = %message.transaction_id,
        account_id = %message.account_id,
        amount_cents = message.amount_cents,
        "Processing transaction"
    );
    
    // Processing logic...
    
    info!(
        transaction_id = %message.transaction_id,
        duration_ms = start.elapsed().as_millis(),
        "Transaction processed successfully"
    );
}
```

#### Log Levels and Output
- **ERROR** - System errors and failures
- **WARN** - Business rule violations and retries
- **INFO** - Transaction processing events
- **DEBUG** - Detailed processing information
- **TRACE** - Fine-grained debugging information

### Metrics Collection

#### Performance Metrics
```rust
use prometheus::{Counter, Histogram, Gauge};

pub struct ProcessorMetrics {
    pub transactions_processed: Counter,
    pub processing_duration: Histogram,
    pub active_connections: Gauge,
    pub queue_depth: Gauge,
}
```

#### Key Metrics
- **Transaction Throughput** - Transactions per second
- **Processing Latency** - Time to process each transaction
- **Error Rates** - Failed transaction percentage
- **Queue Depth** - Messages waiting for processing
- **Resource Usage** - CPU, memory, database connections

### Health Checks

#### Application Health
```rust
#[derive(Debug, Serialize)]
pub struct HealthStatus {
    pub status: String,
    pub timestamp: DateTime<Utc>,
    pub services: ServiceHealth,
    pub metrics: HealthMetrics,
}

#[derive(Debug, Serialize)]  
pub struct ServiceHealth {
    pub database: ServiceStatus,
    pub rabbitmq: ServiceStatus,
    pub redis: ServiceStatus,
}
```

#### Health Check Endpoints
- **Database Connectivity** - PostgreSQL connection test
- **Message Queue** - RabbitMQ connection and queue status
- **Cache Service** - Redis connectivity and performance
- **Processing Queue** - Message backlog and processing rate

</details>

<details>
<summary><strong>🚀 Deployment</strong></summary>

### Production Build

```bash
# Optimized release build
cargo build --release

# Build with specific optimizations
RUSTFLAGS="-C target-cpu=native" cargo build --release

# Strip symbols for smaller binary
cargo build --release
strip target/release/rust-processor
```

### Docker Deployment

#### Multi-Stage Dockerfile
```dockerfile
# Build stage
FROM rust:1.70-slim as builder
WORKDIR /app
COPY Cargo.toml Cargo.lock ./
COPY src ./src
RUN cargo build --release

# Runtime stage
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/target/release/rust-processor /usr/local/bin/
CMD ["rust-processor"]
```

#### Production Environment
```env
# Production configuration
RUST_LOG=info
DATABASE_URL=postgres://user:pass@prod-db:5432/paygateway
RABBITMQ_URL=amqp://user:pass@prod-rabbitmq:5672/
REDIS_URL=redis://prod-redis:6379/

# Performance tuning
PROCESSOR_WORKERS=8
DB_POOL_SIZE=20
REDIS_POOL_SIZE=16
QUEUE_BUFFER_SIZE=2000
```

### Scaling Strategies

#### Horizontal Scaling
- **Multiple Instances** - Run multiple processor instances
- **Load Balancing** - Distribute messages across instances
- **Queue Partitioning** - Separate queues by processing type
- **Geographic Distribution** - Deploy across regions

#### Vertical Scaling
- **CPU Optimization** - Increase worker thread count
- **Memory Tuning** - Optimize connection pool sizes
- **I/O Performance** - SSD storage and network optimization
- **Database Scaling** - Read replicas and connection pooling

### Monitoring in Production

#### System Monitoring
- **Resource Usage** - CPU, memory, disk, network
- **Application Performance** - Request latency, throughput
- **Error Tracking** - Error rates and categorization  
- **Business Metrics** - Transaction volume, success rates

#### Alerting Rules
- **High Error Rate** - >5% transaction failures
- **Processing Delays** - Queue depth >1000 messages
- **Resource Exhaustion** - >80% CPU/memory usage
- **External Service Issues** - Database/Redis connectivity

</details>

---

## 🧪 Testing

### Test Categories

```bash
# Unit tests for individual components
cargo test --lib

# Integration tests with external services
cargo test --test integration_tests

# Performance benchmarks
cargo bench

# Load testing with multiple scenarios
cargo test --test load_tests --release
```

### Performance Benchmarks

Current performance targets:
- **Transaction Processing**: >10,000 transactions/second
- **Balance Updates**: <10ms average latency
- **Memory Usage**: <200MB under normal load
- **CPU Efficiency**: <50% CPU at peak load

## 🔒 Security

### Security Features

- **Memory Safety** - Rust's ownership system prevents common vulnerabilities
- **Input Validation** - Strict validation of all incoming messages
- **Database Security** - Prepared statements prevent SQL injection
- **Connection Security** - TLS encryption for all external connections
- **Audit Logging** - Comprehensive transaction audit trail

## 🤝 Contributing

When contributing to the Rust processor:

1. Follow Rust best practices and idioms
2. Write comprehensive tests for new functionality
3. Ensure thread safety for concurrent operations
4. Add proper error handling and logging
5. Update documentation and examples

## 📄 License

This Rust processor is part of the PayGateway project and follows the same [MIT License](../LICENSE).

---

<div align="center">

**Part of the PayGateway ecosystem**

[🏠 Main Project](../README.md) • [🌐 Frontend](../frontend/README.md) • [🚀 Go API](../go-api/README.md)

</div>
- **SQLx:** Asynchronous SQL toolkit
- **Redis:** Client for the cache
- **Serde:** Data serialization and deserialization