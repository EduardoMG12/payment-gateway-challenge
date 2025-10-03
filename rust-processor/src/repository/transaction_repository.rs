use crate::models::{DbTransaction, TransactionStatus};
use anyhow::Result;
use sqlx::PgPool;
use uuid::Uuid;

use async_trait::async_trait;

#[async_trait]
pub trait TTransactionRepository {
    async fn find_by_id(&self, tx_id: Uuid) -> Result<Option<DbTransaction>>;
    async fn update_status(&self, tx_id: Uuid, status: TransactionStatus) -> Result<()>;
    async fn update_refund_transaction_id(
        &self,
        tx_id: Uuid,
        refund_tx_id: Uuid,
        status: TransactionStatus,
    ) -> Result<()>;
    async fn get_balance(&self, account_id: Uuid) -> Result<i64>;
    async fn has_been_refunded(&self, original_tx_id: Uuid) -> Result<bool>;
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

    async fn has_been_refunded(&self, original_tx_id: Uuid) -> Result<bool> {
        let exists: (bool,) = sqlx::query_as(
            r#"
            SELECT EXISTS (
                SELECT 1 FROM transactions
                WHERE refund_transaction_id = $1 AND status = 'APPROVED'
            )
            "#,
        )
        .bind(original_tx_id)
        .fetch_one(self.pool)
        .await?;

        Ok(exists.0)
    }

    async fn update_refund_transaction_id(
        &self,
        tx_id: Uuid,
        refund_tx_id: Uuid,
        status: TransactionStatus,
    ) -> Result<()> {
        sqlx::query(
            r#"
            UPDATE transactions
            SET refund_transaction_id  = $1, status = $2, amount_cents = (SELECT amount_cents FROM transactions WHERE id = $1)
            WHERE id = $3
            "#,
        )
        .bind(refund_tx_id)
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

    async fn get_balance(&self, account_id: Uuid) -> Result<i64> {
        let row: (Option<i64>,) = sqlx::query_as(
            r#"
            SELECT
                SUM(
                    CASE
                        WHEN t1.type = 'DEPOSIT' THEN t1.amount_cents
                        WHEN t1.type = 'PURCHASE' THEN -t1.amount_cents
                        WHEN t1.type = 'REFUND' THEN
                            CASE t_orig.type
                                WHEN 'DEPOSIT' THEN -t1.amount_cents
                                WHEN 'PURCHASE' THEN t1.amount_cents
                                ELSE 0
                            END
                        ELSE 0
                    END
                )::BIGINT
            FROM
                transactions AS t1
            LEFT JOIN
                transactions AS t_orig ON t1.refund_transaction_id = t_orig.id
            WHERE
                t1.account_id = $1
            AND
                t1.status = 'APPROVED'
            "#,
        )
        .bind(account_id)
        .fetch_one(self.pool)
        .await?;
        Ok(row.0.unwrap_or(0))
    }
}
