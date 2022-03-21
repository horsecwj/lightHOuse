package database

import (
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/common/types"
	"log"
	"math/big"

	ethTypes "github.com/ethereum/go-ethereum/core/types"

	//"help_center/spiderbycolly/database"
	"sync"
)

//import (
//	"help_center/spiderbycolly/common"
//	"help_center/spiderbycolly/common/types"
//	"help_center/spiderbycolly/database"
//	ethTypes "github.com/ethereum/go-ethereum/core/types"
//	"log"
//	"math/big"
//	"sync"
//)
//
type GetBlockNumber func() (uint, error)

type SyncBlockScheduler struct {
	Height           uint           // 区块确认高度
	ChainType        types.Chain    // 公链类型
	ChainBlockNumber GetBlockNumber // 获取区块链的当前高度
}

//
func (scheduler *SyncBlockScheduler) Start() {

	// 获取当前区块高度
	number, err := scheduler.ChainBlockNumber()

	if err != nil {

		common.Logger.Info("查询当前区块高度失败:", err)
		return
	}

	// 计算可同步高度
	current := number - scheduler.Height
	LoclaUnhandle, err := DB().MaxUnHandleBlockNumber(scheduler.ChainType)
	if err != nil {

		common.Logger.Info("查询本地区块高度失败:", err)
		return
	}
	CountUnHandleBlockNumber, err := DB().CountUnHandleBlockNumber(scheduler.ChainType)
	if err != nil {

		common.Logger.Info("查询本地区块高度失败:", err)
		return
	}
	if LoclaUnhandle == 0 {
		LoclaUnhandle = 100000
	}

	if CountUnHandleBlockNumber >= 100 {

		return
	}

	if current >= LoclaUnhandle+500 {

		current = LoclaUnhandle + 500
	}

	common.Logger.Info("发现新高度:", current, " 当前高度:", LoclaUnhandle)

	// 保存至数据库
	err = DB().SaveBlockNumbers(LoclaUnhandle, current, scheduler.ChainType)
	if err != nil {

		common.Logger.Info("保存区块解析任务失败: ", err)
	}
	return
}

type SyncTransactionScheduler struct {
	WaitGroup    *sync.WaitGroup
	BlockLimit   uint
	ChainType    types.Chain
	BlockHandler BlockInfoHandler
	TxHandler    TransactionHandler
}

type TransactionWorker struct {
	Wait      *sync.WaitGroup
	Index     int            // 交易所在的 Index
	BlockInfo *BlockInfo     // 区块信息
	Recharges []*Transaction // 解析出来的充值记录
	Success   bool           // 解析成功状态
}

type BlockInfo struct {
	Number     int
	Hash       string
	Timestamp  int64
	TxSize     int
	Getter     TransactionGetter
	Difficulty int64
}

type MyBlockInfo struct {
	Number     int
	Hash       string
	Timestamp  int64
	TxSize     int
	Difficulty int64
}

type TransactionGetter func(index int) (tx interface{})

type BlockInfoHandler func(blockNumber uint) (info *BlockInfo, err error)

type TransactionHandler func(worker *TransactionWorker)

func (scheduler *SyncTransactionScheduler) Start() {

	array, err := DB().GetUnhandledBlock(scheduler.ChainType, scheduler.BlockLimit)
	if err != nil {

		common.Logger.Info("查询未解析区块任务失败:", err)
		return
	}

	if len(array) == 0 {

		return
	}

	scheduler.WaitGroup.Add(len(array))
	for _, block := range array {

		go scheduler.parseBlock(*block)

	}
}

//
//// 解析区块
func (scheduler *SyncTransactionScheduler) parseBlock(item BcwBlockNumber) {

	defer func() {
		scheduler.WaitGroup.Done()
	}()

	info, err := scheduler.BlockHandler(uint(item.Number))
	if err != nil {
		return
	}

	blocksInfo := MyBlockInfo{Number: info.Number, Hash: info.Hash, Timestamp: info.Timestamp, TxSize: info.TxSize, Difficulty: info.Difficulty}
	blocksInfo.InsertTransTable("block")

	// 分页解析（单次解析数量过多容易导致数据库连接耗尽)
	limit := 100
	page := info.TxSize / limit
	if info.TxSize%limit > 0 {

		page += 1
	}

	pageLimit := limit
	if info.TxSize < limit {
		pageLimit = info.TxSize
	}

	tasks := make([]*TransactionWorker, 0, info.TxSize)
	tasksInfo := make([]*TransInfo, 0, info.TxSize)
	for i := 0; i < page; i++ {

		var txWait sync.WaitGroup

		if i == page-1 && info.TxSize%pageLimit > 0 {

			txWait.Add(info.TxSize % pageLimit)
		} else {

			txWait.Add(pageLimit)
		}

		// DO Something
		for j := 0; j < pageLimit; j++ {

			index := i*pageLimit + j
			if index >= info.TxSize {

				break
			}
			worker := &TransactionWorker{
				Wait:      &txWait,
				Index:     index,
				BlockInfo: info,
			}
			go func(workertemp *TransactionWorker) {
				temp := ParseTransaction(workertemp)
				if len(temp.FromAddr) != 0 {
					temp.InsertTransTable("transaction")
					tasksInfo = append(tasksInfo, &temp)
				}
			}(worker)

			tasks = append(tasks, worker)
		}

		txWait.Wait()
	}

	//err = db.SaveEthTrans(tasksInfo)
	//if err != nil{
	//	log.Print(err)
	//}

	// 检查解析状态
	recharges, succeed := scheduler.checkParseStatus(tasks)

	// 全部解析成功
	if succeed {
		// 保存充值交易
		err = scheduler.blockParseSucceed(&item, recharges)
		if err != nil {
			common.Logger.Info("更新解析状态失败:", err)
		}
	}

}

