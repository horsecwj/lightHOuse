package database

import (
	"fmt"
	"help_center/spiderbycolly/spiderService/model"
	"strings"
)

// 批量保存
func (db *DBConn) SaveTopCkoGameFi(array []*model.TopCkoGameFi) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*8)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?, ?, ?,?)")
		params = append(params, address.ID)
		params = append(params, address.Coin, address.Price)
		params = append(params, address.OneDay, address.OneWeek)
		params = append(params, address.DayVolume, address.MktCap)
		params = append(params, address.GameFi)
	}

	// 拼接SQL
	format := "insert into top_cko_game_fi (id,coin, price,one_day,one_week,day_volume,mkt_cap,game_fi) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// 根据symbol删除记录
func (db *DBConn) DeleteTopCkoGameFiWithCoin(symbol string) error {

	return db.Where("coin = ?", symbol).Delete(&model.TopCkoGameFi{}).Error
}

// 根据symbol删除记录
func (db *DBConn) DeleteTopCkoGameFi() error {

	return db.Delete(&model.TopCkoGameFi{}).Error
}

type BybitArticle struct {
	Id       string `json:"id" gorm:"autoincrement;"`
	Title    string `json:"title" gorm:"primary key;`
	OverView string `json:"over_view"`
	Article  string `json:"article"`
	Time     string `json:"time"`
}

type SlateArticle struct {
	Id       string `json:"id" gorm:"autoincrement;"`
	Title    string `json:"title" gorm:"primary key;`
	OverView string `json:"over_view"`
	Article  string `json:"article"`
	Time     string `json:"time"`
}
