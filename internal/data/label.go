package data

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func LabelAdd(c *Label) error {
	t := time.Now()
	if c.Id == 0 {
		c.Id = t.UnixMilli()
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("labels")
	return tx.Create(&c).Error
}

func LabelDel(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("labels")
	return tx.Delete(Label{}, "id = ?", id).Error
}

func LabelUpdate(c *Label) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("labels")
	return tx.Where("id = ?", c.Id).Updates(&c).Error
}

func SearchLabel() interface{} {
	var list = make([]Label, 0, 20)
	tx := GetDbCli().Session(&gorm.Session{}).Table("labels").Order("id")
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list[1:]
}
