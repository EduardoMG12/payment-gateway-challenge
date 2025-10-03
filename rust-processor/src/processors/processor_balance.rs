use crate::connections::CacheRepository;
use crate::models::BalanceRequest;
use crate::repository::{TAccountRepository, TTransactionRepository};
use anyhow::Result;

pub async fn process_balance_request(
    account_repository: &impl TAccountRepository,
    transaction_repository: &impl TTransactionRepository,
    cache_repository: &CacheRepository,
    req: BalanceRequest,
) -> Result<()> {
    let account = account_repository.find_by_id(req.account_id).await?;
    if account.is_none() {
        return Ok(());
    }

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
