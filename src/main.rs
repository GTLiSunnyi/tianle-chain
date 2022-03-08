mod client;
mod cmd;
mod config;
mod grpc_client;
mod grpc_server;
mod proto;
mod utils;

#[tokio::main]
async fn main() {
    std::env::set_var("RUST_LOG", "info");
    env_logger::init();

    let args = cmd::Args::new();
    
    let client = client::Client::new();

    // 连接到 kad 网络
    if args.ip != "" {
        client.grpc.ping(&args.ip).await;
    } else {
        // 创建一条新的链
        println!("即将创建一条新链!");
    }

    let env_config = config::Config::new();

    // 创建密钥

    grpc_server::start_network_service(args.grpc_port).await;
}
