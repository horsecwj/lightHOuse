package database

import (
	"fmt"
	"help_center/spiderbycolly/common/types"
	"strings"
)

type BcwBlockNumber struct {
	Number uint64           `json:"number" gorm:"primary_key;"`
	Chain  types.Chain      `json:"chain" gorm:"primary_key;"`
	Status BlockParseStatus `json:"status" gorm:"type:enum('0', '1');default:'0';"`
}

type TransInfo struct {
	Hash             string
	BlockHash        string
	Nonce            int64
	BlockNumber      int
	TransactionIndex int
	FromAddr         string
	ToAddr           string
	Value            int64
	Gas              int64
	GasPrice         int64
	BlockTimestamp   int64
	Data             string
}

type BlockParseStatus string

const (
	BlockParseStatusWait     BlockParseStatus = "0"
	BlockParseStatusComplete BlockParseStatus = "1"
)

func (block *BcwBlockNumber) ParseCompleteInDB(db *DBConn) error {

	return db.Model(&block).Update("status", BlockParseStatusComplete).Error
}

func (block *BcwBlockNumber) ParseComplete() error {

	return block.ParseCompleteInDB(DB())
}

// MaxBlockNumber 查询当前已获取的最大区块高度
func (db *DBConn) MaxBlockNumber(chain types.Chain) (uint, error) {

	sql := "select ifnull(max(number), 0) number from bcw_block_number where chain = ?"
	var numbers []uint

	err := db.Raw(sql, chain).Pluck("number", &numbers).Error
	if err != nil {
		return 0, err
	}

	return numbers[0], nil
}

// MaxBlockNumber 查询当前已获取的最大区块高度
func (db *DBConn) MaxUnHandleBlockNumber(chain types.Chain) (uint, error) {

	sql := "select ifnull(max(number), 0) number from bcw_block_number where chain = ? and status = '0' "
	var numbers []uint

	err := db.Raw(sql, chain).Pluck("number", &numbers).Error
	if err != nil {
		return 0, err
	}

	return numbers[0], nil
}
func (db *DBConn) CountUnHandleBlockNumber(chain types.Chain) (uint, error) {

	sql := "select COUNT(*) from bcw_block_number where chain = ? and status = '0' "
	var numbers []uint

	err := db.Raw(sql, chain).Pluck("COUNT(*)", &numbers).Error
	if err != nil {
		return 0, err
	}

	return numbers[0], nil
}

// SaveBlockNumbers 保存区块解析任务
func (db *DBConn) SaveBlockNumbers(start uint, end uint, chain types.Chain) error {

	if end < start {

		return fmt.Errorf("参数错误")
	}

	length := end - start
	if length == 0 {

		length = 1
	}

	values := make([]string, 0, length)
	params := make([]interface{}, 0, length)

	for number := end; number >= start; number-- {

		if start != end && number == start {

			break
		}

		values = append(values, "(?, ?)")
		params = append(params, number, chain)
	}

	format := "insert into bcw_block_number(number, chain) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Debug().Exec(sql, params...).Error
}

// GetUnhandledBlock 查询未解析的区块
func (db *DBConn) GetUnhandledBlock(chain types.Chain, limit uint) (array []*BcwBlockNumber, err error) {

	err = db.Limit(limit).Debug().Find(&array, "status = ? and chain = ?", BlockParseStatusWait, chain).Error
	return
}
