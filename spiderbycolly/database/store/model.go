package store

import "time"

type AggTradeInfo struct {
	Price     float64   `json:"price"`     //成交价格
	Vol       float64   `json:"tra_vol"`   //成交数量
	TraTime   time.Time `json:"tra_time"`  //成交时间
	Market    bool      `json:"market"`    //是否为主动卖出单
	Timestamp time.Time `json:"timestamp"` //入库时间
}
