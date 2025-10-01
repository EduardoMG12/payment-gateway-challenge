use crate::models::QueueTransaction;

pub fn process_purchase(tx: QueueTransaction) {
    println!("Processing purchase transaction: {:?}", tx);
}
