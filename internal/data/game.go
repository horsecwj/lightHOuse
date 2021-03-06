package data

import (
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func GameAdd(a *Game) error {
	t := time.Now()
	if a.Id == 0 {
		a.Id = t.UnixMilli()
	}
	tx := GetDbCli().Session(&gorm.Session{})
	if VerificationGameParameters(a.GameName) != nil {
		tx.Table("game_parameters").Create(&GameParameter{Id: a.Id, GameFi: a.GameName})
	}
	return tx.Table("games").Create(&a).Error
}

func GameDelete(id int64) error {
	tx := GetDbCli().Session(&gorm.Session{})
	err := tx.Table("game_chain").Delete(GameChain{}, "game_id = ?", id).Error
	if err != nil {
		log.Println(err.Error())
	}
	err = tx.Table("game_label").Delete(GameLabel{}, "game_id = ?", id).Error
	if err != nil {
		log.Println(err.Error())
	}
	err = tx.Table("game_class").Delete(GameClass{}, "game_id = ?", id).Error
	if err != nil {
		log.Println(err.Error())
	}
	err = tx.Table("game_currency").Delete(GameCurrency{}, "game_id = ?", id).Error
	if err != nil {
		log.Println(err.Error())
	}
	gamename, err := VerificationGame(id)
	if err != nil {
		log.Println(err.Error())
	}
	err = GameParameterDelete(gamename)
	if err != nil {
		log.Println(err)
	}
	return tx.Table("games").Delete(Game{}, "id = ?", id).Error
}

func GameUpdate(a *Game) error {
	tx := GetDbCli().Session(&gorm.Session{})
	if a.Chain != nil {
		tx.Table("game_chain").Delete(GameChain{}, "game_id = ?", a.Id)
	}
	if a.Label != nil {
		tx.Table("game_label").Delete(GameLabel{}, "game_id = ?", a.Id)
	}
	if a.Class != nil {
		tx.Table("game_class").Delete(GameClass{}, "game_id = ?", a.Id)
	}
	if a.Currency != nil {
		tx.Table("game_currency").Delete(GameCurrency{}, "game_id = ?", a.Id)
	}
	if a.Status == 0 {
		tx.Table("games").Where("id = ?", a.Id).Update("status", 0)
	}
	if a.Release == 0 {
		tx.Table("games").Where("id = ?", a.Id).Update("release", 0)
	}
	return tx.Table("games").Where("id = ?", a.Id).Updates(&a).Error
}

func (a *GameQuery) GameSearch(adm bool) interface{} {
	var list = make([]Game, 0, a.PageSize)
	tx := GetDbCli().Session(&gorm.Session{}).Table("games").Order("created desc").
		Preload("Label").Preload("Chain").Preload("Class").Preload("Currency").Preload("New")
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)
	}
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	if a.GameFi != "" {
		tx = tx.Where("game_name = ?", a.GameFi)
	}
	if a.Status != 0 {
		tx = tx.Where("status = ?", a.Status)
	}
	if a.ClassId != 0 {
		tx = tx.Joins("left join game_class on games.id = game_class.game_id").Where("game_class.class_id = ?", a.ClassId)
	}
	if a.ChainId != 0 {
		tx = tx.Joins("left join game_chain on games.id = game_chain.game_id").Where("game_chain.chain_id = ?", a.ChainId)
	}
	if !adm {

		type game struct {
			Id              int64      `json:"id"`
			GameName        string     `json:"title"`
			Cover           string     `json:"cover"`
			Summary         string     `json:"summary"`
			Lang            string     `json:"lang"`
			Currency        []currency `json:"currency" gorm:"many2many:game_currency"`
			Chain           []Chain    `json:"chain" gorm:"many2many:game_chain"`
			Label           []Label    `json:"label" gorm:"many2many:game_label"`
			Class           []Class    `json:"class" gorm:"many2many:game_class"`
			Telegram        string     `json:"telegram"`
			Facebook        string     `json:"facebook"`
			Twitter         string     `json:"twitter"`
			Youtube         string     `json:"youtube"`
			GameUrl         string     `json:"game_url"`
			Guide           string     `json:"guide"`
			AboutGames      string     `json:"about_games"`
			New             []article  `json:"new" gorm:"many2many:game_article"`
			Stragegy        string     `json:"stragegy"`
			RevenueAnalysis string     `json:"revenue_analysis"`
			Updated         string     `json:"updated"`
		}
		var result = make([]game, 0, a.PageSize)
		err := tx.Find(&result).Error
		if err != nil {
			log.Println(err.Error())
		}
		for i := range result {
			for x := range result[i].Currency {
				result[i].Currency[x].IssueAt = result[i].Currency[x].IssueAt[:10]
				result[i].Updated = result[i].Updated[:10]
			}
		}
		return result
	} else {
		err := tx.Find(&list).Error
		if err != nil {
			log.Println(err.Error())
		}
		return list
	}
}

func GameMatch(subStr string) (interface{}, int) {
	tx := GetDbCli().Session(&gorm.Session{}).Table("games").Order("id")
	tx = tx.Where("game_name like ? ", "%"+subStr+"%").Preload("Class").Preload("Currency").Preload("Chain")
	tx = tx.Limit(10)
	type game struct {
		Id       int64      `json:"id"`
		GameName string     `json:"game_name"`
		Currency []Currency `json:"currency" gorm:"many2many:game_currency"`
		Chain    []Chain    `json:"chain" gorm:"many2many:game_chain"`
		Class    []Class    `json:"class" gorm:"many2many:game_class"`
		Status   int        `json:"status"`
		Release  string     `json:"release"`
	}
	var result []game
	err := tx.Find(&result).Error
	if err != nil {
		log.Println(err.Error())
	}
	return result, len(result)

}

