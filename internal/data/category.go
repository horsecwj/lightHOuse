package data

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func CategoryAdd(c *Category) error {
	t := time.Now()
	if c.Id == 0 {
		c.Id = t.UnixMilli()
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("categories")
	return tx.Create(&c).Error
}

func CategoryDelete(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("categories")
	return tx.Delete(Category{}, "id = ?", id).Error
}

func CategoryUpdate(c *Category) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("categories")
	return tx.Where("id = ? and lang = ?", c.Id, c.Lang).Updates(&c).Error
}

func (c *CategoryQuery) CategorySearch(adm bool) interface{} {
	var list = make([]Category, 0, 20)
	tx := GetDbCli().Session(&gorm.Session{}).Table("categories").Order("id")
	if c.Id != 0 {
		tx = tx.Where("id = ?", c.Id)
	}
	if c.ParentId != 0 {
		tx = tx.Where("parent_id = ?", c.ParentId)
	}
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	if list[0].Id == 1 {
		list[0].Id = 0
	}
	if adm {
		return list[2:]
	} else {
		return list
	}
}

func CategoryCount(pid int64) int64 {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("categories")
	if pid != 0 {
		tx = tx.Where("parent_id = ?", pid)
	}
	err := tx.Count(&count).Error
	if err != nil {
		log.Println(err.Error())
	}
	return count
}
