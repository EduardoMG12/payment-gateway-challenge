# Go API - Payment Gateway

This is the Go API responsible for being the entry point of the payment system.

## Responsibilities

- **Account and Card Management:** Handles the creation of new accounts, generating fictitious cards and initial balances.
- **Tokenization Simulation:** Acts as a tokenization service, receiving fictitious card data and generating a secure hash for internal use.
- **Asynchronous Communication:** Publishes messages to the queue (RabbitMQ) so that the Rust processor can work in a decoupled manner.
- **Status Inquiry:** Provides endpoints for the frontend to query the status of a transaction.

## API Endpoints

The complete API documentation is available via Swagger. After starting the API, access:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Main endpoints:

- `POST /accounts`: Creates a new account.
- `GET /accounts`: Lists existing accounts.
- `POST /cards`: Creates a new card for an account.
- `GET /cards/{accountId}`: Lists the cards of an account.
- `POST /transactions`: Creates a new transaction.
- `GET /transactions/{accountId}`: Lists the transactions of an account.