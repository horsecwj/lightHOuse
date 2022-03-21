package database

import (
	"fmt"
	"help_center/spiderbycolly/spiderService/model"
	"strings"
)

func (db *DBConn) SaveEthTrans(array []*TransInfo) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)
	for _, address := range array {
		values = append(values, "(?, ?, ?,  ?, ?,?,  ?,?,?, ?,?,?)")
		params = append(params, address.Hash, address.BlockHash)
		params = append(params, address.Nonce, address.BlockNumber)
		params = append(params, address.TransactionIndex, address.FromAddr, address.Value)
		params = append(params, address.Gas, address.GasPrice, address.BlockTimestamp)
		params = append(params, address.Data, address.ToAddr)
	}

	format := "insert into trans_info (hash,block_hash,nonce,block_number,transaction_index,from_addr,value,gas,gas_price,block_timestamp,data,to_addr) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

func (db *DBConn) SaveBlocks(array []*MyBlockInfo) error {

	if len(array) == 0 {

		return fmt.Errorf("参数错误")
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*9)

	for _, tx := range array {

		values = append(values, "(?, ?, ? , ?, ?)")
		params = append(params, tx.Number)
		params = append(params, tx.Hash, tx.TxSize, tx.Timestamp)
		params = append(params, tx.Difficulty)
	}

	format := "insert into my_block_info(number,hash, tx_size,timestamp, difficulty) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

func (db *DBConn) GetEthTrans() ([]*model.BybitNewlyArticle, error) {
	var addr []*model.BybitNewlyArticle
	err := db.Model(&addr).Debug().Order("timestamp desc limit 1").Scan(&addr).Error
	return addr, err
}

// 根据symbol删除记录
func (db *DBConn) DeleteEthTrans(symbol string) error {
	return db.Where("id = ?", symbol).Delete(&model.BybitNewlyArticle{}).Error
}
