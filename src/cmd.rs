use clap::Parser;

/// tianle-chain
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
pub struct Args {
    /// ip of node to connect
    #[clap(short, long, default_value = "")]
    pub ip: String,
    /// port of grpc
    #[clap(short, long, default_value = "50001")]
    pub grpc_port: i32,
}

impl Args {
    pub fn new() -> Args {
        Args::parse()
    }
}
