mod config;
mod connections;
mod models;
mod processors;
mod repository;
mod services;

use std::sync::Arc;

use crate::{
    models::{BalanceRequest, QueueTransaction},
    processors::{processor_balance, processor_transaction},
};

use anyhow::Result;

#[tokio::main]
async fn main() -> Result<()> {
    dotenv::dotenv().ok();
    let config = config::Config::load()?;

    let pool = connections::db::setup_sqlx_pool(&config.database_url()).await?;

    let amqp_addr = config.rabbitmq_amqp_addr();
    let channel = connections::rabbitmq::create_channel(&amqp_addr).await?;

    let pool = Arc::new(pool);

    connections::rabbitmq::consume_queue(channel, "transactions_queue", {
        let pool = Arc::clone(&pool);
        move |msg: String| {
            let pool = Arc::clone(&pool);
            async move {
                match serde_json::from_str::<QueueTransaction>(&msg) {
                    Ok(tx) => {
                        if let Err(err) =
                            processor_transaction::process_transaction(&pool, tx).await
                        {
                            eprintln!("❌ Error processing transaction: {:?}", err);
                        }
                    }
                    Err(err) => eprintln!("❌ Error deserializing message: {:?}", err),
                }
            }
        }
    })
    .await?;

    let channel_balance = connections::rabbitmq::create_channel(&amqp_addr).await?;
    let pool_balance = Arc::clone(&pool);
    tokio::spawn(async move {
        let handler = move |msg: String| {
            let pool = Arc::clone(&pool_balance);
            async move {
                match serde_json::from_str::<BalanceRequest>(&msg) {
                    Ok(req) => {
                        if let Err(err) =
                            processor_balance::process_balance_request(&pool, req).await
                        {
                            eprintln!("❌ Error to process balance: {:?}", err);
                        }
                    }
                    Err(err) => eprintln!("❌ Error deserializing BalanceRequest: {:?}", err),
                }
            }
        };
        if let Err(e) = connections::rabbitmq::consume_queue(
            channel_balance,
            "calculate_balance_queue",
            handler,
        )
        .await
        {
            eprintln!("Error in balance consumer: {}", e);
        }
    });

    println!(
        "🚀 Processor begin listening and waiting for messages on queues: 'transactions_queue' and 'calculate_balance_queue'"
    );

    println!("Connected to: {}", amqp_addr);

    loop {
        tokio::time::sleep(std::time::Duration::from_secs(60)).await;
    }
}
