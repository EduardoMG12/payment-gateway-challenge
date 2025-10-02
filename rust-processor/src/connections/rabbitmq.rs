use anyhow::Result;
use futures_lite::stream::StreamExt;
use lapin::{
    Channel, Connection, ConnectionProperties,
    options::{BasicAckOptions, BasicConsumeOptions, QueueDeclareOptions},
    types::FieldTable,
};
use std::future::Future;

pub async fn create_channel(amqp_addr: &str) -> Result<Channel> {
    let conn = Connection::connect(amqp_addr, ConnectionProperties::default()).await?;
    println!("✅ Conected to RabbitMQ!");
    let channel = conn.create_channel().await?;
    Ok(channel)
}

pub async fn consume_queue<F, Fut>(channel: Channel, queue_name: &str, handler: F) -> Result<()>
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

// Commented out because not used in this project now, fot don't show warning logs but can be useful for future implementations
// pub async fn publish<T: serde::Serialize>(
//     channel: &Channel, // Recebe uma referência ao canal existente
//     queue_name: &str,
//     message: T
// ) -> Result<()> {
//     let payload = serde_json::to_vec(&message)?;

//     channel
//         .basic_publish(
//             "",
//             queue_name,
//             lapin::options::BasicPublishOptions::default(),
//             &payload,
//             lapin::BasicProperties::default().with_delivery_mode(2.into()),
//         )
//         .await?
//         .await?;

//     println!("✅ Published to {}: {:?}", queue_name, message.serialize(serde_json::value::Serializer).unwrap());

//     Ok(())
// }
