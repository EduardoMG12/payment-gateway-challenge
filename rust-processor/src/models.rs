use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct Transaction {
    pub id: String,
    pub account_id: String,
    pub amount_cents: i64,
    pub r#type: String, // "PURCHASE" | "DEPOSIT" in the future i will implement "REFUND"
}
