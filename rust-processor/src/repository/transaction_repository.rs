use crate::models::TransactionStatus;
use anyhow::Result;
use sqlx::PgPool;
use uuid::Uuid;

pub struct TransactionRepository<'a> {
    pool: &'a PgPool,
}

impl<'a> TransactionRepository<'a> {
    pub fn new(pool: &'a PgPool) -> Self {
        Self { pool }
    }

    pub async fn update_status(&self, tx_id: Uuid, status: TransactionStatus) -> Result<()> {
        sqlx::query(
            r#"
            UPDATE transactions
            SET status = $1
            WHERE id = $2
            "#,
        )
        .bind(status.as_str())
        .bind(tx_id)
        .execute(self.pool)
        .await?;

        Ok(())
    }
}
