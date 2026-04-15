use crate::auth::engine::AuthSecretEngine;
use crate::auth::server::pb::auth_secret_service_server::AuthSecretService;
use crate::auth::server::pb::{
    HashRequest, HashResponse, SignJwtRequest, SignJwtResponse, ValidateJwtRequest,
    ValidateJwtResponse, VerifyRequest, VerifyResponse,
};
use tonic::{Request, Response, Status};

pub mod pb {
    tonic::include_proto!("auth_secret");
}

#[derive(Debug)]
pub struct AuthSecretServiceImpl {
    pub engine: AuthSecretEngine,
}

impl AuthSecretServiceImpl {
    pub fn new(engine: AuthSecretEngine) -> Self {
        Self { engine }
    }
}

#[tonic::async_trait]
impl AuthSecretService for AuthSecretServiceImpl {
    async fn hash_password(
        &self,
        request: Request<HashRequest>,
    ) -> Result<Response<HashResponse>, Status> {
        let req = request.into_inner();
        println!("[INFO] Received request to hash a password");

        let hash_string = self
            .engine
            .hash_password(&req.plain_password)
            .map_err(|e| Status::internal(e))?;

        Ok(Response::new(HashResponse { hash_string }))
    }

    async fn verify_password(
        &self,
        request: Request<VerifyRequest>,
    ) -> Result<Response<VerifyResponse>, Status> {
        let req = request.into_inner();
        println!("[INFO] Received request to verify a password");

        let is_valid = self
            .engine
            .verify_password(&req.plain_password, &req.hash_string);

        Ok(Response::new(VerifyResponse { is_valid }))
    }

    async fn sign_jwt(
        &self,
        request: Request<SignJwtRequest>,
    ) -> Result<Response<SignJwtResponse>, Status> {
        let req = request.into_inner();
        println!(
            "[INFO] received request to sign JWT for user: {}",
            req.user_id
        );

        let token = self
            .engine
            .sign_jwt(&req.user_id, &req.username)
            .map_err(|e| Status::internal(e))?;

        Ok(Response::new(SignJwtResponse { token }))
    }

    async fn validate_jwt(
        &self,
        request: Request<ValidateJwtRequest>,
    ) -> Result<Response<ValidateJwtResponse>, Status> {
        let req = request.into_inner();
        println!("[INFO] received request to validate JWT");

        let response = match self.engine.validate_jwt(&req.token) {
            Some((user_id, username)) => ValidateJwtResponse {
                user_id,
                username,
                is_valid: true,
            },
            None => ValidateJwtResponse {
                user_id: String::new(),
                username: String::new(),
                is_valid: false,
            },
        };

        Ok(Response::new(response))
    }
}
