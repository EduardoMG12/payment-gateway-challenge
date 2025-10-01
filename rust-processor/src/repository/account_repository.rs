use anyhow::Result;
use sqlx::PgPool;
use uuid::Uuid;

pub struct AccountRepository<'a> {
    pool: &'a PgPool,
}

impl<'a> AccountRepository<'a> {
    pub fn new(pool: &'a PgPool) -> Self {
        Self { pool }
    }

    pub async fn get_balance(&self, account_id: Uuid) -> Result<i64> {
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
}
