// use uuid::Uuid;
// use sqlx::FromRow;
// use chrono::{DateTime, Utc};
use serde::Deserialize;
use serde::Serialize;
use std::str::FromStr;

#[derive(Debug, Clone, Copy, Deserialize)]
pub enum TransactionStatus {
    PENDING,
    APPROVED,
    REJECTED,
    ERROR,
}

impl TransactionStatus {
    // pub fn from_str(s: &str) -> Option<Self> {
    //     match s {
    //         "PENDING" => Some(TransactionStatus::PENDING),
    //         "APPROVED" => Some(TransactionStatus::APPROVED),
    //         "REJECTED" => Some(TransactionStatus::REJECTED),
    //         "ERROR" => Some(TransactionStatus::ERROR),
    //         _ => None,
    //     }
    // }

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

// #[derive(Debug, FromRow)]
// pub struct DbTransaction {
//     pub id: Uuid,
//     pub account_id: Uuid,
//     pub card_id: Option<Uuid>,

//     pub amount_cents: i64,
//     #[sqlx(rename = "type")]
//     pub transaction_type: String,
//     pub idempotency_key: String,
//     pub created_at: DateTime<Utc>,
// }

// #[derive(Debug, FromRow)]
// pub struct Account {
//     pub id: Uuid,
//     pub balance_cents: i64,
// }
