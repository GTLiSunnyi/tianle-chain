package grpc

import (
	"context"
	"log"
	"time"

	"github.com/GTLiSunnyi/tianle-chain/proto"
	"google.golang.org/grpc"
)

type GrpcClient struct{}

// TODO grpc.WithInsecure()
func (g GrpcClient) Ping(contract_address string) {
	grpcClient, err := grpc.Dial(contract_address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer grpcClient.Close()

	dSpaceClient := proto.NewNetworkServiceClient(grpcClient)

	// 设置 grpc 超时时间
	clientDeadline := time.Now().Add(time.Duration(3 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	pingReq := &proto.PingRequest{}
	_, err = dSpaceClient.Ping(ctx, pingReq)
	if err != nil {
		log.Fatal(err)
	}
}
