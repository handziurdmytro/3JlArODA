use argon2::{
    Argon2,
    password_hash::{PasswordHash, PasswordHasher, PasswordVerifier, SaltString, rand_core::OsRng},
};

#[derive(Debug, Clone)]
pub struct AuthSecretEngine {
    pepper: String,
}

impl AuthSecretEngine {
    pub fn new(pepper: String) -> Self {
        Self { pepper }
    }

    pub fn hash_password(&self, plain_password: &str) -> Result<String, String> {
        let mut peppered_password = String::with_capacity(plain_password.len() + self.pepper.len());
        peppered_password.push_str(plain_password);
        peppered_password.push_str(&self.pepper);

        let salt = SaltString::generate(&mut OsRng);
        let argon2 = Argon2::default();

        let password_hash = argon2
            .hash_password(peppered_password.as_bytes(), &salt)
            .map_err(|e| format!("Hashing error: {}", e))?;

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
}

mod tests {}
#[cfg(test)]
#[test]
fn simple_test() {
    let pepper = "123456".to_string();
    let engine = AuthSecretEngine::new(pepper);

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
