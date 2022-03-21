package cmd

import (
	"fmt"
	"help_center/config"
	"help_center/internal/server"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath string

var rootCommand = &cobra.Command{
	Use:   "blockchain-wallet",
	Short: "run blockchain wallet",
	PreRun: func(cmd *cobra.Command, args []string) {
		// 协程跑爬虫
		//go
		log.Print(111)
	},
	Run: func(cmd *cobra.Command, args []string) {
		//跑 light-house
		server.RunApp()
	},
}

func init() {

	rootCommand.PersistentFlags().StringVar(&configPath, "config-path", "./config", "system config path")
	cobra.OnInitialize(initConfig)

	_ = viper.BindPFlag("config-path", rootCommand.PersistentFlags().Lookup("config-path"))
}

func initConfig() {

	viper.AddConfigPath(configPath)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {

		fmt.Println("配置读取出错1: ", err)
		return
	}
	// 监听配置
	viper.OnConfigChange(func(in fsnotify.Event) {
		config.RefreshConf()
	})
	viper.WatchConfig()
}

func Execute() {

	err := rootCommand.Execute()
	if err != nil {
		fmt.Println("启动失败: ", err)
		os.Exit(1)
	}
	log.Print(111)
}
