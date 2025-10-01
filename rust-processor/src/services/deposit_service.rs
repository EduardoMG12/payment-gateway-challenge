use crate::models::QueueTransaction;

pub fn process_deposit(tx: QueueTransaction) {
    println!("Processing deposit transaction: {:?}", tx);
}
