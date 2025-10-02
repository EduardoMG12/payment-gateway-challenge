use anyhow::Result;
use futures_lite::stream::StreamExt;
use lapin::{
    Channel, Connection, ConnectionProperties,
    options::{BasicAckOptions, BasicConsumeOptions, QueueDeclareOptions},
    types::FieldTable,
};
use std::{future::Future, sync::Arc};

pub async fn create_connection(addr: &str) -> Result<Connection> {
    let conn = Connection::connect(addr, ConnectionProperties::default()).await?;
    Ok(conn)
}

pub async fn create_channel(conn: &Connection) -> Result<Channel> {
    let channel = conn.create_channel().await?;
    Ok(channel)
}
pub async fn consume_queue<F, Fut>(
    channel: Arc<Channel>,
    queue_name: &str,
    handler: F,
) -> Result<()>
where
    F: Fn(String) -> Fut + Send + Sync + 'static,
    Fut: Future<Output = ()> + Send + 'static,
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

                    // chama handler assíncrono
                    handler(msg.clone()).await;

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
