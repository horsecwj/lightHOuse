package data

import (
	"time"

	"gorm.io/gorm"
)

func ClassAdd(d *Class) error {
	t := time.Now()
	if d.Id == 0 {
		d.Id = t.UnixMilli()
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("classes")
	return tx.Create(&d).Error
}

func ClassDel(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("classes")
	return tx.Delete(Class{}, "id = ?", id).Error
}

func ClassUpdate(c *Class) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("classes")
	return tx.Where("id = ?", c.Id).Updates(&c).Error
}
