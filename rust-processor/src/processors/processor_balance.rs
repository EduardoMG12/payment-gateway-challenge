use crate::models::BalanceRequest;
use crate::repository::{
    AccountRepository, TAccountRepository, TTransactionRepository, TransactionRepository,
};
use anyhow::Result;
use sqlx::PgPool;

pub async fn process_balance_request(pool: &PgPool, req: BalanceRequest) -> Result<()> {
    println!(
        "📥 Balance calculation request received for account: {}",
        req.account_id
    );

    let transaction_repository = TransactionRepository::new(pool);

    let balance = transaction_repository.get_balance(req.account_id).await?;
    println!(
        "💰 Calculated balance for account {}: {} cents",
        req.account_id, balance
    );
    // i need to connect my redis and sum all transactions to get the balance and set on redis
    // and i need remember split my listeners and call function on main
    Ok(())
}
