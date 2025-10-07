# Fictitious Payment Gateway

This is a learning project focused on **Rust** and **Go**, where I build a simplified payment gateway. The goal is not to create a market-ready product, but rather to apply software engineering and systems architecture concepts in a practical **fintech** domain.

The project is a **monorepo** with three main services:
- A simple **frontend** in React to simulate client interaction.
- A **Go API** to receive transaction requests.
- A **Rust processing engine** to execute the core business logic.

## System Architecture

The project is divided into three main components that communicate asynchronously using RabbitMQ.

A[Frontend] -->|1. Initiate Transaction| B(Go API);
B -->|2. Publish to queue| C[RabbitMQ];
C -->|3. Process Transaction| D(Rust Service);
D -->|4. Save to Ledger| E[PostgreSQL];
D -->|5. Update Status| E;
E --> D;
B -->|6. Polling for status| A;

## Technologies Used

* **Frontend:** React, TypeScript, Vite, Tailwind CSS
* **API:** Go, Gin, Swagger
* **Processor:** Rust, Tokio, Lapin
* **Messaging:** RabbitMQ
* **Database:** PostgreSQL
* **Cache:** Redis
* **Containerization:** Docker, Docker Compose


## How to Run the Project

### Prerequisites

- Docker and Docker Compose
- `sqlx-cli` installed (`cargo install sqlx-cli`)

### 1. Quick Setup (Development)

To start the project in development mode, follow the steps below:

1.  **Make the setup script executable:**
    ```sh
    chmod +x scripts/setup.sh
    ```

2.  **Run the setup script:**
    ```sh
    ./scripts/setup.sh
    ```
   
This command will copy the `.env` environment files and install the Node.js dependencies.

3.  **Start the Docker containers:**
    ```sh
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build -d
    ```

4.  **Run the database migrations:**
    ```sh
    export DATABASE_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable"
    sqlx migrate run --source go-api/migrations
    ```

### 2. Running in Production

To run the project in production mode:

1.  **Copy the production environment file:**
    ```sh
    ./scripts/setup.sh
    ```

2.  **Start the Docker containers:**
    ```sh
    docker-compose up -d --build
    ```

3.  **Run the database migrations:**
    ```sh
    export DATABASE_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable"
    sqlx migrate run --source go-api/migrations
    ```
## Monorepo Structure

* `/frontend`: React application.
* `/go-api`: Go API.
* `/processor-rust`: Rust processing service.
* `/scripts`: Setup and utility scripts.