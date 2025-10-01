use chrono::{DateTime, Utc};
use sqlx::FromRow;
use uuid::Uuid;

#[derive(Debug, FromRow)]
pub struct DbAccount {
    pub id: Uuid,
    pub username: String,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

#[derive(Debug, FromRow)]
pub struct Account {
    pub id: Uuid,
    pub balance_cents: i64,
}
