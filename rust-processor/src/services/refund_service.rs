use crate::models::QueueTransaction;

pub fn process_refund(tx: QueueTransaction) {
    println!("Processando reembolso: {:?}", tx);
}
