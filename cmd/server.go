package cmd

import (
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"help_center/spiderbycolly/spiderService/server"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

	rootCommand.AddCommand(createServerCommand())
}

// 启动API命令
func createServerCommand() *cobra.Command {

	serverCommand := &cobra.Command{
		Use:   "server",
		Short: "run api server",
		Run: func(cmd *cobra.Command, args []string) {

			// 初始化Logger
			common.InitLogger("server")

			// 尝试初始化数据库连接
			err := database.Init(false)
			// 关闭连接
			defer database.CloseConn()

			if err != nil {

				common.Logger.Info("database 连接失败:", err)
				os.Exit(1)
				return
			}

			database.AutoMigrate()
			viper.Set("name", "server")
			if err := server.Run(":8888", false); err != nil {
				common.Logger.Info("fail to run server with error:", err)
				os.Exit(1)
			}
		},
	}
	return serverCommand
}
