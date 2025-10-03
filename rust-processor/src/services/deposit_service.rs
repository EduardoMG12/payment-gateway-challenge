use crate::models::{QueueTransaction, TransactionStatus};
use crate::repository::{TAccountRepository, TTransactionRepository};
use anyhow::{Result, anyhow};

pub async fn process_deposit(
    transaction_repo: &impl TTransactionRepository,
    account_repo: &impl TAccountRepository,
    tx: QueueTransaction,
) -> Result<()> {
    let account = account_repo.find_by_id(tx.account_id).await?;
    if account.is_none() {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;
        return Ok(());
    }

    let maybe_tx = transaction_repo.find_by_id(tx.id).await?;
    if maybe_tx.is_none() {
        return Err(anyhow!("Transaction {} not found in DB.", tx.id));
    }

    let existing_tx = maybe_tx.unwrap();
    if existing_tx.status != TransactionStatus::PENDING {
        println!("Transaction {} already processed. Ignoring.", tx.id);
        return Ok(());
    }

    transaction_repo
        .update_status(tx.id, TransactionStatus::APPROVED)
        .await?;
    println!("Status da transação {} atualizado para APPROVED.", tx.id);

    Ok(())
}
