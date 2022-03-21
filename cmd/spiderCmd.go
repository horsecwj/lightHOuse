package cmd

import (
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"help_center/spiderbycolly/spiderService"
	"os"

	"github.com/spf13/cobra"
)

var SpiderCommand = &cobra.Command{

	Use:   "spiderCmd",
	Short: "start spiderCmd sync service",
	Run: func(cmd *cobra.Command, args []string) {
		// 尝试初始化数据库连接
		err := database.Init(false)
		// 关闭连接
		defer database.CloseConn()
		// 初始化Logger
		common.InitLogger("spiderCmd")
		if err != nil {

			common.Logger.Info("spiderCmd -> ", err)
			os.Exit(1)
			return
		}
		database.AutoMigrate()

		if err := spiderService.Run(); err != nil {
			common.Logger.Info(" service 启动失败:%s", err)
			os.Exit(1)
			return
		}

	},
}

func init() {

	rootCommand.AddCommand(SpiderCommand)
}
