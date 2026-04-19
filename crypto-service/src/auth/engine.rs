use crate::auth::errors::AuthError;
use argon2::{
    Argon2,
    password_hash::{PasswordHash, PasswordHasher, PasswordVerifier, SaltString, rand_core::OsRng},
};
use jsonwebtoken::{DecodingKey, EncodingKey, Header, Validation, decode, encode};
use serde::{Deserialize, Serialize};
use std::time::{SystemTime, UNIX_EPOCH};

#[derive(Debug, Serialize, Deserialize)]
struct Claims {
    sub: String,
    username: String,
    exp: u64,
    iat: u64,
}
#[derive(Debug, Clone)]
pub struct AuthSecretEngine {
    pepper: String,
    jwt_secret: String,
    jwt_exp_seconds: u64,
}

impl AuthSecretEngine {
    pub fn new(pepper: String, jwt_secret: String, jwt_exp_seconds: u64) -> Self {
        Self {
            pepper,
            jwt_secret,
            jwt_exp_seconds,
        }
    }

    pub fn hash_password(&self, plain_password: &str) -> Result<String, AuthError> {
        let mut peppered_password = String::with_capacity(plain_password.len() + self.pepper.len());
        peppered_password.push_str(plain_password);
        peppered_password.push_str(&self.pepper);

        let salt = SaltString::generate(&mut OsRng);
        let argon2 = Argon2::default();

        let password_hash = argon2
            .hash_password(peppered_password.as_bytes(), &salt)
            .map_err(|e| AuthError::Hashing(e.to_string()))?;

        Ok(password_hash.to_string())
    }

    pub fn verify_password(&self, plain_password: &str, hash_str: &str) -> bool {
        let mut peppered_password = String::with_capacity(plain_password.len() + self.pepper.len());
        peppered_password.push_str(plain_password);
        peppered_password.push_str(&self.pepper);

        let parsed_hash = match PasswordHash::new(hash_str) {
            Ok(hash) => hash,
            Err(_) => return false,
        };

        Argon2::default()
            .verify_password(peppered_password.as_bytes(), &parsed_hash)
            .is_ok()
    }

    pub fn sign_jwt(&self, user_id: &str, username: &str) -> Result<String, AuthError> {
        let now = SystemTime::now().duration_since(UNIX_EPOCH)?.as_secs();

        let claims = Claims {
            sub: user_id.to_string(),
            username: username.to_string(),
            exp: now + self.jwt_exp_seconds,
            iat: now,
        };

        encode(
            &Header::default(),
            &claims,
            &EncodingKey::from_secret(self.jwt_secret.as_bytes()),
        )
        .map_err(|e| AuthError::JwtSign(e.to_string()))
    }

    pub fn validate_jwt(&self, token: &str) -> Option<(String, String)> {
        let result = decode::<Claims>(
            token,
            &DecodingKey::from_secret(self.jwt_secret.as_bytes()),
            &Validation::default(),
        );

        match result {
            Ok(data) => Some((data.claims.sub, data.claims.username)),
            Err(_) => None,
        }
    }
}

mod tests {}
#[cfg(test)]
#[test]
fn simple_test() {
    let pepper = "123456".to_string();
    let engine = AuthSecretEngine::new(pepper, "".to_string(), 0);

    let pass1 = "c1$c0".to_string();
    // println!("{}", pass1);
    let hash1 = engine.hash_password(&pass1).unwrap();
    // println!("{}", hash1);
    assert!(engine.verify_password(&pass1, &hash1));

    let pass2 = "p4$$w0rd".to_string();
    // println!("{}", pass2);
    let hash2 = engine.hash_password(&pass2).unwrap();
    // println!("{}", hash2);
    assert!(engine.verify_password(&pass2, &hash2));
}
