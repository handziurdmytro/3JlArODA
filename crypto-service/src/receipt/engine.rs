use crate::receipt::errors::ReceiptError;
use ed25519_dalek::{Signature, Signer, SigningKey, VerifyingKey};

#[derive(Clone, Debug)]
pub struct ReceiptEngine {
    signing_key: SigningKey,
}

impl ReceiptEngine {
    pub fn new(hex: &str) -> Result<Self, ReceiptError> {
        let bytes = hex::decode(hex)?;

        let key_bytes: [u8; 32] = bytes
            .try_into()
            .map_err(|_| ReceiptError::InvalidKeyLength)?;

        let signing_key = SigningKey::from_bytes(&key_bytes);
        Ok(Self { signing_key })
    }

    pub fn get_public_key(&self) -> String {
        let verifying_key: VerifyingKey = self.signing_key.verifying_key();
        hex::encode(verifying_key.to_bytes())
    }

    pub fn sign_data(&self, raw_data: &str) -> String {
        let signature: Signature = self.signing_key.sign(raw_data.as_bytes());
        hex::encode(signature.to_bytes())
    }
}
