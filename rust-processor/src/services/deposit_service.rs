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
        eprintln!(
            "Conta {} não encontrada. Transação {} rejeitada.",
            tx.account_id, tx.id
        );
        return Ok(());
    }

    let maybe_tx = transaction_repo.find_by_id(tx.id).await?;
    if maybe_tx.is_none() {
        eprintln!(
            "ERRO CRÍTICO: Transação {} não encontrada no banco de dados!",
            tx.id
        );
        return Err(anyhow!("Transação {} não encontrada no DB.", tx.id));
    }

    let existing_tx = maybe_tx.unwrap();
    if existing_tx.status != TransactionStatus::PENDING {
        println!("Transação {} já processada. Ignorando.", tx.id);
        return Ok(());
    }

    // This logic simple register money input but in real system here should comuniquete with external payment provider
    println!(
        "Deposit para a conta {} processado com sucesso.",
        tx.account_id
    );

    transaction_repo
        .update_status(tx.id, TransactionStatus::APPROVED)
        .await?;
    println!("Status da transação {} atualizado para APPROVED.", tx.id);

    // I thinking the next step is create redis and implement update Balance on redis after make transaction
    // let new_balance = account_repo.get_balance(tx.account_id).await?;
    // redis_service::update_balance(tx.account_id, new_balance).await?;
    // println!("Saldo da conta {} atualizado no Redis para: {}", tx.account_id, new_balance);

    // send for other queue to my frontend listen?
    Ok(())
}
