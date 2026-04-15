use dotenvy::dotenv;
use ed25519_dalek::SigningKey;
use rand_core::{OsRng, RngCore};
use std::env;
use std::fs::OpenOptions;
use std::io::Write;
use tracing::{info, warn};

pub struct AppConfig {
    pub port: u16,
    pub pepper: String,
    pub signing_key_hex: String,
    pub jwt_secret: String,
    pub jwt_exp_seconds: u64,
}

impl AppConfig {
    pub fn load() -> Self {
        dotenv().ok();

        let port = env::var("PORT")
            .unwrap_or_else(|_| "3030".into())
            .parse::<u16>()
            .expect("PORT must be a valid number (1-65535)");

        let jwt_exp_seconds = env::var("JWT_ESP_SECONDS")
            .unwrap_or_else(|_| "86400".into())
            .parse::<u64>()
            .expect("JWT_EXP_SECONDS must be a valid number");

        Self {
            port,
            jwt_exp_seconds,
            pepper: env::var("PEPPER").expect("FATAL: PEPPER is not set"),
            signing_key_hex: env::var("SIGNING_KEY").expect("FATAL: SIGNING_KEY is not set"),
            jwt_secret: env::var("JWT_SECRET").expect("FATAL: JWT_SECRET is not set"),
        }
    }
}
