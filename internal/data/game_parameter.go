package data

import (
	"log"

	"gorm.io/gorm"
)

func GameParameterAdd(GameFi string) interface{} {
	tx := GetDbCli().Session(&gorm.Session{}).Table("top_cko_game_fi").Where("game_fi = ?", GameFi)
	var result = GameParameter{}
	err := tx.First(&result).Error
	if err != nil {
		log.Println(err)
	}
	gp := GetDbCli().Session(&gorm.Session{}).Table("game_parameters")
	return gp.Where("game_fi = ?", GameFi).
		Updates(GameParameter{Coin: result.Coin, Price: result.Price, OneDay: result.OneDay, OneWeek: result.OneWeek, DayVolume: result.DayVolume, MktCap: result.MktCap}).Error
}

func (a *GameQuery) SearchGameParameter(isAdm bool) interface{} {
	var list = make([]GameParameter, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("game_parameters").Order("mkt_cap desc").Where("status = ?", 1)
	if !isAdm {
		if a.ClassId != 0 {
			row := &GameClass{}
			class := GetDbCli().Session(&gorm.Session{}).Table("game_class").Where("class_id = ?", a.ClassId)
			err := class.Find(&row).Error
			if err != nil {
				log.Println(err.Error())
			}
			ty := GetDbCli().Session(&gorm.Session{}).Table("games").Where("id = ?", row.GameId)
			errr := ty.Find(&list).Error
			if errr != nil {
				log.Println(err.Error())
			}
		}
		return list
	}
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list
}

func (a *GameQuery) SearchGameCmk(losers bool) interface{} {
	var list = make([]Cmk, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("top_cmk_game_fi")
	if !losers {
		tx.Order("id")
		err := tx.Find(&list).Error
		if err != nil {
			log.Println(err.Error())
		}
		return list
	}
	tx.Order("id desc")
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list
}
