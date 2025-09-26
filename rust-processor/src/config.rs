use anyhow::{Result, bail};
use serde::Deserialize;

#[derive(Debug, Deserialize, Clone)]
pub struct Config {
    pub rabbitmq_default_user: String,
    pub rabbitmq_default_pass: String,
    pub rabbitmq_host: String,
    pub rabbitmq_port: String,

    pub postgres_user: String,
    pub postgres_password: String,
    pub postgres_db: String,
    pub postgres_host: String,
    pub postgres_port: String,
}

impl Config {
    pub fn load() -> Result<Self> {
        match envy::from_env::<Config>() {
            Ok(config) => Ok(config),
            Err(e) => {
                bail!("Fail to loading ambient envirements: {}", e)
            }
        }
    }

    pub fn rabbitmq_amqp_addr(&self) -> String {
        format!(
            "amqp://{}:{}@{}:{}/%2f",
            self.rabbitmq_default_user,
            self.rabbitmq_default_pass,
            self.rabbitmq_host,
            self.rabbitmq_port
        )
    }

    pub fn database_url(&self) -> String {
        format!(
            "postgres://{}:{}@{}:{}/{}?sslmode=disable",
            self.postgres_user,
            self.postgres_password,
            self.postgres_host,
            self.postgres_port,
            self.postgres_db
        )
    }
}
