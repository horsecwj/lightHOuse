package data

import (
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

func (a *ArticleQuery) ArticleSearch(adm bool) interface{} {
	var list = make([]article, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles").Order("id desc").Preload("Label")
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
			Id       int64     `json:"id"`
			Lang     string    `json:"lang"`
			CateId   int64     `json:"cate_id"`
			Title    string    `json:"title"`
			Summary  string    `json:"summary"`
			Cover    string    `json:"cover"`
			RichText string    `json:"rich_text"`
			Label    []Label   `json:"label" gorm:"many2many:article_label"`
			Hot      int       `json:"hot"`
			Created  time.Time `json:"created"`
		}
		var result = make([]article, 0, a.PageSize)
		tx.Where("status = ?", 2)
		err := tx.Find(&result).Error
		if err != nil {
			log.Println(err.Error())
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

func ArticleMatch(subStr string) (interface{}, int) {
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles").Order("created desc, id")
	tx = tx.Where("title like ? ", "%"+subStr+"%")
	tx = tx.Limit(30)
	type Result struct {
		Id       int64  `json:"id"`
		Lang     string `json:"lang"`
		CateId   int64  `json:"cate_id"`
		Title    string `json:"title"`
		Summary  string `json:"summary"`
		RichText string `json:"rich_text"`
		Uri      string `json:"uri"`
	}
	var result []Result
	tx = tx.Select("id", "lang", "cate_id", "title", "summary", "rich_text", "uri")
	err := tx.Scan(&result).Error
	if err != nil {
		log.Println(err.Error())
	}
	return result, len(result)
}

func (a *ArticleQuery) LikeArticle() interface{} {
	var row []likeArticle
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles").
		Where("cate_id = ?", a.CateId).Where("status = ?", 2).Not("id = ?", a.Id)
	err := tx.Find(&row).Error
	if err != nil {
		log.Println(err.Error())
	}
	return row
}

func OutArticleAdd() error {
	var (
		row   []OutArticle
		cover string
	)
	err := GetDbCli().Session(&gorm.Session{}).Table("slate_article").Find(&row).Error
	if err != nil {
		log.Println(err)
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("articles")
	for i := 0; i < len(row); i++ {
		if row[i].Cover == "" {
			cover = "/upload/1646706814654.png"
		}
		if VerificationArticle(row[i].Title) != nil {
			timeStr := time.Unix(row[i].Timestamp, 0).Format("2006-01-02 15:04:05")
			err := tx.Create(article{Id: time.Now().UnixMilli(), Title: row[i].Title, Lang: "cn", Cover: cover, Summary: row[i].OverView, Markdown: row[i].Article, RichText: row[i].Articletext, Created: timeStr, Updated: timeStr}).Error
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
	tx = tx.Table("articles").Order("id desc").Where("cate_id in ?", arr).Preload("Category")
	if a.CateId != 0 {
		tx = tx.Where("cate_id = ?", a.CateId)
	}
	var row []CourseBanner
	err := tx.Find(&row).Error
	if err != nil {
		log.Println(err.Error())
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
