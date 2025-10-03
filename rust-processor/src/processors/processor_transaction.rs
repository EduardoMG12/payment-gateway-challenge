use std::f32::consts::E;

use anyhow::Result;
use sqlx::PgPool;

use crate::connections::CacheRepository;
use crate::models::{BalanceRequest, QueueTransaction, TransactionStatus, TransactionType};
use crate::repository::{AccountRepository, TTransactionRepository, TransactionRepository};
use crate::services;

pub async fn process_transaction(
    pool: &PgPool,
    cache_repository: &CacheRepository,
    tx: QueueTransaction,
) -> Result<()> {
    let account_repo = AccountRepository::new(pool);
    let transaction_repo = TransactionRepository::new(pool);

    println!("Processando transação: {{");
    println!("  ID: {}", tx.id);
    println!("  Account ID: {}", tx.account_id);
    println!("  Amount (cents): {}", tx.amount_cents);
    println!("  Type: {}", tx.transaction_type);
    println!(
        "  Refund Transaction ID: {}",
        tx.refund_transaction_id.value
    );
    println!("  Status: {}", tx.status);
    println!("}}");

    if tx.amount_cents <= 0 {
        transaction_repo
            .update_status(tx.id, TransactionStatus::REJECTED)
            .await?
    }

    match tx.transaction_type.parse::<TransactionType>() {
        Ok(TransactionType::DEPOSIT) => {
            services::deposit_service::process_deposit(
                &transaction_repo,
                &account_repo,
                tx.clone(),
            )
            .await?;
        }
        Ok(TransactionType::PURCHASE) => {
            services::purchase_service::process_purchase(tx.clone());
        }
        Ok(TransactionType::REFUND) => {
            services::refund_service::process_refund(tx.clone());
        }
        Err(e) => {
            transaction_repo
                .update_status(tx.id, TransactionStatus::REJECTED)
                .await?;
            return Err(anyhow::anyhow!("Invalid transaction type: {}", e));
        }
    }

    let balance_request = BalanceRequest {
        account_id: tx.account_id,
    };

    crate::processors::processor_balance::process_balance_request(
        &transaction_repo,
        &cache_repository,
        balance_request,
    )
    .await?;

    Ok(())
    // Annotations to i thunking about nexts steps
    // case i cant process the transaction throw message to queue again
    // case send message to queue again 3 times, update status to ERROR
    // thinking about my way is the best way to do this
    // thinking about my code is clean and easy to understand, my imporsts are organized
    // thinking about how can i make a documentation for my rust code
}
