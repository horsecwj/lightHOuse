package data

import (
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func BannerAdd(c *Banner, r *http.Request) error {
	t := time.Now()
	if c.Id == 0 {
		c.Id = t.UnixMilli()
	}
	if c.Cover == "" {
		c.Cover = "/upload/1645606276524.png"
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner")
	return tx.Create(&c).Error
}

func BannerDel(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner")
	return tx.Delete(Label{}, "id = ?", id).Error
}

func BannerUpdate(c *Banner) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner")
	return tx.Where("id = ?", c.Id).Updates(&c).Error
}

func BannerSearch() interface{} {
	type banner struct {
		Id       int64  `json:"id"`
		CadeId   int64  `json:"cade_id"`
		ParentId int64  `json:"parent_id"`
		Cover    string `json:"cover"`
		Chain    int64  `json:"chain"`
		Title    string `json:"title"`
		Rows     int64  `json:"rows"`
		Cols     int64  `json:"cols"`
	}
	var list = make([]banner, 0, 20)
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner").Order("id")
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	for i := 0; i < len(list); i++ {
		title, err := VerificationTitle(list[i].Chain)
		if err != nil {
			log.Println(err.Error())
		}
		list[i].Title = title.Title
		list[i].CadeId = title.CateId
		list[i].Rows = 1
		list[i].Cols = 1
		var category []Category
		err = GetDbCli().Session(&gorm.Session{}).Table("categories").Where("id = ?", title.CateId).Find(&category).Error
		if err != nil {
			log.Println(err.Error())
		}
		if len(category) > 0 {
			list[i].ParentId = category[0].ParentId
		}
	}
	list[0].Rows = 2
	list[0].Cols = 2
	return list
}
