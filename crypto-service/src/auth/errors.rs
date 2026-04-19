use thiserror::Error;

#[derive(Error, Debug)]
pub enum AuthError {
    #[error("Failed to hash password: {0}")]
    Hashing(String),

    #[error("Failed to sign JWT: {0}")]
    JwtSign(String),

    #[error("System time error: {0}")]
    Time(#[from] std::time::SystemTimeError),
}
