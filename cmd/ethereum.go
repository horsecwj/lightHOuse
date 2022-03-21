package cmd

import (
	"help_center/config"
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"help_center/spiderbycolly/spiderService/ethereum"

	ethUtil "help_center/spiderbycolly/spiderService/ethereum/util"
	"os"

	"github.com/spf13/cobra"
)

var ethereumCommand = &cobra.Command{
	Use:   "ethereum",
	Short: "run ethereum service",
	Run: func(cmd *cobra.Command, args []string) {

		// 尝试初始化数据库连接
		err := database.CreatTransTable("transaction")
		// 关闭连接
		defer database.Tc.Close()

		// 初始化Logger
		common.InitLogger("ethereum")

		if err != nil {

			common.Logger.Info("ethereum -> ", err)
			os.Exit(1)
			return
		}

		err = database.CreatBlockTable("block")
		// 关闭连接
		defer database.Bc.Close()
		if err != nil {
			common.Logger.Info("ethereum -> ", err)
			os.Exit(1)
			return
		}

		ethUtil.InitETHInstance()
		if err := ethereum.Run(config.ETHConf()); err != nil {

			common.Logger.Info("ethereum service 启动失败:%s", err)
			os.Exit(1)
			return
		}
	},
}

func init() {

	rootCommand.AddCommand(ethereumCommand)
}
