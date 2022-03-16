package data

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func GameParameterAdd() error {
	var row []GameParameter
	err := GetDbCli().Session(&gorm.Session{}).Table("top_cko_game_fi").Find(&row).Error
	if err != nil {
		log.Println(err)
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("game_parameters")
	for i := 0; i < len(row); i++ {
		if VerificationGameParameters(row[i].GameFi) != nil {
			err := tx.Create(GameParameter{Id: time.Now().UnixMilli(), Coin: row[i].Coin, GameFi: row[i].GameFi, Price: row[i].Price, OneDay: row[i].OneDay, OneWeek: row[i].OneWeek, DayVolume: row[i].DayVolume, MktCap: row[i].MktCap}).Error
			if err != nil {
				log.Println(err)
			}
		}
	}
	return err
}

func (a *GameQuery) SearchGameParameter(isAdm bool) interface{} {
	var list = make([]GameParameter, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("game_parameters").Order("id")
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list
}

func GameParameterDelete(name string) error {
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("game_parameters").Delete(GameParameter{}, "game_fi = ?", name).Error
}

func (a *GameQuery) SearchGameCmk(losers bool) interface{} {
	var list = make([]Cmk, 0, a.PageSize)
	if !losers {
		tx := GetDbCli().Session(&gorm.Session{}).Table("top_cmk_game_fi")
		tx.Order("id")
		err := tx.Find(&list).Error
		if err != nil {
			log.Println(err.Error())
		}
		return list
	} else {
		tx := GetDbCli().Session(&gorm.Session{}).Table("top_cmk_game_fi_losers")
		tx.Order("id")
		err := tx.Find(&list).Error
		if err != nil {
			log.Println(err.Error())
		}
		for i := range list {
			list[i].OneDay = "-" + list[i].OneDay
		}
		return list
	}
}
