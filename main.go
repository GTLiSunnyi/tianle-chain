package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/GTLiSunnyi/tianle-chain/cmd"
)

// rootCmd is the entry
var (
	rootCmd = &cobra.Command{
		Use: "tianle",
	}
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(cmd.StartCmd())
	rootCmd.AddCommand(cmd.InitCmd())
	// rootCmd.AddCommand(keysCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