func (a *GameQuery) GameCount() int {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("games")
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)
	}
	if a.Status != 0 {
		tx = tx.Where("status = ?", a.Status)
	}
	if a.ClassId != 0 {
		tx = tx.Joins("left join game_class on games.id = game_class.game_id").Where("game_class.class_id = ?", a.ClassId)
	}
	if a.ChainId != 0 {
		tx = tx.Joins("left join game_chain on games.id = game_chain.game_id").Where("game_chain.chain_id = ?", a.ChainId)
	}
	err := tx.Count(&count).Error
	if err != nil {
		log.Println(err.Error())
	}
	strCount := strconv.FormatInt(count, 10)
	intCount, err := strconv.Atoi(strCount)
	if err != nil {
		log.Println(err.Error())
	}
	return intCount
}

func (a *GameQuery) ParameterCount() int {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("game_parameters")
	if a.Id != 0 {
		tx = tx.Where("id = ?", a.Id)
	}
	var row = make([]game, 0)
	ty := GetDbCli().Session(&gorm.Session{}).Table("games").Preload("Class").Preload("Chain")
	if a.Status != 0 {
		ty = ty.Where("status = ?", a.Status).Find(&row)
		var name []string
		for i := 0; i < len(row); i++ {
			name = append(name, row[i].GameName)
		}
		tx = tx.Where("game_fi in ?", name)
	}
	if a.ClassId != 0 {
		ty.Joins("left join game_class on games.id = game_class.game_id").Where("game_class.class_id = ?", a.ClassId).Find(&row)
		var name []string
		for i := 0; i < len(row); i++ {
			name = append(name, row[i].GameName)
		}
		tx = tx.Where("game_fi in ?", name)
	}
	if a.ChainId != 0 {
		ty.Joins("left join game_chain on games.id = game_chain.game_id").Where("game_chain.chain_id = ?", a.ChainId).Find(&row)
		var name []string
		for i := 0; i < len(row); i++ {
			name = append(name, row[i].GameName)
		}
		tx = tx.Where("game_fi in ?", name)
	}
	err := tx.Count(&count).Error
	if err != nil {
		log.Println(err.Error())
	}
	strCount := strconv.FormatInt(count, 10)
	intCount, err := strconv.Atoi(strCount)
	if err != nil {
		log.Println(err.Error())
	}
	return intCount
}

func (a *GameQuery) LikeGame() interface{} {
	var list = make([]Label, 0, 20)
	tx := GetDbCli().Session(&gorm.Session{}).Table("labels").Order("id").Preload("Game").Where("id = ?", a.LabelId)
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	if a.Id != 0 {
		tx = tx.Joins("left join game_label on labels.id = game_label.label_id").Not("game_label.label_id = ?", a.Id)
	}
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}
	return list
}

func (a *GameQuery) GameValue() interface{} {
	tx := GetDbCli().Session(&gorm.Session{}).Table("game_parameters").Preload("Game")
	if a.Page > 0 && a.PageSize > 0 {
		tx = tx.Limit(a.PageSize).Offset((a.Page - 1) * a.PageSize)
	}
	var row = make([]game, 0)
	ty := GetDbCli().Session(&gorm.Session{}).Table("games").Preload("Class").Preload("Chain")
	if a.Status != 0 {
		ty = ty.Where("status = ?", a.Status).Find(&row)
		var name []string
		for i := 0; i < len(row); i++ {
			name = append(name, row[i].GameName)
		}
		tx = tx.Where("game_fi in ?", name)
	}
	if a.ClassId != 0 {
		ty.Joins("left join game_class on games.id = game_class.game_id").Where("game_class.class_id = ?", a.ClassId).Find(&row)
		var name []string
		for i := 0; i < len(row); i++ {
			name = append(name, row[i].GameName)
		}
		tx = tx.Where("game_fi in ?", name)
	}
	if a.ChainId != 0 {
		ty = ty.Joins("left join game_chain on games.id = game_chain.game_id").Where("game_chain.chain_id = ?", a.ChainId).Find(&row)
		var name []string
		for i := 0; i < len(row); i++ {
			name = append(name, row[i].GameName)
		}
		tx = tx.Where("game_fi in ?", name)
	}
	err := ty.Find(&row).Error
	if err != nil {
		log.Println(err.Error())
	}
	type gamevalue struct {
		Id        int64  `json:"id"`
		Coin      string `json:"coin"`
		GameFi    string `json:"game_fi"`
		Price     string `json:"price"`
		OneDay    string `json:"one_day"`
		OneWeek   string `json:"one_week"`
		DayVolume string `json:"day_volume"`
		MktCap    string `json:"mkt_cap"`
		Game      game   `json:"game" gorm:"foreignkey:game_name;references:game_fi"`
	}
	var result = make([]gamevalue, 0, a.PageSize)
	err = tx.Find(&result).Error
	if err != nil {
		log.Println(err.Error())
	}
	for x := 0; x < len(row); x++ {
		for y := 0; y < len(result); y++ {
			if result[y].GameFi == row[x].GameName {
				result[y].Game.Chain = row[x].Chain
				result[y].Game.Class = row[x].Class

			}
		}
	}
	return result
}
