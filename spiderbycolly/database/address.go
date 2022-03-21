package database

import (
	"fmt"
	"help_center/spiderbycolly/common/types"
	"strings"
)

type Address struct {
	Id      uint32        `json:"id" gorm:"primary key;autoincrement;"`
	Address string        `json:"address"`
	Chain   types.Chain   `json:"chain" gorm:"type:enum('BTC', 'ETH','Tron');default:'BTC';"`
	Status  AddressStatus `json:"status" gorm:"type:enum('0', '1', '2');default:'0';"`

	Balance     float64 `json:"balance"`
	UsdtBalance float64 `json:"usdt_balance"`
}

type AddressStatus string

const (
	AddressStatusUnused  AddressStatus = "0"
	AddressStatusUsed    AddressStatus = "1"
	AddressStatusUnknown AddressStatus = "2"
)

// NewAddress 根据地址跟类型创建对象
func NewAddress(address string, chain types.Chain) *Address {

	return &Address{Address: address, Chain: chain}
}

// Create 单个地址保存
func (address *Address) Create(db *DBConn) error {

	return db.Create(address).Error
}

// UsedInDB 地址已使用
func (address *Address) UsedInDB(db *DBConn) error {

	return db.Model(address).Update("status", AddressStatusUsed).Error
}

// Used 在普通数据库连接中更新
func (address *Address) Used() error {

	return address.UsedInDB(DB())
}

// SaveAddresses 批量保存地址
func (db *DBConn) SaveAddresses(array []*Address) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*2)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?, ?)")
		params = append(params, address.Address, address.Chain)
	}

	// 拼接SQL
	format := "insert into bcw_address (address, chain) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// UnusedAddrWithType 查询指定类型的地址未使用数量
func (db *DBConn) UnusedAddrWithType(chain types.Chain) (count uint, err error) {

	err = db.Model(&Address{}).Where("chain = ? and status = ?", chain, AddressStatusUnused).Count(&count).Error
	return
}

// GetUnusedAddress 获取一个未使用的地址
func (db *DBConn) GetUnusedAddress(chain types.Chain) (*Address, error) {

	var addr Address
	err := db.Model(&addr).Order("id asc limit 1 for update").Where("chain = ? and status = ?", chain, AddressStatusUnused).Scan(&addr).Error
	return &addr, err
}

// IsExistsAddress 查询是否存在
func (db *DBConn) IsExistsAddress(addr string) (isExists bool, err error) {
	var count uint64
	err = db.Model(&Address{}).Where("address = ?", addr).Count(&count).Error
	if err != nil {

		return
	}

	if count > 0 {

		isExists = true
	}

	return
}

// GetAddressesLimit 批量查询已使用地址
func (db *DBConn) GetAddressesLimit(id, limit uint64, chain types.Chain) (array []*Address, err error) {

	err = db.Limit(limit).Find(&array, "id > ? and chain = ?", id, chain).Error

	return
}

func (db *DBConn) GetAddressesBalanceLimit(id, limit uint64, chain types.Chain) (array []*Address, err error) {

	err = db.Limit(limit).Find(&array, "id > ? and chain = ? and balance is null", id, chain).Error

	return
}

func (db *DBConn) UpdateBalance(balance float64, usdtBalance float64, address string) error {

	params := map[string]interface{}{

		"balance":      balance,
		"usdt_balance": usdtBalance,
	}

	return db.Model(&Address{}).Where("address = ?", address).Updates(params).Error
}
