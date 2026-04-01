use ed25519_dalek::{Signature, Signer, SigningKey, VerifyingKey};
use rand_core::OsRng;

#[derive(Clone, Debug)]
pub struct ReceiptCryptoEngine {
    signing_key: SigningKey,
}

impl ReceiptCryptoEngine {
    pub fn new() -> Self {
        let mut csprng = OsRng;
        let signing_key: SigningKey = SigningKey::generate(&mut csprng);

        /* reading from file logic will be here */

        Self { signing_key }
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
