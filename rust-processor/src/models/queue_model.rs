use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct QueueCardId {
    #[serde(rename = "String")]
    pub value: String,
    #[serde(rename = "Valid")]
    pub is_valid: bool,
}

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct RefundTransactionId {
    #[serde(rename = "String")]
    pub value: String,
    #[serde(rename = "Valid")]
    pub is_valid: bool,
}

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct QueueTransaction {
    pub id: Uuid,
    pub account_id: Uuid,
    pub card_id: QueueCardId,
    pub amount_cents: i64,
    pub status: String,
    #[serde(rename = "type")]
    pub transaction_type: String,
    pub refund_transaction_id: RefundTransactionId,
    pub idempotency_key: String,
    pub created_at: DateTime<Utc>,
    #[serde(default)]
    pub retry_count: i32,
}
