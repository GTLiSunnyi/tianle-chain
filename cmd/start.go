package cmd

import (
	"github.com/spf13/cobra"

	"github.com/GTLiSunnyi/tianle-chain/grpc"
	"github.com/GTLiSunnyi/tianle-chain/types"
)

func StartCmd() *cobra.Command {
	var port int32
	var home string
	var ip string

	cmd := &cobra.Command{
		Use:     "start",
		Short:   "启动 tianle-chain",
		Example: `tianle start [--port] [--home] [--ip]`,
		Args:    cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// ping 节点
			grpcClient := new(grpc.GrpcClient)
			grpcClient.Ping(ip)

			// TODO 检查有没有 home 目录完整性

			// 开启 grpc server
			go func() {
				grpc.StartNetwork(port)
			}()

			select {}
		},
	}

	cmd.Flags().Int32VarP(&port, "port", "p", types.DefaultPort, "port of chain")
	cmd.Flags().StringVarP(&home, "home", "", types.DefaultHome, "path of data")
	cmd.Flags().StringVarP(&ip, "ip", "", "", "ip of node to connect")

	return cmd
}
