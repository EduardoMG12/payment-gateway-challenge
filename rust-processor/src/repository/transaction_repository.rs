use crate::models::{DbTransaction, TransactionStatus};
use anyhow::Result;
use sqlx::PgPool;
use uuid::Uuid;

use async_trait::async_trait;

#[async_trait]
pub trait TTransactionRepository {
    async fn find_by_id(&self, tx_id: Uuid) -> Result<Option<DbTransaction>>;
    async fn update_status(&self, tx_id: Uuid, status: TransactionStatus) -> Result<()>;
}

pub struct TransactionRepository<'a> {
    pool: &'a PgPool,
}

impl<'a> TransactionRepository<'a> {
    pub fn new(pool: &'a PgPool) -> Self {
        Self { pool }
    }
}

#[async_trait]
impl TTransactionRepository for TransactionRepository<'_> {
    async fn update_status(&self, tx_id: Uuid, status: TransactionStatus) -> Result<()> {
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

    async fn find_by_id(&self, tx_id: Uuid) -> Result<Option<DbTransaction>> {
        let maybe_transaction = sqlx::query_as::<_, DbTransaction>(
            r#"
            SELECT id, account_id, amount_cents, "type", refund_transaction_id, status, created_at
            FROM transactions
            WHERE id = $1
            "#,
        )
        .bind(tx_id)
        .fetch_optional(self.pool)
        .await?;

        Ok(maybe_transaction)
    }
}
