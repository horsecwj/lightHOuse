package schedule

import (
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"strings"
	"time"

	"github.com/robfig/cron"
)

var transferSyncing = false

func createTransferScheduler(c *cron.Cron) (err error) {
	err = c.AddFunc("@every 3s", syncTransfer)
	if err != nil {

		return
	}

	return
}

func syncTransfer() {

	if transferSyncing {

		return
	}

	transferSyncing = true
	defer func() {

		transferSyncing = false
	}()
	//db := database.DB()
	//conf := EVMConfig.Chain
	//log.Print(conf, db)
	// 查询数据
	array, err := database.DB().GetUnConfirmTransfer(EVMConfig.Chain)
	if err != nil {

		common.Logger.Info("查询交易失败:", err)
		return
	}

	if len(array) == 0 {

		return
	}

	//
	for _, item := range array {

		common.Logger.Info("获取交易信息:", item.TxHash)
		t, pending, err := EVMInstance().GetTransaction(item.TxHash)
		if err != nil {
			common.Logger.Infof("查询链上状态失败:%v,%v", item.TxHash, err)
			// 10分钟后去查询记录不存在报错
			if (time.Since(time.Unix(item.Timestamp, 0)) > time.Minute*10) && strings.Contains(err.Error(), "not found") {
				if e := item.Failed("交易没发送出去"); e != nil {
					common.Logger.Info("更新失败状态出错:", e, "-", item.TxHash)
				}
				common.Logger.Info("交易没发送出去:", item.TxHash)
			}
			continue
		}

		// 被回滚
		if t == nil && pending == false {

			if e := item.Failed("交易被回滚"); e != nil {

				common.Logger.Info("更新失败状态出错:", e, "-", item.TxHash)
			}

			common.Logger.Info("交易被回滚:", item.TxHash)
			continue
		}

		// 被确认
		if t != nil && strings.EqualFold(t.Hash().String(), item.TxHash) && pending == false {

			// ERC20 再查询执行状态
			if !strings.EqualFold(item.Symbol, database.ETHSymbol) {

				receipt, err := EVMInstance().GetTransactionReceipt(item.TxHash)
				if err != nil {

					common.Logger.Info("查询状态失败:", item.TxHash, "-- err:", err)
					continue
				}

				// 合约执行失败
				if receipt.Status != 1 {

					if e := item.Failed("合约执行失败"); e != nil {

						common.Logger.Info("更新失败状态出错:", e, "-", item.TxHash)
					}
					continue
				}

			}
			if e := item.Succeed(); e != nil {

				common.Logger.Info("更新不可逆状态出错:", e, "-", item.TxHash)
			}
		}
	}
}
