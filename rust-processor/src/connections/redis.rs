use anyhow::Result;
use redis::AsyncCommands;
use uuid::Uuid;

const BALANCE_EXPIRATION_SECONDS: u64 = 24 * 60 * 60;

#[derive(Clone)]
pub struct CacheRepository {
    client: redis::Client,
}

impl CacheRepository {
    pub fn new(redis_url: &str) -> Result<Self> {
        let client = redis::Client::open(redis_url)?;
        Ok(Self { client })
    }

    pub async fn set_balance(&self, account_id: Uuid, balance: i64) -> Result<()> {
        let mut conn = self.client.get_multiplexed_tokio_connection().await?;

        let cache_key = format!("balance:{}", account_id);

        let _: () = conn
            .set_ex(cache_key, balance, BALANCE_EXPIRATION_SECONDS)
            .await?;

        println!(
            "   CACHE: Balance account {} updated to {} in Redis (expires in 24h)",
            account_id, balance
        );

        Ok(())
    }
}
