package database

import (
	"encoding/base64"
	"fmt"
	"help_center/spiderbycolly/spiderService/model"
	"strings"
)

// 批量保存
func (db *DBConn) SaveSlateArt(array []model.SlateArticle) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?,?,?,?)")
		params = append(params, address.Title, address.OverView)
		params = append(params, address.Article, address.Link)
		params = append(params, address.Time, address.Timestamp, address.Articletext, base64.StdEncoding.EncodeToString([]byte(address.Pic)))
	}

	format := "insert into slate_article (title,over_view,article,link,time,timestamp,articletext,pic) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// 获取一个未使用的地址
func (db *DBConn) GetSlateArt() ([]*model.SlateArticle, error) {
	var addr []*model.SlateArticle
	err := db.Model(&addr).Debug().Order("timestamp desc limit 1").Scan(&addr).Error
	return addr, err
}
func (db *DBConn) SlateArtLink() ([]string, error) {

	sql := "SELECT slate_article.link links from slate_article"
	var links []string

	err := db.Raw(sql).Pluck("links", &links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

// 获取多个未使用的地址
func (db *DBConn) GetManySlateArt() (map[string]bool, error) {
	resLink, err := db.SlateArtLink()
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
