package database

import (
	"crypto/sha256"

	"github.com/jinzhu/gorm"
	//"more.top/blockchain-server-wallet/common/kms"
	"help_center/spiderbycolly/common/types"
)

type Transfer struct {
	ID          uint64           `json:"-"`
	CustomId    string           `json:"custom_id" gorm:"type:varchar(128);unique_index;comment:'请求端的唯一ID';not null;"`
	TxHash      string           `json:"tx_hash" gorm:"comment:'链上交易ID';"`
	Symbol      string           `json:"symbol" gorm:"comment:'对应Symbol表';"`
	FromAddress string           `json:"from_address" gorm:"comment:'转出地址';"`
	ToAddress   string           `json:"to_address" gorm:"comment:'收款地址';"`
	Memo        string           `json:"memo" gorm:"comment:'填写的memo (EOS独有）';"`
	Quantity    uint64           `json:"quantity" gorm:"comment:'转账数量（乘以symbol精度后的值）';"`
	Fee         uint64           `json:"fee" gorm:"comment:'自定义转账手续费'"`
	Timestamp   int64            `json:"timestamp"`
	Signature   string           `json:"signature" gorm:"comment:'数据签名(KMS)';type:text;"`
	Summed      bool             `json:"summed" gorm:"type tinyint(1) default 0;not null;"`
	Message     string           `json:"message" gorm:"type:text;"` // 错误消息
	Flag        TransferFlag     `json:"flag" gorm:"type:enum('0', '1', '2', '3');default:'0';comment:'0: 待发送，1: 已发出, 2: 不可逆确认完成, 3: 失败（被回滚）';"`
	SyncFlag    TransferSyncFlag `json:"sync_flag" gorm:"type:enum('0', '1', '2', '3');default:'0';"`
	FeeAddress  string           `json:"fee_address" gorm:"type:varchar(256);comment:'手续费地址'"`
}

type TransferFlag string

const (
	TransferWait    TransferFlag = "0"
	TransferSend    TransferFlag = "1"
	TransferSuccess TransferFlag = "2"
	TransferFail    TransferFlag = "3"
)

type TransferSyncFlag string

const (
	TransferSyncWait    TransferSyncFlag = "0"
	TransferSyncSend    TransferSyncFlag = "1"
	TransferSyncSuccess TransferSyncFlag = "2"
	TransferSyncFail    TransferSyncFlag = "3"
)

// CidHash 获取CID sha256
func (t *Transfer) CidHash() []byte {

	hash := sha256.New()
	hash.Write([]byte(t.CustomId))
	return hash.Sum(nil)
}

// CreateInDB 保存
func (t *Transfer) CreateInDB(db *DBConn) error {

	return db.Create(t).Error
}

func (t *Transfer) Create() error {

	return t.CreateInDB(DB())
}

// SendInDB 已发送
func (t *Transfer) SendInDB(db *DBConn, txHash string, timestamp int64) error {

	return db.Model(t).Updates(Transfer{
		TxHash:    txHash,
		Flag:      TransferSend,
		Timestamp: timestamp,
	}).Error
}

func (t *Transfer) Send(txHash string, timestamp int64) error {

	return t.SendInDB(DB(), txHash, timestamp)
}

// FailedInDB 已失败
func (t *Transfer) FailedInDB(db *DBConn, message string) error {

	params := map[string]interface{}{
		"flag":    TransferFail,
		"message": message,
	}
	return db.Model(t).Updates(params).Error
}

func (t *Transfer) Failed(message string) error {

	return t.FailedInDB(DB(), message)
}

// SucceedInDB 已确认
func (t *Transfer) SucceedInDB(db *DBConn) error {

	return db.Model(t).Update("flag", TransferSuccess).Error
}

func (t *Transfer) Succeed() error {

	return t.SucceedInDB(DB())
}

// SyncInDB 同步状态
func (t *Transfer) SyncInDB(db *DBConn) error {

	return db.Exec("update bcw_transfer set sync_flag = flag where id = ?", t.ID).Error
}

func (t *Transfer) Sync() error {

	return t.SyncInDB(DB())
}

// GetSyncTransfer 查询未同步状态的转账
func (db *DBConn) GetNeedSyncTransfer() (array []*Transfer, err error) {

	err = db.Find(&array, "sync_flag != flag").Error
	return
}

// GetUnConfirmTransfer 查询已发送状态的交易
func (db *DBConn) GetUnConfirmTransfer(chain types.Chain) (array []*Transfer, err error) {

	return db.getTransferByChainStatus(chain, TransferSend)
}

// GetTransferByCId 查询提现信息
func (db *DBConn) GetTransferByCId(cId string) (transfer *Transfer, err error) {

	var t Transfer
	err = db.Where("custom_id = ?", cId).First(&t).Error

	if err != nil {

		return
	}

	transfer = &t
	return

}

// IsTransferHash 根据TxHash查询是否存在提现记录
func (db *DBConn) IsTransferHash(hash string) (exists bool, err error) {

	err = db.Find(&Transfer{}, "tx_hash = ?", hash).Error

	if err != nil {

		exists = !gorm.IsRecordNotFoundError(err)
		if !exists {

			err = nil
		}
	}

	return
}

// GetNeedSendTransfersByChain 根据公链类型查询
func (db *DBConn) GetNeedSendTransfersByChain(chain types.Chain) (array []*Transfer, err error) {

	return db.getTransferByChainStatus(chain, TransferWait)
}

// 根据公链类型，状态查询数据
func (db *DBConn) getTransferByChainStatus(chain types.Chain, flag TransferFlag) (array []*Transfer, err error) {

	err = db.Table("bcw_transfer").Joins("inner join bcw_symbol on bcw_transfer.symbol = bcw_symbol.symbol").Where("bcw_symbol.chain = ? ", chain).Scan(&array).Error
	return
}

// GetTransferByHash 根据txHash查询
func (db *DBConn) GetTransferByHash(txHash string) (*Transfer, error) {

	var transfer Transfer
	err := db.First(&transfer, "tx_hash = ?", txHash).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {

		return nil, err
	}

	return &transfer, err
}

// GetETHTransfer 查询手续费
func (db *DBConn) GetETHTransfer(limit int) (array []*Transfer, err error) {

	err = db.Limit(limit).Find(&array, "symbol like '%ETH%' and fee is null and tx_hash is not null").Error
	return
}

func (db *DBConn) UpdateETHFees(txHash string, fee float64) error {

	return db.Model(&Transfer{}).Where("tx_hash = ?", txHash).Update("fee", fee).Error
}
