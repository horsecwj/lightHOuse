package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"help_center/internal/server"
	"os"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use:   "light-house",
	Short: "arbitrage",
	PreRun: func(cmd *cobra.Command, args []string) {
		// 协程跑爬虫
	},
	Run: func(cmd *cobra.Command, args []string) {
		//跑 light-house
		server.RunApp()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
