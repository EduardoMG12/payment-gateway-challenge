use anyhow::Result;
use sqlx::PgPool;

use crate::models::{QueueTransaction, TransactionStatus, TransactionType};
use crate::repository::{
    account_reposistory::AccountRepository, transaction_repository::TransactionRepository,
};
use crate::services;

pub async fn process_transaction(pool: &PgPool, tx: QueueTransaction) -> Result<()> {
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
            .await?;
    }

    match tx
        .transaction_type
        .parse::<TransactionType>()
        .map_err(anyhow::Error::msg)?
    {
        TransactionType::DEPOSIT => {
            services::deposit_service::process_deposit(tx);
            return Ok(());
        }
        TransactionType::PURCHASE => {
            services::purchase_service::process_purchase(tx);
            return Ok(());
        }
        TransactionType::REFUND => {
            services::refund_service::process_refund(tx);
            return Ok(());
        }
    }

    // Annotations to i thunking about nexts steps
    // case i cant process the transaction throw message to queue again
    // case send message to queue again 3 times, update status to ERROR
    // case transcions is proccess with success update status to APPROVED and make request to get new balance and send value to redis
    // thinking about whats difference and CHARGE and DEPOSIT and how to process this two types
    // thinking about how can i make the refund process
    // thinking about my way is the best way to do this
    // thinking about my code is clean and easy to understand, my imporsts are organized
    // thinking about how can i make a documentation for my rust code

    let balance = account_repo.get_balance(tx.account_id).await?;
    println!("Novo saldo da conta {}: {}", tx.account_id, balance);

    Ok(())
}
