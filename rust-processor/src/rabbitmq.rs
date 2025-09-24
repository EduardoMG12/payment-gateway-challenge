use anyhow::Result;
use futures_lite::stream::StreamExt;
use lapin::{
    Channel, Connection, ConnectionProperties,
    options::{BasicAckOptions, BasicConsumeOptions, QueueDeclareOptions},
    types::FieldTable,
};

pub async fn create_channel(amqp_addr: &str) -> Result<Channel> {
    let conn = Connection::connect(amqp_addr, ConnectionProperties::default()).await?;
    println!("✅ Conected to RabbitMQ!");
    let channel = conn.create_channel().await?;
    Ok(channel)
}

pub async fn consume_queue<F>(channel: Channel, queue_name: &str, handler: F) -> Result<()>
where
    F: Fn(String) + Send + Sync + 'static,
{
    let queue = channel
        .queue_declare(
            queue_name,
            QueueDeclareOptions {
                durable: true,
                ..Default::default()
            },
            FieldTable::default(),
        )
        .await?;

    println!("📥 Waiting message on queue: {}", queue.name());

    let mut consumer = channel
        .basic_consume(
            queue_name,
            "rust_consumer",
            BasicConsumeOptions::default(),
            FieldTable::default(),
        )
        .await?;

    tokio::spawn(async move {
        while let Some(delivery) = consumer.next().await {
            match delivery {
                Ok(delivery) => {
                    let msg = String::from_utf8_lossy(&delivery.data).to_string();

                    handler(msg.clone());

                    if let Err(err) = channel
                        .basic_ack(delivery.delivery_tag, BasicAckOptions::default())
                        .await
                    {
                        eprintln!("❌ Fail to set ACK: {:?}", err);
                    }
                }
                Err(err) => eprintln!("❌ Consumer Erro: {:?}", err),
            }
        }
    });

    Ok(())
}
