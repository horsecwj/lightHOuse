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
	// if !isAdm {
	// 	row := &GameClass{}
	// 	class := GetDbCli().Session(&gorm.Session{}).Table("game_class").Where("class_id = ?", a.ClassId)
	// 	err := class.Find(&row).Error
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	}
	// 	ty := GetDbCli().Session(&gorm.Session{}).Table("games").Where("id = ?", row.GameId)
	// 	errr := ty.Find(&list).Error
	// 	if errr != nil {
	// 		log.Println(err.Error())
	// 	}
	// 	return list
	// }
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
