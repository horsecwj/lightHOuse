package data

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func ArticleAdd(a *Article) error {
	t := time.Now()
	if a.Id == 0 {
		a.Id = t.UnixMilli()
	}
	if a.Cover == "" {
		a.Cover = "/upload/1646706814654.png"
	}
	var row GameArticle
	if a.GameId != 0 {
		row.ArticleId = a.Id
		row.GameId = a.GameId
		err := GetDbCli().Session(&gorm.Session{}).Table("game_article").Create(&row).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("articles").Preload("Label").Create(&a).Error
}

func ArticleDelete(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{})
	err := tx.Table("article_label").Delete(ArticleLabel{}, "article_id = ?", id).Error
	if err != nil {
		log.Println(err.Error())
	}
	err = tx.Table("game_article").Delete(GameArticle{}, "article_id = ?", id).Error
	if err != nil {
		log.Println(err.Error())
	}
	return tx.Table("articles").Delete(Article{}, "id = ?", id).Error
}

func ArticleUpdate(a *Article) error {
	tx := GetDbCli().Session(&gorm.Session{})
	var row GameArticle
	if a.GameId != 0 {
		row.ArticleId = a.Id
		row.GameId = a.GameId
		err := tx.Table("game_article").Delete(GameArticle{}, "article_id = ?", a.Id).Create(&row).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
	if a.Label != nil {
		tx.Table("article_label").Delete(ArticleLabel{}, "article_id = ?", a.Id)
	}
	if a.Status == 0 {
		tx.Table("articles").Where("id = ?", a.Id).Update("status", 0)
	}
	if a.Hot == 0 {
		tx.Table("articles").Where("id = ?", a.Id).Update("hot", 0)
	}
	return tx.Table("articles").Where("id = ?", a.Id).Omit("created").Preload("Label").Updates(&a).Error
}

func MoreArticleUpdate(d []MoreArticle) error {
	tx := GetDbCli().Session(&gorm.Session{})
	for i := range d {
		if d[i].Id != 0 {
			err := tx.Table("articles").Where("id = ?", d[i].Id).Updates(Article{CateId: d[i].CateId}).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *ArticleQuery) ArticleSearch(adm bool) interface{} {
	var list = make([]article, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles").Order("updated desc").Preload("Label")
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)
	}
	if a.Status != 0 {
		tx = tx.Not("status = ?", 2)
	}
	if a.CateId != 0 {
		tx = tx.Where("cate_id = ?", a.CateId)
	}
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	if a.Hot == 1 {
		tx = tx.Not("hot = ?", 0)
	}
	if !adm {
		type article struct {
			Id       int64   `json:"id"`
			Lang     string  `json:"lang"`
			CateId   int64   `json:"cate_id"`
			Title    string  `json:"title"`
			Summary  string  `json:"summary"`
			Cover    string  `json:"cover"`
			RichText string  `json:"rich_text"`
			Label    []Label `json:"label" gorm:"many2many:article_label"`
			Hot      int     `json:"hot"`
			Updated  string  `json:"updated"`
		}
		var result = make([]article, 0, a.PageSize)
		if a.Id == 0 {
			tx.Where("status = ?", 2)
		}
		err := tx.Find(&result).Error
		if err != nil {
			log.Println(err.Error())
		}
		for i := range result {
			result[i].Updated = result[i].Updated[:10]
		}
		return result
	} else {
		err := tx.Find(&list).Error
		if err != nil {
			log.Println(err.Error())
		}
		return list
	}
}

func (a *ArticleQuery) ArticleCount() int {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles")
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)
	}
	if a.Status != 0 {
		tx = tx.Not("status = ?", 2)
	}
	if a.CateId != 0 {
		tx = tx.Where("cate_id = ?", a.CateId)
	}
	if a.Hot == 1 {
		tx = tx.Not("hot = ?", 0)
	}
	err := tx.Count(&count).Where("status = ?", 2).Error
	if err != nil {
		log.Println(err.Error())
	}
	strCount := strconv.FormatInt(count, 10)
	intCount, err := strconv.Atoi(strCount)
	if err != nil {
		log.Println(err.Error())
	}
	return intCount
}

func ArticleMatch(subStr string, user bool) (interface{}, int) {
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles").Order("created desc, id").Preload("Label")
	tx = tx.Where("title like ? ", "%"+subStr+"%")
	tx = tx.Limit(30)
	if user {
		tx = tx.Where("status = ?", 2)
	}
	var result []article
	err := tx.Find(&result).Error
	if err != nil {
		log.Println(err.Error())
	}
	return result, len(result)
}

