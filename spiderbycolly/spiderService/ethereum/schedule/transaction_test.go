package schedule_test

import (
	"fmt"
	"help_center/config"
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"help_center/spiderbycolly/spiderService/ethereum"

	ss "help_center/spiderbycolly/spiderService/ethereum/schedule"
	ethUtil "help_center/spiderbycolly/spiderService/ethereum/util"
	"log"
	"os"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func initDatabse() {
	viper.AddConfigPath("../../config")
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {

		fmt.Println("配置读取出错: ", err)
		return
	}

	// 监听配置
	viper.OnConfigChange(func(in fsnotify.Event) {

		config.RefreshConf()
	})
	viper.WatchConfig()
	// 尝试初始化数据库连接
	err := database.Init(true)
	// 关闭连接
	//defer database.CloseConn()

	if err != nil {

		common.Logger.Info("数据库初始化失败:", err)
		os.Exit(1)
		return
	}
}

func TestRun(t *testing.T) {
	initDatabse()
	_ = database.DB()
	common.InitLogger("ethereum")

	re := ethUtil.InitETHInstance()
	log.Print(re)
	if err := ethereum.Run(config.ETHConf()); err != nil {

		common.Logger.Info("ethereum service 启动失败:%s", err)
		os.Exit(1)
		return
	}
}

func TestSyncBlockNumber(t *testing.T) {
	initDatabse()
	var err error
	//// 尝试初始化数据库连接
	err = database.CreatTransTable("transaction")
	tc := database.Tc
	log.Print(tc)
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
	conf := config.APPConf()
	bc := database.Bc
	log.Print(conf, bc)
	common.InitLogger("ethereum")

	ss.SyncTransaction()
}
