package schedule

import (
	"help_center/spiderbycolly/database"

	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"help_center/config"
	"help_center/spiderbycolly/common"

	"math/big"
	"strings"
	"sync"

	"github.com/robfig/cron"
)

var (
	BlockSyncing       = false
	TransactionSyncing = false
)

func createTransactionScheduler(c *cron.Cron) (err error) {

	err = c.AddFunc("@every 15s", SyncBlockNumber)
	if err != nil {

		return
	}

	err = c.AddFunc("@every 1s", SyncTransaction)
	if err != nil {

		return
	}

	return
}

// 同步区块高度解析任务
func SyncBlockNumber() {
	if BlockSyncing {

		return
	}
	BlockSyncing = true

	defer func() {

		BlockSyncing = false
	}()

	// 判断是否开启充值立即通知
	diffHeight := config.ETHConf().ConfirmHeight
	if config.APPConf().Notice {

		diffHeight = 0
	}

	// 创建执行者
	scheduler := &database.SyncBlockScheduler{

		Height:           uint(diffHeight),
		ChainType:        EVMConfig.Chain,
		ChainBlockNumber: EVMInstance().LastBlockNumber,
	}

	scheduler.Start()
}

// 同步交易
func SyncTransaction() {
	if TransactionSyncing {

		return
	}
	if EVMConfig == nil {
		EVMConfig = config.ETHConf()

	}
	TransactionSyncing = true

	defer func() {

		TransactionSyncing = false
	}()

	var limit uint = 3
	var wait sync.WaitGroup
	scheduler := &database.SyncTransactionScheduler{
		ChainType:  EVMConfig.Chain,
		WaitGroup:  &wait,
		BlockLimit: limit,
		BlockHandler: func(blockNumber uint) (info *database.BlockInfo, err error) {

			block, err := EVMInstance().GetBlockByNumber(int64(blockNumber))

			if err != nil {

				return
			}
			info = &database.BlockInfo{
				Number:    int(blockNumber),
				Hash:      block.Hash().Hex(),
				TxSize:    block.Transactions().Len(),
				Timestamp: int64(block.Time()),

				Difficulty: int64(block.Difficulty().Uint64()),
				Getter: func(index int) (tx interface{}) {

					return block.Transactions()[index]
				},
			}

			return
		},
		TxHandler: parseTransaction,
	}

	scheduler.Start()
	wait.Wait()
}

// 解析交易

// 解析交易
func parseTransaction(worker *database.TransactionWorker) {
	defer func() {

		worker.Wait.Done()
	}()

	tx := worker.BlockInfo.Getter(worker.Index).(*ethTypes.Transaction)
	receiver := tx.To()
	if receiver == nil {

		worker.Success = true
		common.Logger.Info(tx.Hash().Hex(), "：无接收地址")
		return
	}

	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if msg, err := tx.AsMessage(ethTypes.NewEIP155Signer(chainID)); err == nil {
	//	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
	//}

	// 不是平台地址充值，直接完成
}

func parseInputDataQuantityHex(value string) string {

	if strings.HasPrefix(value, "0") {

		return parseInputDataQuantityHex(strings.Replace(value, "0", "", 1))
	} else {

		return "0x" + value
	}
}

// 验证
func validateERC20Balance(contract string, method string, address string, amount *big.Int) bool {
	balance, err := EVMInstance().GetERC20Balance(contract, method, address)
	if err != nil {

		common.Logger.Warn("解析ERC20余额出错:", err)
		return false
	}

	return balance.Cmp(amount) >= 0
}
