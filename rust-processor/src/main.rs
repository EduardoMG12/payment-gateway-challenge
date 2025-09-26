mod config;
mod db;
mod processor;
mod rabbitmq;

use anyhow::Result;

#[tokio::main]
async fn main() -> Result<()> {
    dotenv::dotenv().ok();
    let config = config::Config::load()?;

    let pool = db::setup_sqlx_pool(&config.database_url()).await?;

    let amqp_addr = config.rabbitmq_amqp_addr();
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
