use crate::connections::CacheRepository;
use crate::models::BalanceRequest;
use crate::repository::{TTransactionRepository, account_repository};
use anyhow::Result;

pub async fn process_balance_request(
    transaction_repository: &impl TTransactionRepository,
    cache_repository: &CacheRepository,
    req: BalanceRequest,
) -> Result<()> {
    println!(
        "📥 Balance calculation request received for account: {}",
        req.account_id
    );

    let balance = transaction_repository.get_balance(req.account_id).await?;
    println!(
        "💰 Calculated balance for account {}: {} cents",
        req.account_id, balance
    );

    cache_repository
        .set_balance(req.account_id, balance)
        .await?;
    Ok(())
}
