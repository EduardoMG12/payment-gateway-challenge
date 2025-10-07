# Processor-Rust - Processing Engine

This is the transaction processing engine of the payment gateway, developed in Rust.

## Responsibilities

- **Message Consumption:** Listens to the RabbitMQ message queue and consumes transactions for processing.
- **Business Validation:** Applies business rules, such as idempotency validation and balance verification.
- **Ledger Management:** Interacts with the PostgreSQL database to save and update transaction states.
- **Balance Calculation:** Processes transactions to calculate the updated balance and caches it in Redis.

## How to Run the Processor (Development)

1.  **Start the infrastructure:**
    ```sh
    docker-compose up -d postgres rabbitMQ redis
    ```

2.  **Start the processor in watch mode:**
    ```sh
    cargo watch -x run
    ```

The service will connect to RabbitMQ and start processing messages from the `transactions_queue` and `calculate_balance_queue` queues.

## Technologies

- **Rust**
- **Tokio:** Asynchronous runtime
- **Lapin:** RabbitMQ client for Rust
- **SQLx:** Asynchronous SQL toolkit
- **Redis:** Client for the cache
- **Serde:** Data serialization and deserialization