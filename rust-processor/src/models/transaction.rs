use chrono::DateTime;
use chrono::Utc;
use serde::Deserialize;
use serde::Serialize;
use sqlx::FromRow;
use std::str::FromStr;
use uuid::Uuid;

#[derive(Debug, sqlx::Type, PartialEq)]
#[sqlx(type_name = "VARCHAR")]
pub enum TransactionStatus {
    PENDING,
    APPROVED,
    REJECTED,
    ERROR,
}

impl TransactionStatus {
    pub fn as_str(&self) -> &str {
        match self {
            TransactionStatus::PENDING => "PENDING",
            TransactionStatus::APPROVED => "APPROVED",
            TransactionStatus::REJECTED => "REJECTED",
            TransactionStatus::ERROR => "ERROR",
        }
    }
}

#[derive(Debug, Serialize, Deserialize)]
pub enum TransactionType {
    DEPOSIT,
    PURCHASE,
    REFUND,
}

impl FromStr for TransactionType {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.to_uppercase().as_str() {
            "DEPOSIT" => Ok(TransactionType::DEPOSIT),
            "PURCHASE" => Ok(TransactionType::PURCHASE),
            "REFUND" => Ok(TransactionType::REFUND),
            _ => Err(format!("Invalid transaction type: {}", s)),
        }
    }
}

#[derive(Debug, FromRow)]
pub struct DbTransaction {
    pub id: Uuid,
    pub account_id: Uuid,
    pub amount_cents: i64,
    #[sqlx(rename = "type")]
    pub transaction_type: String,
    pub refund_transaction_id: Option<Uuid>,
    pub status: TransactionStatus,
    pub created_at: DateTime<Utc>,
}

#[derive(Debug, Deserialize)]
pub struct BalanceRequest {
    pub account_id: Uuid,
}
