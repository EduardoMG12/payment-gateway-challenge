use anyhow::Result;
use sqlx::{PgPool, Pool, Postgres};

pub type DbPool = Pool<Postgres>;

pub async fn setup_sqlx_pool(database_url: &str) -> Result<DbPool> {
    let pool = PgPool::connect(database_url)
        .await
        .map_err(|e| anyhow::anyhow!("Erro to conect data base: {}", e))?;

    sqlx::query("SELECT 1")
        .execute(&pool)
        .await
        .map_err(|e| anyhow::anyhow!("Erro to test connection data base: {}", e))?;

    println!("✅ Connection with data-base is stabelised with succes!");

    Ok(pool)
}
