mod processor;
mod rabbitmq;

use anyhow::Result;

#[tokio::main]
async fn main() -> Result<()> {
    dotenv::dotenv().ok();

    let rabbit_user = std::env::var("RABBITMQ_DEFAULT_USER")
        .expect("Variable RABBITMQ_DEFAULT_USER not found. Verify your file .env");
    let rabbit_pass = std::env::var("RABBITMQ_DEFAULT_PASS")
        .expect("Variable RABBITMQ_DEFAULT_PASS not found. Verify your file .env");
    let rabbit_host = std::env::var("RABBITMQ_HOST").unwrap_or("localhost".into());
    let rabbit_port = std::env::var("RABBITMQ_PORT").unwrap_or("5672".into());

    let amqp_addr = format!(
        "amqp://{}:{}@{}:{}/%2f",
        rabbit_user, rabbit_pass, rabbit_host, rabbit_port
    );

    let channel = rabbitmq::create_channel(&amqp_addr).await?;

    rabbitmq::consume_queue(
        channel,
        "transactions_queue",
        processor::process_transaction,
    )
    .await?;

    println!("Connected to: {}", amqp_addr);

    loop {
        tokio::time::sleep(std::time::Duration::from_secs(60)).await;
    }
}
