package database

import (
	"fmt"
	"help_center/spiderbycolly/spiderService/model"

	"strings"
)

//type TopCkoGameFi struct {
//	Coin      string `json:"coin" gorm:"primary key;"`
//	Price     string `json:"price"`
//	OneDay    string `json:"one_day"`
//	OneWeek   string `json:"one_week"`
//	DayVolume string `json:"day_volume"`
//	MktCap    string `json:"mkt_cap"`
//	LastWeek  string `json:"last_week"`
//}

// 批量保存
func (db *DBConn) SaveTopCmkGameFi(array []*model.TopCmkGameFi) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*6)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?, ?)")
		params = append(params, address.ID)
		params = append(params, address.Coin, address.Price)
		params = append(params, address.OneDay)
		params = append(params, address.DayVolume)
		params = append(params, address.GameFi)
	}

	// 拼接SQL
	format := "insert into top_cmk_game_fi (id,coin, price,one_day,day_volume,game_fi) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// 批量保存
func (db *DBConn) SaveTopCmkGameFiLosers(array []*model.TopCmkGameFi) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*6)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?, ?)")
		params = append(params, address.ID)
		params = append(params, address.Coin, address.Price)
		params = append(params, address.OneDay)
		params = append(params, address.DayVolume)
		params = append(params, address.GameFi)
	}

	// 拼接SQL
	format := "insert into top_cmk_game_fi_losers (id,coin, price,one_day,day_volume,game_fi) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// 根据symbol删除记录
func (db *DBConn) DeleteTopCmkGameFiWithCoin(symbol string) error {

	return db.Where("coin = ?", symbol).Delete(&model.TopCmkGameFi{}).Error
}

// 根据symbol删除记录
func (db *DBConn) DeleteTopCmkGameFi() error {

	return db.Delete(&model.TopCmkGameFi{}).Error
}

// 根据symbol删除记录
func (db *DBConn) DeleteTopCmkGameFiLosers() error {

	return db.Delete(&model.TopCmkGameFiLosers{}).Error
}

//type BybitArticle struct {
//	Title      string `json:"title"`
//	OverView     string `json:"over_view"`
//	Article    string `json:"article"`
//	Time string `json:"time"`
//}
//type SlateArticle struct {
//	Title      string `json:"title"`
//	OverView     string `json:"over_view"`
//	Article    string `json:"article"`
//	Time string `json:"time"`
//}