func (a *ArticleQuery) LikeArticle() interface{} {
	var (
		row  []likeArticle
		data []likeArticle
	)
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles").
		Where("cate_id = ?", a.CateId).Where("status = ?", 2).Order("id desc")
	err := tx.Find(&row).Error
	if err != nil {
		log.Println(err.Error())
	}
	for i := range row {
		row[i].Updated = row[i].Updated[:10]
	}
	i := 1
	for i = range row {
		if row[i].Id == a.Id {
			break
		}
	}
	fmt.Println(len(row), i)
	if i == 0 {
		if len(row) <= 4 {
			return row
		}
		return row[1:5]
	}
	if i == 1 {
		if len(row) <= 4 {
			return row
		}
		data = append(data, row[0], row[2], row[3], row[4])
		return data
	}
	if i+1 == len(row) {
		if len(row) <= 4 {
			return row
		}
		return row[len(row)-4:]
	}
	if i+1 == len(row)-1 {
		if len(row) <= 4 {
			return row
		}
		data = append(data, row[len(row)-5], row[len(row)-4], row[len(row)-3], row[len(row)-1])
		return data
	}
	data = append(data, row[i-2], row[i-1], row[i+1], row[i+2])
	return data
}

func OutArticleAdd() error {
	var (
		row   []OutArticle
		bybit []OutArticle
		cover string
	)
	err := GetDbCli().Session(&gorm.Session{}).Table("slate_article").Find(&row).Error
	if err != nil {
		log.Println(err)
	}
	err = GetDbCli().Session(&gorm.Session{}).Table("bybit_article").Find(&bybit).Error
	if err != nil {
		log.Println(err)
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles")
	for i := 0; i < len(row); i++ {
		if row[i].Pic == "" {
			cover = "/upload/1646706814654.png"
		} else {
			cover = row[i].Pic
		}
		if VerificationArticle(row[i].Timestamp) != nil {
			link := fmt.Sprintf("<p>Contents Sourced from: %s</p>", row[i].Link)
			row[i].Article = row[i].Article + link
			row[i].Articletext = row[i].Articletext + "\n" + row[i].Link
			timeStr := time.Unix(row[i].Timestamp, 0).Format("2006-01-02 15:04:05")
			err := tx.Create(article{Id: row[i].Timestamp, Title: row[i].Title, Lang: "cn", CateId: 1645582941827, Cover: cover, Summary: row[i].OverView, Markdown: row[i].Article, RichText: row[i].Articletext, Created: timeStr, Updated: timeStr}).Error
			if err != nil {
				log.Println(err)
			}
		}
	}
	tx = GetDbCli().Session(&gorm.Session{}).Table("articles")
	for i := 0; i < len(bybit); i++ {
		if bybit[i].Pic == "" {
			cover = "/upload/1646706814654.png"
		} else {
			cover = row[i].Pic
		}
		if VerificationArticle(row[i].Timestamp) != nil {
			timeStr := time.Unix(bybit[i].Timestamp, 0).Format("2006-01-02 15:04:05")
			err := tx.Create(article{Id: row[i].Timestamp, Title: bybit[i].Title, Lang: "cn", CateId: 1645582941827, Cover: cover, Summary: bybit[i].OverView, Markdown: bybit[i].Article, RichText: bybit[i].Articletext, Created: timeStr, Updated: timeStr}).Error
			if err != nil {
				log.Println(err)
			}
		}
	}
	return err
}

func (a *ArticleQuery) Course(Video bool, Image bool) interface{} {
	tx := GetDbCli().Session(&gorm.Session{})
	var (
		video []DelQuery
		image []DelQuery
		arr   []int64
	)
	if Video {
		err := tx.Table("categories").Where("parent_id = ?", 1646124402108).Find(&video).Error
		if err != nil {
			log.Println(err)
		}
	}
	if Image {
		err := tx.Table("categories").Where("parent_id = ?", 1646104648579).Find(&image).Error
		if err != nil {
			log.Println(err)
		}
	}
	for i := 0; i < len(video); i++ {
		arr = append(arr, video[i].Id)
	}
	for i := 0; i < len(image); i++ {
		arr = append(arr, image[i].Id)
	}
	tx = tx.Table("articles").Where("cate_id in ?", arr).Preload("Category").Order("updated desc")
	if a.CateId != 0 {
		tx = tx.Where("cate_id = ?", a.CateId)
	}
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	var row []CourseBanner
	err := tx.Find(&row).Error
	if err != nil {
		log.Println(err.Error())
	}
	for i := range row {
		row[i].Updated = row[i].Updated[:10]
	}
	return row
}

func (a *ArticleQuery) CourseCount(Video bool, Image bool) int {
	tx := GetDbCli().Session(&gorm.Session{})
	var (
		course []DelQuery
		arr    []int64
		count  int64
	)
	if Video {
		err := tx.Table("categories").Where("parent_id = ?", 1646124402108).Find(&course).Error
		if err != nil {
			log.Println(err)
		}
	} else {
		err := tx.Table("categories").Where("parent_id = ?", 1646104648579).Find(&course).Error
		if err != nil {
			log.Println(err)
		}
	}
	for i := 0; i < len(course); i++ {
		arr = append(arr, course[i].Id)
	}
	tx = tx.Table("articles").Where("cate_id in ?", arr)
	if a.CateId != 0 {
		tx = tx.Where("cate_id = ?", a.CateId)
	}
	err := tx.Count(&count).Error
	if err != nil {
		log.Println(err.Error())
	}
	strCount := strconv.FormatInt(count, 10)
	intCount, err := strconv.Atoi(strCount)
	if err != nil {
		log.Println(err.Error())
	}
	return intCount
}
