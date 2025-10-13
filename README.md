# 💳 PayGateway - Fictitious Payment Gateway

<div align="center">

**A modern distributed payment processing system built with Rust, Go, and React**

*This is a learning project focused on exploring advanced software engineering and systems architecture concepts in the fintech domain*

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org/)
[![Rust](https://img.shields.io/badge/Rust-1.70+-000000?style=flat-square&logo=rust&logoColor=white)](https://www.rust-lang.org/)
[![React](https://img.shields.io/badge/React-18+-61DAFB?style=flat-square&logo=react&logoColor=black)](https://reactjs.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)

</div>

## 🎯 Project Overview

PayGateway is a **distributed payment processing system** designed as a learning platform for modern fintech architecture patterns. This project demonstrates:

- **Microservices Architecture** with async communication
- **Event-Driven Design** using message queues
- **Multi-language System** (Rust + Go + TypeScript)
- **Modern DevOps Practices** with containerization
- **Financial Domain Modeling** with proper data consistency

### 🏗️ Architecture Components

The system consists of three main services working together:

1. **🌐 Frontend (React)** - User interface for transaction management
2. **🚀 Go API** - HTTP gateway and business orchestration 
3. **⚡ Rust Processor** - High-performance transaction engine

---

## 📋 Table of Contents

<details>
<summary><strong>🏛️ System Architecture</strong></summary>

### Architecture Diagram

```mermaid
graph TB
    A[React Frontend] -->|HTTP Requests| B[Go API Gateway]
    B -->|Publish Messages| C[RabbitMQ]
    C -->|Consume Messages| D[Rust Processor]
    D -->|Read/Write| E[PostgreSQL Database]
    D -->|Cache| F[Redis Cache]
    B -->|Query Status| E
    A -->|Real-time Updates| B
    
    subgraph "Frontend Layer"
        A
    end
    
    subgraph "API Gateway Layer"
        B
    end
    
    subgraph "Message Broker"
        C
    end
    
    subgraph "Processing Engine"
        D
    end
    
    subgraph "Data Layer"
        E
        F
    end
```

### Communication Flow

1. **Frontend** sends transaction requests to Go API
2. **Go API** validates request and publishes to RabbitMQ
3. **Rust Processor** consumes messages and processes transactions
4. **Database** stores transaction ledger and account data
5. **Redis** caches balance calculations for performance
6. **Frontend** polls for transaction status updates

</details>

<details>
<summary><strong>🛠️ Technology Stack</strong></summary>

### Frontend
- **React 18** - Modern UI library
- **TypeScript** - Type-safe development
- **Vite** - Fast build tool
- **Tailwind CSS** - Utility-first styling
- **React Query** - Server state management
- **React Router** - Client-side routing

### Backend API (Go)
- **Go 1.21+** - High-performance language
- **Gin** - HTTP web framework
- **Swagger** - API documentation
- **GORM** - ORM for database operations
- **RabbitMQ Client** - Message publishing

### Processing Engine (Rust)
- **Rust 1.70+** - Systems programming language
- **Tokio** - Async runtime
- **Lapin** - RabbitMQ client
- **SQLx** - Compile-time verified queries
- **Serde** - Serialization framework

### Infrastructure
- **PostgreSQL** - Primary database
- **Redis** - Caching layer
- **RabbitMQ** - Message broker
- **Docker** - Containerization
- **Docker Compose** - Multi-container orchestration

</details>

<details>
<summary><strong>🚀 Quick Start Guide</strong></summary>

### Prerequisites

Before running the project, ensure you have:

- **Docker** and **Docker Compose** installed
- **Node.js 18+** (for frontend development)
- **Go 1.21+** (for API development)
- **Rust 1.70+** (for processor development)
- **sqlx-cli** (`cargo install sqlx-cli`)

### 🎯 One-Command Setup

```bash
# Clone the repository
git clone https://github.com/EduardoMG12/payment-gateway-challenge.git
cd payment-gateway-challenge

# Make setup script executable
chmod +x scripts/setup.sh

# Run automated setup
./scripts/setup.sh

# Start all services (recommended for development)
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up rust-processor
```

### 🔧 Development Mode (Step by Step)

1. **Environment Setup**
   ```bash
   # Copy environment files
   ./scripts/setup.sh
   ```

2. **Start Infrastructure**
   ```bash
   # Start database, cache, and message broker
   docker-compose -f docker-compose.yml -f docker-compose.dev.yml up postgres redis rabbitmq -d
   ```

3. **Run Database Migrations**
   ```bash
   export DATABASE_URL="postgres://paygateway_user:paygateway_pass@localhost:5432/paygateway_db?sslmode=disable"
   sqlx migrate run --source go-api/migrations
   ```

4. **Start Services**
   ```bash
   # Terminal 1: Start Go API
   cd go-api && go run cmd/main.go

   # Terminal 2: Start Rust Processor
   cd rust-processor && cargo run

   # Terminal 3: Start Frontend
   cd frontend && npm run dev
   ```

### 🏭 Production Mode

```bash
# Start all services in production mode
docker-compose up -d --build

# Run migrations
export DATABASE_URL="postgres://paygateway_user:paygateway_pass@localhost:5432/paygateway_db?sslmode=disable"
sqlx migrate run --source go-api/migrations
```

</details>

<details>
<summary><strong>📁 Project Structure</strong></summary>

```
payment-gateway-challenge/
├── 🌐 frontend/                 # React application
│   ├── src/
│   │   ├── components/         # Reusable UI components
│   │   ├── pages/             # Page components
│   │   ├── hooks/             # Custom React hooks
│   │   ├── services/          # API client services
│   │   └── store/             # State management
│   └── public/                # Static assets
│
├── 🚀 go-api/                   # Go HTTP API
│   ├── cmd/                   # Application entrypoints
│   ├── internal/              # Private application code
│   │   ├── account/          # Account management
│   │   ├── card/             # Card operations
│   │   ├── transaction/      # Transaction handling
│   │   └── config/           # Configuration
│   ├── migrations/           # Database migrations
│   └── docs/                 # Swagger documentation
│
├── ⚡ rust-processor/           # Rust processing engine
│   ├── src/
│   │   ├── processors/       # Transaction processors
│   │   ├── services/         # Business services
│   │   ├── models/           # Data models
│   │   ├── database/         # Database operations
│   │   └── connections/      # External connections
│   └── migrations/           # Rust migrations
│
├── 🐳 docker-compose.yml        # Production configuration
├── 🐳 docker-compose.dev.yml   # Development configuration
├── 📜 scripts/                 # Setup and utility scripts
└── 📚 docs/                    # Additional documentation
```

</details>

<details>
<summary><strong>🌟 Features</strong></summary>

### Core Features

#### 👤 Account Management
- Create and manage user accounts
- Account balance tracking
- Transaction history

#### 💳 Card Operations
- Generate virtual payment cards
- Secure tokenization system
- Card-specific transaction tracking

#### 💰 Transaction Processing
- **Purchase** - Debit transactions
- **Deposit** - Credit transactions  
- **Refund** - Reverse previous transactions
- Real-time balance updates
- Idempotency protection

#### 📊 Financial Reporting
- Detailed transaction statements
- Balance history tracking
- Card-specific filtering
- Export capabilities

### Technical Features

#### 🔒 Security
- Secure card tokenization
- Idempotent transaction processing
- Input validation and sanitization
- Error handling and logging

#### ⚡ Performance
- Async message processing
- Redis caching layer
- Connection pooling
- Optimized database queries

#### 🔧 Monitoring
- Health check endpoints
- Structured logging
- Transaction flow tracking
- Error reporting

</details>

<details>
<summary><strong>🔗 API Documentation</strong></summary>

### Swagger Documentation

When running the Go API, comprehensive API documentation is available at:

**Development:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Key Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/accounts` | Create new account |
| `GET` | `/accounts` | List all accounts |
| `POST` | `/cards` | Create new card |
| `GET` | `/cards/{accountId}` | List account cards |
| `POST` | `/transactions` | Process transaction |
| `GET` | `/transactions/{accountId}` | Get transaction history |
| `GET` | `/accounts/{accountId}/balance` | Get account balance |
| `GET` | `/health` | Health check |

### Postman Collection

A complete Postman collection is available in the frontend at:
`frontend/public/payment-gateway.postman_collection.json`

</details>

<details>
<summary><strong>🚀 Development</strong></summary>

### Running Individual Services

#### Frontend Development
```bash
cd frontend
npm install
npm run dev
# Runs on http://localhost:8081
```

#### Go API Development
```bash
cd go-api
go mod download
go run cmd/main.go
# Runs on http://localhost:8080
```

#### Rust Processor Development
```bash
cd rust-processor
cargo build
cargo run
# Connects to RabbitMQ queues
```

### Testing

#### Frontend Tests
```bash
cd frontend
npm run test
```

#### Go API Tests
```bash
cd go-api
go test ./...
```

#### Rust Processor Tests
```bash
cd rust-processor
cargo test
```

### Database Operations

#### Create Migration
```bash
# Go migrations
sqlx migrate add -r <migration_name> --source go-api/migrations

# Apply migrations
sqlx migrate run --source go-api/migrations
```

</details>

<details>
<summary><strong>🐳 Docker Operations</strong></summary>

### Available Services

- **postgres** - PostgreSQL database
- **redis** - Redis cache
- **rabbitmq** - RabbitMQ message broker
- **go-api** - Go HTTP API
- **rust-processor** - Rust processing engine
- **frontend** - React application

### Common Commands

```bash
# Start all services
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up

# Start specific service
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up rust-processor

# View logs
docker-compose logs -f rust-processor

# Stop all services
docker-compose down

# Rebuild and start
docker-compose up --build
```

### Service URLs

- **Frontend:** http://localhost:8081
- **Go API:** http://localhost:8080
- **Swagger Docs:** http://localhost:8080/swagger/index.html
- **RabbitMQ Management:** http://localhost:15672 (guest/guest)

</details>

---

## 🤝 Contributing

This is a learning project, but contributions and suggestions are welcome! Feel free to:

- Open issues for bugs or feature requests
- Submit pull requests with improvements
- Share feedback on architecture decisions
- Suggest new learning opportunities

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👨‍💻 Author

**Eduardo MG**
- GitHub: [@EduardoMG12](https://github.com/EduardoMG12)
- LinkedIn: [Eduardo MG](https://linkedin.com/in/eduardomg12)

---

<div align="center">

**⭐ If you found this project helpful, please give it a star!**

*Built with ❤️ for learning and sharing knowledge*

</div>