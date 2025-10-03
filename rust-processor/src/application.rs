use crate::{
    config::Config,
    connections::{self, CacheRepository},
    models::{BalanceRequest, QueueTransaction},
    processors::{processor_balance, processor_transaction},
    repository::{AccountRepository, TransactionRepository},
};
use anyhow::Result;
use lapin::Channel;
use sqlx::PgPool;
use std::sync::Arc;
use tokio::task::JoinHandle;
pub struct Application {
    db_pool: Arc<PgPool>,
    cache_repo: Arc<CacheRepository>,
    transaction_channel: Arc<Channel>,
    balance_channel: Arc<Channel>,
}

impl Application {
    pub async fn build(config: &Config) -> Result<Self> {
        let db_pool = Arc::new(connections::db::setup_sqlx_pool(&config.database_url()).await?);

        let cache_repo = Arc::new(connections::CacheRepository::new(&config.redis_url())?);

        let amqp_conn =
            connections::rabbitmq::create_connection(&config.rabbitmq_amqp_addr()).await?;

        let transaction_channel =
            Arc::new(connections::rabbitmq::create_channel(&amqp_conn).await?);
        let balance_channel = Arc::new(connections::rabbitmq::create_channel(&amqp_conn).await?);

        println!("✅ Connected to Database and RabbitMQ successfully.");

        Ok(Self {
            db_pool,
            cache_repo,
            transaction_channel,
            balance_channel,
        })
    }

    pub async fn run_with_loop(self) -> Result<()> {
        println!(
            "🚀 Processor starting to listen on queues: 'transactions_queue' and 'calculate_balance_queue'"
        );

        self.run_transaction_consumer();
        self.run_balance_consumer();

        loop {
            tokio::time::sleep(std::time::Duration::from_secs(60)).await;
        }
    }

    fn run_transaction_consumer(&self) -> JoinHandle<()> {
        let pool = Arc::clone(&self.db_pool);
        let channel = Arc::clone(&self.transaction_channel);
        let cache_repo = Arc::clone(&self.cache_repo);
        tokio::spawn(async move {
            let handler = move |msg: String| {
                let pool = Arc::clone(&pool);
                let cache_repo_clone = Arc::clone(&cache_repo);
                async move {
                    match serde_json::from_str::<QueueTransaction>(&msg) {
                        Ok(tx) => {
                            if let Err(err) = processor_transaction::process_transaction(
                                &pool,
                                &cache_repo_clone,
                                tx,
                            )
                            .await
                            {
                                eprintln!("❌ Error processing transaction: {:?}", err);
                            }
                        }
                        Err(err) => {
                            eprintln!("❌ Error deserializing transaction message: {:?}", err)
                        }
                    }
                }
            };

            if let Err(e) =
                connections::rabbitmq::consume_queue(channel, "transactions_queue", handler).await
            {
                eprintln!("FATAL: Transaction consumer failed: {}", e);
            }
        })
    }

    fn run_balance_consumer(&self) -> JoinHandle<()> {
        let pool = Arc::clone(&self.db_pool);
        let channel = Arc::clone(&self.balance_channel);
        let cache_repo = Arc::clone(&self.cache_repo);

        tokio::spawn(async move {
            let handler = move |msg: String| {
                let pool = Arc::clone(&pool);
                let cache_repo_clone = Arc::clone(&cache_repo);
                async move {
                    match serde_json::from_str::<BalanceRequest>(&msg) {
                        Ok(req) => {
                            let transaction_repo = TransactionRepository::new(&pool);
                            let account_repo = AccountRepository::new(&pool);
                            if let Err(err) = processor_balance::process_balance_request(
                                &account_repo,
                                &transaction_repo,
                                &cache_repo_clone,
                                req,
                            )
                            .await
                            {
                                eprintln!("❌ Error processing balance request: {:?}", err);
                            }
                        }
                        Err(err) => eprintln!("❌ Error deserializing balance request: {:?}", err),
                    }
                }
            };
            if let Err(e) =
                connections::rabbitmq::consume_queue(channel, "calculate_balance_queue", handler)
                    .await
            {
                eprintln!("FATAL: Balance consumer failed: {}", e);
            }
        })
    }
}
