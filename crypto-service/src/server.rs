use crate::crypto::engine::ReceiptCryptoEngine;
use crate::server::pd::receipt_signer_server::ReceiptSigner;
use crate::server::pd::{SignRequest, SignResponse};
use tonic::{Request, Response, Status};

pub mod pd {
    tonic::include_proto!("receipt");
}

#[derive(Debug)]
pub struct ReceiptService {
    pub crypto_engine: ReceiptCryptoEngine,
}

impl ReceiptService {
    pub fn new(crypto_engine: ReceiptCryptoEngine) -> Self {
        Self { crypto_engine }
    }
}

#[tonic::async_trait]
impl ReceiptSigner for ReceiptService {
    async fn sign_receipt(
        &self,
        request: Request<SignRequest>,
    ) -> Result<Response<SignResponse>, Status> {
        let req = request.into_inner();

        let card = if req.card_number.is_empty() {
            "NO_CARD"
        } else {
            &req.card_number
        };

        let raw_string = format!(
            "{}|{}|{}|{}|{}|{}",
            req.check_number, req.id_employee, card, req.print_date, req.sum_total, req.vat,
        );

        println!("[INFO] received request form signing a bill: {}", req.check_number);

        let signature_hex = self.crypto_engine.sign_data(&raw_string);
        let public_key_hex = self.crypto_engine.get_public_key();

        let response = SignResponse {
            signature_hex,
            public_key_hex,
        };

        Ok(Response::new(response))
    }
}
