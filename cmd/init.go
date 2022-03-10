package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/GTLiSunnyi/tianle-chain/crypto"
	"github.com/GTLiSunnyi/tianle-chain/types"
	"github.com/GTLiSunnyi/tianle-chain/utils"
)

func InitCmd() *cobra.Command {
	var home string

	cmd := &cobra.Command{
		Use:     "init",
		Short:   "初始化一条新的链",
		Example: `tianle init [--home]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !utils.IsDirEmpty(home) {
				log.Fatal("home 目录不为空!")
			}

			// 新建 chain_config.json 文件
			utils.CopyFile("./script/chain_config.json", home+"/chain_config.json")

			// 创建密钥
			var pwd = utils.GetPassword()
			wallet := crypto.NewWallet()
			wallet.ExportWallet(home, pwd)

			// TODO 创建创世区块

			fmt.Println("初始化完成")
			fmt.Println("请配置好 home/data/chain_config.json")
			os.Exit(-1)

			return nil
		},
	}

	cmd.Flags().StringVarP(&home, "home", "", types.DefaultHome, "path of data")

	return cmd
}
