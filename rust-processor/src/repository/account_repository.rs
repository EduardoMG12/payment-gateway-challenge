use anyhow::Result;
use sqlx::PgPool;
use uuid::Uuid;

use async_trait::async_trait;

use crate::models::DbAccount;

#[async_trait]
pub trait TAccountRepository {
    async fn get_balance(&self, account_id: Uuid) -> Result<i64>;
    async fn find_by_id(&self, account_id: Uuid) -> Result<Option<DbAccount>>;
}

pub struct AccountRepository<'a> {
    pool: &'a PgPool,
}

impl<'a> AccountRepository<'a> {
    pub fn new(pool: &'a PgPool) -> Self {
        Self { pool }
    }
}

#[async_trait]
impl TAccountRepository for AccountRepository<'_> {
    async fn get_balance(&self, account_id: Uuid) -> Result<i64> {
        let row: (i64,) = sqlx::query_as(
            r#"
            SELECT COALESCE(SUM(
                CASE 
                    WHEN amount_cents >= 0 THEN amount_cents
                    ELSE amount_cents
                END
            )::BIGINT, 0)
            FROM transactions
            WHERE account_id = $1
              AND status = 'APPROVED'
            "#,
        )
        .bind(account_id)
        .fetch_one(self.pool)
        .await?;

        Ok(row.0)
    }

    async fn find_by_id(&self, tx_id: Uuid) -> Result<Option<DbAccount>> {
        let maybe_account = sqlx::query_as::<_, DbAccount>(
            r#"
            SELECT id, username, created_at, updated_at
            FROM accounts
            WHERE id = $1
            "#,
        )
        .bind(tx_id)
        .fetch_optional(self.pool)
        .await?;

        Ok(maybe_account)
    }
}
