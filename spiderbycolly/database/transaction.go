package database

import (
	"fmt"
	"help_center/spiderbycolly/common/types"
	"strings"
)

type Transaction struct {
	Id          uint64            `json:"id" gorm:"primary_key;autoincrement;"`
	ChainType   types.Chain       `json:"chain_type"`
	BlockNumber uint64            `json:"block_number"`
	TxHash      string            `json:"tx_hash" gorm:"type:varchar(128);unique_index:idx_tx_unique_txhash_receiver;"`
	Timestamp   int64             `json:"timestamp"`
	Payer       string            `json:"payer"`
	Receiver    string            `json:"receiver" gorm:"type:varchar(60);unique_index:idx_tx_unique_txhash_receiver;"`
	Quantity    uint64            `json:"quantity"`
	Symbol      string            `json:"symbol" gorm:"type:varchar(32);"`
	Memo        string            `json:"memo"`
	Status      TransactionStatus `json:"status" gorm:"type:enum('0', '1');default:'0';"`
}

type TransactionStatus string

const (
	TransactionStatusWait TransactionStatus = "0"
	TransactionStatusSend TransactionStatus = "1"
)

func (db *DBConn) SaveTransactions(array []*Transaction, chain types.Chain) error {

	if len(array) == 0 {

		return fmt.Errorf("参数错误")
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*9)

	for _, tx := range array {

		values = append(values, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		params = append(params, tx.BlockNumber, tx.TxHash, tx.Payer, tx.Receiver)
		params = append(params, tx.Quantity, tx.Symbol, tx.Memo, tx.Timestamp)
		params = append(params, chain)
	}

	format := "insert into bcw_transaction(block_number, tx_hash, payer, receiver, quantity, symbol, memo, `timestamp`, chain_type) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

func (db *DBConn) GetWaitTransaction(chain types.Chain) (array []*Transaction, err error) {

	err = db.Find(&array, "status = ? and chain_type = ?", TransactionStatusWait, chain).Error
	return
}

func (db *DBConn) UpdateTransactionStatus(maxId uint64, chain types.Chain) error {

	return db.Exec("update bcw_transaction set `status` = ? where id <= ? and chain_type = ?", TransactionStatusSend, maxId, chain).Error
}

type FailedTransaction struct {
	Id        uint64      `json:"id"`
	ChainType types.Chain `json:"chain_type"`
	Tx        string      `json:"tx"`
}

func (tx *FailedTransaction) Create(db *DBConn) error {

	return db.Create(tx).Error
}

func (tx *FailedTransaction) DeleteInDB(db *DBConn) error {

	return db.Delete(tx).Error
}

func (tx *FailedTransaction) Delete() error {

	return tx.DeleteInDB(DB())
}

func (db *DBConn) FailedTxWithType(chain types.Chain, limit uint) (array []*FailedTransaction, err error) {

	err = db.Limit(limit).Find(&array, "chain_type = ?", chain).Error
	return
}