func ParseTransaction(worker *TransactionWorker) TransInfo {
	defer func() {
		worker.Success = true
		worker.Wait.Done()
	}()

	tx := worker.BlockInfo.Getter(worker.Index).(*ethTypes.Transaction)
	receiver := tx.To()
	if receiver == nil {

		worker.Success = true
		common.Logger.Info(tx.Hash().Hex(), "：无接收地址")
		return TransInfo{}
	}

	receiverAddr := receiver.Hex()
	q := tx.Value()
	quantity := q.Uint64()

	asMessage, e := tx.AsMessage(ethTypes.LatestSignerForChainID(big.NewInt(int64(1))), tx.GasPrice())

	if e != nil {
		log.Println(e)
	}

	fromAddr := asMessage.From().Hex()

	temp := TransInfo{
		Hash:             tx.Hash().Hex(),
		BlockHash:        worker.BlockInfo.Hash,
		Nonce:            int64(tx.Nonce()),
		BlockNumber:      worker.BlockInfo.Number,
		TransactionIndex: worker.Index,
		FromAddr:         fromAddr,
		ToAddr:           receiverAddr,
		Value:            int64(quantity),
		Data:             string(tx.Data()),
		Gas:              int64(tx.Gas()),
		GasPrice:         tx.GasPrice().Int64(),
		BlockTimestamp:   worker.BlockInfo.Timestamp,
	}

	return temp

}

// 解析成功更新状态
func (scheduler *SyncTransactionScheduler) blockParseSucceed(block *BcwBlockNumber, recharges []*Transaction) error {

	db := DB()

	// 没有本平台的充值记录
	if len(recharges) == 0 {

		return block.ParseComplete()
	}

	_, err := db.Transaction(func(tx *DBConn) (i interface{}, e error) {

		e = tx.SaveTransactions(recharges, scheduler.ChainType)
		if e != nil {
			return
		}

		e = block.ParseCompleteInDB(tx)

		return
	})

	return err
}

// 检查解析结果
func (scheduler *SyncTransactionScheduler) checkParseStatus(tasks []*TransactionWorker) (array []*Transaction, succeed bool) {

	succeed = true
	//
	for _, worker := range tasks {

		// 出现解析失败
		if !worker.Success {

			succeed = false
			break
		}

		if worker.Recharges != nil && len(worker.Recharges) > 0 {

			array = append(array, worker.Recharges...)
		}
	}

	return
}

type SyncRechargeScheduler struct {
	ChainType        types.Chain    // 公链类型
	ChainBlockNumber GetBlockNumber // 获取区块链的当前高度
	Confirmations    uint64         // 不可逆高度差
}

type SyncFailedTxScheduler struct {
	Wait      *sync.WaitGroup
	ChainType types.Chain
	Handler   FailedTxHandler
	Limit     uint
}

type FailedTxHandler func(wait *sync.WaitGroup, transaction *FailedTransaction)

func (scheduler *SyncFailedTxScheduler) Start() {

	array, err := DB().FailedTxWithType(scheduler.ChainType, scheduler.Limit)
	if err != nil {

		common.Logger.Info("查询", scheduler.ChainType, "失败交易错误:", err)
		return
	}

	// 没有失败交易
	if array == nil || len(array) == 0 {

		return
	}

	scheduler.Wait.Add(len(array))

	for _, tx := range array {

		go scheduler.Handler(scheduler.Wait, tx)
	}
}

type SyncStatusScheduler struct {
	ChainType        types.Chain
	Confirmations    uint64
	CheckInterval    int64
	ChainBlockNumber GetBlockNumber
	TxStatus         StatusHandler
}

type StatusHandler func(txHash string) (exists bool, err error)
