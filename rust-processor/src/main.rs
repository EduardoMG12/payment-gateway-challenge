mod application;
mod config;
mod connections;
mod models;
mod processors;
mod repository;
mod services;

use crate::application::Application;
use anyhow::Result;

#[tokio::main]
async fn main() -> Result<()> {
    dotenv::dotenv().ok();
    let config = config::Config::load()?;

    let app = Application::build(&config).await?;

    app.run_with_loop().await?;

    Ok(())
}
