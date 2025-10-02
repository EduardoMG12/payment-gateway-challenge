use crate::models::BalanceRequest;
use crate::repository::{AccountRepository, TAccountRepository};
use anyhow::Result;
use sqlx::PgPool;

pub async fn process_balance_request(pool: &PgPool, req: BalanceRequest) -> Result<()> {
    println!(
        "📥 Balance calculation request received for account: {}",
        req.account_id
    );

    // i need to connect my redis and sum all transactions to get the balance and set on redis
    // and i need remember split my listeners and call function on main
    Ok(())
}
