use crate::auth::engine::AuthSecretEngine;
use crate::auth::server::AuthSecretServiceImpl;
use crate::auth::server::pb::auth_secret_service_server::AuthSecretServiceServer;
use crate::crypto::engine::ReceiptEngine;
use crate::crypto::server::ReceiptServiceImpl;
use crate::crypto::server::pd::receipt_service_server::ReceiptServiceServer;
use std::net::SocketAddr;
use tonic::transport::Server;
use crate::config::AppConfig;

mod auth;
mod config;
mod crypto;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let config = AppConfig::load();

    let address: SocketAddr = format!("0.0.0.0:{}", &config.port).parse()?;

    println!("[INFO] Cryptographic module initialization");

    let receipt_engine = ReceiptEngine::new(&config.signing_key_hex)
        .expect("Failed to initialize ReceiptEngine");

    let auth_engine = AuthSecretEngine::new(config.pepper);

    let receipt_service = ReceiptServiceImpl::new(receipt_engine);
    let auth_service = AuthSecretServiceImpl::new(auth_engine);

    Server::builder()
        .add_service(ReceiptServiceServer::new(receipt_service))
        .add_service(AuthSecretServiceServer::new(auth_service))
        .serve(address)
        .await?;

    Ok(())
}
