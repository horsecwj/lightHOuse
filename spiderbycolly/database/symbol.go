package database

import "help_center/spiderbycolly/common/types"

var CacheSymbol = map[string]*Symbol{}

const (
	ETHSymbol  string = "ETH"
	BTCSymbol  string = "BTC"
	USDTSymbol string = "USDT"
	TRXSymbol  string = "TRX"
	BNBSymbol  string = "BNB"
)

type Symbol struct {
	ID                  uint        `json:"id"`
	Chain               types.Chain `json:"chain"`
	Symbol              string      `json:"symbol" gorm:"type:varchar(32);unique_index;"`
	Token               string      `json:"token" gorm:"comment:'代币名称（可重复）';"`
	Contract            string      `json:"contract,omitempty" gorm:"comment:'合约名称，可为空';"`
	Method              string      `json:"method,omitempty" gorm:"comment:'以太坊ERC20转账方法hex';"`
	Decimal             uint        `json:"decimal" gorm:"comment:'代币精度';"`
	ChainDecimal        uint        `json:"chain_decimal" gorm:"comment:'代币在链上精度';"`
	MinRechargeQuantity uint64      `json:"min_recharge_quantity" gorm:"comment:'最小充值数量'"`
	BalanceMethod       string      `json:"balance_method" gorm:"comment:'以太坊合约ERC20查询余额方法'"`
	GasLimit            uint64      `json:"gas_limit" gorm:"comment:'ERC20 合约执行时常';default:21000;"`
	Flag                string      `json:"flag" gorm:"type:enum('0', '1', '2');comment:'0: 未启用, 1: 正常状态，2: 已作废'"`
}

// 保存
func (symbol *Symbol) Create(db *DBConn) error {

	return db.Create(symbol).Error
}

// 获取当前本地精度与链上精度差
func (symbol *Symbol) DiffDecimal() uint {
	if symbol.ID == 0 {
		return 0
	}
	return symbol.ChainDecimal - symbol.Decimal
}

// 根据symbol删除记录
func (db *DBConn) DeleteWithSymbol(symbol string) error {

	return db.Where("symbol = ?", symbol).Delete(&Symbol{}).Error
}

// 获取当前所有的symbol
func (db *DBConn) GetAllSymbol() (symbols []*Symbol, err error) {

	err = db.Find(&symbols).Error
	return
}

// 根据symbol查询对象
func (db *DBConn) GetSymbol(symbol string) (item *Symbol, err error) {

	// 优先读取缓存
	if CacheSymbol[symbol] != nil {

		return CacheSymbol[symbol], nil
	}

	var s Symbol
	err = db.Debug().First(&s, "symbol = ?", symbol).Error
	if err != nil {

		return
	}

	item = &s

	// 存入缓存
	CacheSymbol[symbol] = item
	return
}

// 根据ERC20 代币方法查询
func (db *DBConn) GetSymbolByContract(contract string) (item *Symbol, err error) {

	var symbol Symbol
	err = db.Where("contract = ?", contract).First(&symbol).Error
	if err != nil {

		return
	}

	item = &symbol
	return
}
