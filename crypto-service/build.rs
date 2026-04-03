fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=../proto/receipt.proto");
    println!("cargo:rerun-if-changed=../proto/auth_secret.proto");

    tonic_prost_build::compile_protos("../proto/receipt.proto")?;
    tonic_prost_build::compile_protos("../proto/auth_secret.proto")?;

    Ok(())
}