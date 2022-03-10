package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/GTLiSunnyi/tianle-chain/proto"
)

type network struct {
	proto.UnimplementedNetworkServiceServer
}

func (n *network) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	// TODO
	return &proto.PingResponse{Config: &proto.ChainConfig{ChainId: "main"}}, nil
}

func StartNetwork(port int32) {
	rpc := grpc.NewServer()
	proto.RegisterNetworkServiceServer(rpc, new(network))

	listener, err := net.Listen("tcp", ":"+fmt.Sprintf("%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("network start at 127.0.0.1:" + fmt.Sprintf("%d", port))

	log.Fatal(rpc.Serve(listener))
}
