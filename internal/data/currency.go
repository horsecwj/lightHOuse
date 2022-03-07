package data

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func (a *ArticleQuery) CurrencySearch() interface{} {
	var list = make([]Currency, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("currencies").Order("id")
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)

	}
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list
}

func CurrencyAdd(a *Currency) error {
	t := time.Now()
	if a.Id == 0 {
		a.Id = t.UnixMilli()
	}
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("currencies").Create(&a).Error
}

func CurrencyDelete(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("currencies").Delete(Currency{}, "id = ?", id).Error
}

func CurrencyUpdate(a *Currency) error {
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("currencies").Where("id = ?", a.Id).Omit("created").Updates(&a).Error
}
