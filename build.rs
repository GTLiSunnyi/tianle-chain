fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=proto");

    tonic_build::configure()
        .build_client(true)
        .format(true)
        .compile(
            &[
                "network.proto",
            ],
            &["proto"],
        )?;

    Ok(())
}
