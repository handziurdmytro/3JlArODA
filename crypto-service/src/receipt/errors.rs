use thiserror::Error;

#[derive(Error, Debug)]
pub enum ReceiptError {
    #[error("Invalid SIGNING_KEY hex format: {0}")]
    HexDecode(#[from] hex::FromHexError),

    #[error("PRIVATE_KEY must be exactly 32 bytes (64 hex chars)")]
    InvalidKeyLength,
}
