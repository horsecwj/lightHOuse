package database

import (
	"encoding/base64"
	"fmt"
	"help_center/spiderbycolly/spiderService/model"
	"strings"
)

// 批量保存
func (db *DBConn) SaveBybitHighLightArt(array []model.BybitArticle) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?,?,?,?)")

		params = append(params, address.Title, address.OverView)
		params = append(params, address.Article, address.Link)
		params = append(params, address.Time, address.Timestamp, address.Articletext, base64.StdEncoding.EncodeToString([]byte(address.Pic)))
	}

	// 拼接SQL
	format := "insert into bybit_article (title,over_view,article,link,time,timestamp,articletext,pic) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// 获取一个
func (db *DBConn) GetBybitArt() ([]model.BybitArticle, error) {
	var addr []model.BybitArticle
	err := db.Model(&addr).Debug().Order("timestamp  limit 1").Scan(&addr).Error
	return addr, err
}

// 根据symbol删除记录
func (db *DBConn) DeleteBybitArt() error {

	return db.Delete(&model.BybitArticle{}).Error
}

func (db *DBConn) BybitNewArtLink() ([]string, error) {

	sql := "SELECT bybit_newly_article.link links from bybit_newly_article"
	var links []string

	err := db.Raw(sql).Pluck("links", &links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

// 获取多个未使用的地址
func (db *DBConn) GetManyBybitNewArt() (map[string]bool, error) {
	resLink, err := db.BybitNewArtLink()
	if err != nil {
		return nil, err
	}
	var linkMap map[string]bool
	linkMap = make(map[string]bool, len(resLink)+1)
	for _, item := range resLink {
		linkMap[item] = true
	}
	return linkMap, err
}
