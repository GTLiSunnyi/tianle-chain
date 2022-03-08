use crate::grpc_client::GrpcClient;

pub struct Client {
    pub grpc: GrpcClient,
}

impl Client {
    pub fn new() -> Client {
        Client {
            grpc: GrpcClient::new(),
        }
    }
}
