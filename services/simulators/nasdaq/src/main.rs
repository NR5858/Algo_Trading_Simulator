use futures_util::stream::StreamExt;
use lapin::{options::*, types::FieldTable, Connection, ConnectionProperties};
use tokio_amqp::*;

#[tokio::main]
async fn main() {
    // Connect to RabbitMQ server
    let addr = std::env::var("AMQP_ADDR")
        .unwrap_or_else(|_| "amqp://guest:guest@localhost:5672/%2f".into());
    let conn = Connection::connect(&addr, ConnectionProperties::default().with_tokio())
        .await
        .expect("connection error");

    // Create a channel
    let channel = conn.create_channel().await.expect("Error creating channel");

    // Declare a queue
    let _queue = channel
        .queue_declare(
            "fix_messages",
            QueueDeclareOptions::default(),
            FieldTable::default(),
        )
        .await
        .expect("queue declare error");

    // Consume messages from the queue
    let mut consumer = channel
        .basic_consume(
            "fix_messages",
            "my_consumer",
            BasicConsumeOptions::default(),
            FieldTable::default(),
        )
        .await
        .expect("basic consume error");

    println!("Waiting for messages...");

    while let Some(delivery) = consumer.next().await {
        match delivery {
            Ok((_, delivery)) => {
                let message = std::str::from_utf8(&delivery.data).expect("invalid utf-8");
                println!("Received message: {}", message);

                // Acknowledge the message
                delivery
                    .ack(BasicAckOptions::default())
                    .await
                    .expect("ack error");

                // Forward the message to your simulator logic
                simulate(message).await;
            }
            Err(error) => {
                eprintln!("Error receiving message: {:?}", error);
            }
        }
    }
}

async fn simulate(message: &str) {
    // Your simulator logic here
    println!("Simulating with message: {}", message);
}
