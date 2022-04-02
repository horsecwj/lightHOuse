package data

import (
	"errors"
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
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner")
	var row []Banner
	err := tx.Find(&row).Error
	if err != nil {
		log.Println(err.Error())
	}

	c.Number = len(row) + 1

	if c.Cover == "" {
		c.Cover = "/upload/1645606276524.png"
	}
	return tx.Create(&c).Error
}

func BannerDel(id int64) error {
	var data []Banner
	var row Banner
	err := GetDbCli().Session(&gorm.Session{}).Table("banner").Where("id = ?", id).First(&row).Error
	if err != nil {
		log.Println(err)
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner")
	err = tx.Where("id = ?", id).Delete(Banner{}).Error
	if err != nil {
		log.Println(err)
	}
	err = GetDbCli().Session(&gorm.Session{}).Table("banner").Find(&data).Order("number").Error
	if err != nil {
		log.Println(err)
	}
	for i := row.Number - 1; i < len(data); i++ {
		err := GetDbCli().Session(&gorm.Session{}).Table("banner").Where("id = ?", data[i].Id).Update("number", i+1).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
	return nil
}

func BannerUpdate(c *Banner) error {
	var (
		data []Banner
		row  Banner
		err  error
	)
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner")
	if c.Number > 0 && c.Number < 6 {
		err = tx.Where("id = ?", c.Id).Updates(&c).Error
		if err != nil {
			log.Println(err.Error())
		}
		err = tx.Where("id = ?", c.Id).First(&row).Error
		if err != nil {
			log.Println(err.Error())
		}
		if row.Number != c.Number {
			err = GetDbCli().Session(&gorm.Session{}).Table("banner").Order("number").Find(&data).Error
			if err != nil {
				log.Println(err)
			}
			for i := 1; i <= len(data); i++ {
				if data[i-1].Id != c.Id {
					var tmpData = data[i-1]
					var num = data[i-1].Number + 1
					if num == 5 {
						tmpData.Number = num
					} else {
						tmpData.Number = (num) % 5
					}
					err = tx.Where("id = ?", c.Id).Updates(&tmpData).Error
					if err != nil {
						log.Println(err.Error())
					}
				}
			}
		}
	} else {
		err = errors.New("number is zero or more than 5")
	}
	return err
}

func BannerSearch() interface{} {
	type banner struct {
		Id       int64  `json:"id"`
		CadeId   int64  `json:"cade_id"`
		ParentId int64  `json:"parent_id"`
		Cover    string `json:"cover"`
		Chain    int64  `json:"chain"`
		Title    string `json:"title"`
		Number   int    `json:"number"`
		Rows     int64  `json:"rows"`
		Cols     int64  `json:"cols"`
	}
	var list = make([]banner, 0, 20)
	tx := GetDbCli().Session(&gorm.Session{}).Table("banner").Order("number")
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	for i := 0; i < len(list); i++ {
		if list[i].Chain != 0 {
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
	}
	if len(list) > 0 {
		list[0].Rows = 2
		list[0].Cols = 2
	}
	return list
}
