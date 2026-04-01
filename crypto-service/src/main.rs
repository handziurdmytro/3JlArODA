use crate::crypto::engine::ReceiptCryptoEngine;
use crate::server::pd::receipt_signer_server::ReceiptSignerServer;
use crate::server::ReceiptService;
use std::net::SocketAddr;
use tonic::transport::Server;

mod server;
mod crypto;
mod reports;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let address: SocketAddr = "0.0.0.0:3030".parse()?;

    println!("[INFO] Cryptographic module initialization");
    let crypto_engine = ReceiptCryptoEngine::new();
    println!("[INFO] Server public key: {}", crypto_engine.get_public_key());

    let receipt_service = ReceiptService::new(crypto_engine);
    println!("[INFO] gRPC server is running on: {}", address);

    Server::builder()
        .add_service(ReceiptSignerServer::new(receipt_service))
        .serve(address)
        .await?;

    Ok(())
}
