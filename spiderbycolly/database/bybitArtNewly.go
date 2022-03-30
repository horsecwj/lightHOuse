package database

import (
	"fmt"
	"help_center/spiderbycolly/spiderService/model"
	"strings"
)

func (db *DBConn) SaveBybitNewlyArt(array []model.BybitNewlyArticle) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?,?,?,?)")
		params = append(params, address.Title, address.OverView)
		params = append(params, address.Article, address.Link)
		params = append(params, address.Time, address.Timestamp, address.Articletext, address.Pic)
	}

	format := "insert into bybit_newly_article (title,over_view,article,link,time,timestamp,articletext,pic) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

func (db *DBConn) GetBybitNewlyArt() ([]*model.BybitNewlyArticle, error) {
	var addr []*model.BybitNewlyArticle
	err := db.Model(&addr).Debug().Order("timestamp desc limit 1").Scan(&addr).Error
	return addr, err
}

// 根据symbol删除记录
func (db *DBConn) DeletebitNewlyArtCoin(symbol string) error {

	return db.Where("id = ?", symbol).Delete(&model.BybitNewlyArticle{}).Error
}
