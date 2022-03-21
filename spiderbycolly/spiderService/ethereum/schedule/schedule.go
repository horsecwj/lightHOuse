package schedule

import (
	"help_center/config"
	ethUtil "help_center/spiderbycolly/spiderService/ethereum/util"

	"github.com/robfig/cron"
)

var (
	EVMConfig *config.ETHConfig
)

func EVMInstance() *ethUtil.ETHClient {
	c := EVMConfig
	ethUtil.EVMInstance(c.Chain)
	return ethUtil.EVMInstance(EVMConfig.Chain)
}

// 创建定时执行任务
func create() (c *cron.Cron, err error) {

	c = cron.New()

	//添加任务执行  (同比eth 区块 ，发现交易，存入 bcw—trancscation)
	err = createTransactionScheduler(c)
	if err != nil {

		return
	}
	return
}

func Run(conf *config.ETHConfig) (err error) {

	EVMConfig = conf
	c, err := create()

	if err != nil {

		return
	}

	c.Run()
	return
}
