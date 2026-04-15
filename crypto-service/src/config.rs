use ed25519_dalek::SigningKey;
use rand_core::{OsRng, RngCore};
use std::fs::OpenOptions;
use std::io::Write;
use tracing::{info, warn};

fn get_or_generate(key: &str, generate: impl Fn() -> String) -> String {
    if let Ok(val) = std::env::var(key) {
        if !val.is_empty() {
            return val;
        }
    }

    let val = generate();

    let mut file = OpenOptions::new()
        .create(true)
        .append(true)
        .open(".env")
        .expect("Cannot open/create .env file");

    writeln!(file, "{}={}", key, val).expect("Cannot write to .env file");

    info!("Generated and saved new key: {}", key);
    val
}

fn get_or_default(key: &str, default: &str) -> String {
    if let Ok(val) = std::env::var(key) {
        if !val.is_empty() {
            return val;
        }
    }
    warn!("{} not set, using default: {}", key, default);
    default.to_string()
}

pub struct AppConfig {
    pub port: u16,
    pub pepper: String,
    pub signing_key_hex: String,
    pub jwt_secret: String,
    pub jwt_exp_seconds: u64,
}

impl AppConfig {
    pub fn load() -> Self {
        let _ = dotenvy::dotenv();

        let port = std::env::var("PORT")
            .unwrap_or_else(|_| "3030".to_string())
            .parse::<u16>()
            .expect("PORT must be a valid number (1-65535)");

        let pepper = get_or_generate("PEPPER", || {
            let mut bytes = [0u8; 32];
            OsRng.fill_bytes(&mut bytes);
            hex::encode(bytes)
        });

        let signing_key_hex = get_or_generate("SIGNING_KEY", || {
            let signing_key = SigningKey::generate(&mut OsRng);
            hex::encode(signing_key.to_bytes())
        });

        let jwt_secret = get_or_generate("JWT_SECRET", || {
            let mut bytes = vec![0u8; 32];
            OsRng.fill_bytes(&mut bytes);
            hex::encode(bytes)
        });

        let jwt_exp_seconds = get_or_default("JWT_EXP_SECONDS", "86400")
            .parse::<u64>()
            .expect("JWT_EXP_SECONDS must be a valid number");

        Self {
            port,
            pepper,
            signing_key_hex,
            jwt_secret,
            jwt_exp_seconds,
        }
    }
}
