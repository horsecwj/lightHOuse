package spiderbycolly

import (
	"github.com/spf13/viper"
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"help_center/spiderbycolly/spiderService"
	"help_center/spiderbycolly/spiderService/server"
	"os"
)

func RunSpiderSpot() {
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
}

func RunSpiderApi() {
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
}
