use crate::proto::network::{PingRequest, network_service_client::NetworkServiceClient};

pub struct GrpcClient {}

// 对 server 发起的 grpc
impl GrpcClient {
    pub fn new() -> GrpcClient {
        GrpcClient{}
    }

    // TODO grpc 超时处理
    pub async fn ping(&self, addr: &String) {
        let mut client = NetworkServiceClient::connect(format!("http://{}", addr)).await.unwrap();

        let request = tonic::Request::new(PingRequest{});

        let response = client.ping(request).await.unwrap();

        println!("RESPONSE={:?}", response);
    }
}
