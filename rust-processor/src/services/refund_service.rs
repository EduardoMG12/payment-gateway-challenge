use anyhow::{Result, anyhow};
use uuid::Uuid;

use crate::{
    models::{QueueTransaction, TransactionStatus},
    repository::{TAccountRepository, TTransactionRepository},
};

pub async fn process_refund(
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
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;

        return Err(anyhow!("Transaction {} not found in DB.", tx.id));
    }

    let existing_tx = maybe_tx.unwrap();
    if existing_tx.status != TransactionStatus::PENDING {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;

        return Ok(());
    }
    let refund_uuid: Uuid;

    if !tx.refund_transaction_id.is_valid {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;
        return Err(anyhow!(
            "Refund transaction ID is required for REFUND type."
        ));
    }

    match Uuid::parse_str(&tx.refund_transaction_id.value) {
        Ok(id) => {
            refund_uuid = id;
        }
        Err(_) => {
            transaction_repo
                .update_status(tx.id, TransactionStatus::REJECTED)
                .await?;
            return Err(anyhow!(
                "Refund transaction ID is marked as valid, but value is not a valid UUID: {}",
                tx.refund_transaction_id.value
            ));
        }
    }

    let maybe_refund_tx = transaction_repo.find_by_id(refund_uuid).await?;
    if maybe_refund_tx.is_none() {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;
        return Err(anyhow!(
            "Refund transaction {} not found in DB.",
            refund_uuid
        ));
    }

    let existing_refund_tx = maybe_refund_tx.unwrap();

    let already_refunded = transaction_repo.has_been_refunded(refund_uuid).await?;
    if already_refunded {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;
        return Err(anyhow!(
            "Transaction {} has already been refunded.",
            refund_uuid
        ));
    }

    if existing_refund_tx.status != TransactionStatus::APPROVED {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;

        return Ok(());
    }

    if existing_refund_tx.account_id != tx.account_id {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;
        return Ok(());
    }

    if existing_refund_tx.transaction_type == "REFUND" {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?;
        return Ok(());
    }

    transaction_repo
        .update_refund_transaction_id(tx.id, refund_uuid, TransactionStatus::APPROVED)
        .await?;
    Ok(())
}
