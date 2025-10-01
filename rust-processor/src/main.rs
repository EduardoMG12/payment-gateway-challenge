mod config;
mod db;
mod models;
mod processor;
mod rabbitmq;
mod repository;
mod services;

use std::sync::Arc;

use crate::models::QueueTransaction;

use anyhow::Result;

#[tokio::main]
async fn main() -> Result<()> {
    dotenv::dotenv().ok();
    let config = config::Config::load()?;

    let pool = db::setup_sqlx_pool(&config.database_url()).await?;

    let amqp_addr = config.rabbitmq_amqp_addr();
    let channel = rabbitmq::create_channel(&amqp_addr).await?;

    let pool = Arc::new(pool);

    rabbitmq::consume_queue(channel, "transactions_queue", {
        let pool = Arc::clone(&pool);
        move |msg: String| {
            let pool = Arc::clone(&pool);
            async move {
                match serde_json::from_str::<QueueTransaction>(&msg) {
                    Ok(tx) => {
                        if let Err(err) = processor::process_transaction(&pool, tx).await {
                            eprintln!("❌ Erro no processamento: {:?}", err);
                        }
                    }
                    Err(err) => eprintln!("❌ Erro ao deserializar mensagem: {:?}", err),
                }
            }
        }
    })
    .await?;

    println!("Connected to: {}", amqp_addr);

    loop {
        tokio::time::sleep(std::time::Duration::from_secs(60)).await;
    }
}
