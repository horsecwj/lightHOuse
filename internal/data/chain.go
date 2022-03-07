package data

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func ChainAdd(d *Chain) error {
	t := time.Now()
	if d.Id == 0 {
		d.Id = t.UnixMilli()
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("chains")
	return tx.Create(&d).Error
}

func ChainDel(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("chains")
	return tx.Delete(Class{}, "id = ?", id).Error
}

func ChainUpdate(c *Chain) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("chains")
	return tx.Where("id = ?", c.Id).Updates(&c).Error
}

func CategoryChain() interface{} {
	var list = make([]Chain, 0, 20)
	tx := GetDbCli().Session(&gorm.Session{}).Table("chains").Order("id")
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list
}
