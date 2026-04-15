use crate::auth::engine::AuthSecretEngine;
use crate::auth::server::AuthSecretServiceImpl;
use crate::auth::server::pb::auth_secret_service_server::AuthSecretServiceServer;
use crate::config::AppConfig;
use crate::receipt::engine::ReceiptEngine;
use crate::receipt::server::ReceiptServiceImpl;
use crate::receipt::server::pd::receipt_service_server::ReceiptServiceServer;
use std::net::SocketAddr;
use tonic::transport::Server;
use tracing::info;
use tracing_subscriber::EnvFilter;
use tracing_subscriber::layer::SubscriberExt;
use tracing_subscriber::util::SubscriberInitExt;

mod auth;
mod config;
mod receipt;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    tracing_subscriber::registry()
        .with(EnvFilter::try_from_default_env().unwrap_or_else(|_| EnvFilter::new("info")))
        .with(tracing_subscriber::fmt::layer().json())
        .init();

    let config = AppConfig::load();

    let address: SocketAddr = format!("0.0.0.0:{}", &config.port).parse()?;

    info!("Cryptographic module initialization");

    let receipt_engine =
        ReceiptEngine::new(&config.signing_key_hex).expect("Failed to initialize ReceiptEngine");

    let auth_engine =
        AuthSecretEngine::new(config.pepper, config.jwt_secret, config.jwt_exp_seconds);

    let receipt_service = ReceiptServiceImpl::new(receipt_engine);
    let auth_service = AuthSecretServiceImpl::new(auth_engine);

    info!("Starting gRPC server...");
    Server::builder()
        .add_service(ReceiptServiceServer::new(receipt_service))
        .add_service(AuthSecretServiceServer::new(auth_service))
        .serve(address)
        .await?;

    Ok(())
}
