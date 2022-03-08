use tonic::{Request, Response, Status, transport::Server};
use crate::proto::network::{PingRequest, PingResponse, network_service_server::{NetworkService, NetworkServiceServer}};

// grpc server
pub struct GrpcNetworkServer {}

#[tonic::async_trait]
impl NetworkService for GrpcNetworkServer {
    async fn ping(&self, _: Request<PingRequest>) -> Result<Response<PingResponse>, Status> {
        let res = PingResponse{
            res: String::from("pong")
        };

        Ok(Response::new(res))
    }
}

pub async fn start_network_service(port: i32) {
    let addr = format!("127.0.0.1:{}", port).parse().unwrap();
    let grpc_network = GrpcNetworkServer {};

    println!("network server running on {}", addr);

    Server::builder()
        .add_service(NetworkServiceServer::new(grpc_network))
        .serve(addr)
        .await.unwrap();
}
