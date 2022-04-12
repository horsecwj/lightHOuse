package data

import (
	"net"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func VerificationLabel(label string) (*Label, error) {
	row := &Label{}
	err := db.Where("word = ?", label).First(&row).Error
	return row, err
}

func VerificationChain(chain string) (*Chain, error) {
	row := &Chain{}
	err := db.Where("name = ?", chain).First(&row).Error
	return row, err
}

func VerificationClass(class string) (*Class, error) {
	row := &Class{}
	err := db.Where("class = ?", class).First(&row).Error
	return row, err
}

func VerificationTitle(id int64) (*Article, error) {
	row := &Article{}
	err := db.Where("id = ?", id).First(&row).Error
	return row, err
}

func VerificationArticle(Id int64) error {
	row := &article{}
	err := db.Where("Id = ?", Id).First(&row).Error
	return err
}

func VerificationArticleHot(id int64) (*article, error) {
	row := &article{}
	err := db.Where("id = ?", id).First(&row).Error
	return row, err
}

func VerificationHot(hot int64) error {
	row := &article{}
	err := db.Where("hot = ?", hot).First(&row).Error
	return err
}

func VerificationGames(name string) error {
	row := &Game{}
	err := db.Where("game_name = ?", name).First(&row).Error
	return err
}

func VerificationGame(id int64) (string, error) {
	row := &Game{}
	err := db.Where("id = ?", id).First(&row).Error
	return row.GameName, err
}

func VerificationGameParameters(name string) error {
	row := &GameParameter{}
	err := db.Where("game_fi = ?", name).First(&row).Error
	return err
}

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func VerificationBanner() ([]Banner, error) {
	var data []Banner
	err := GetDbCli().Session(&gorm.Session{}).Table("banner").Order("number").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func VerificationUserLogin(Sub string) (*UserLogin, error) {
	row := &UserLogin{}
	err := db.Where("subject = ?", Sub).First(&row).Update("login_time", time.Now()).Error
	return row, err
}

func VerificationUserLoginByCode(Code string) error {
	row := &UserLogin{}
	err := db.Where("code = ?", Code).First(&row).Error
	return err
}

func UserLoginByCode(Code string) error {
	row := &UserLogin{}
	tx := db.Where("code = ?", Code).First(&row)
	if err := tx.Error; err != nil {
		return err
	}
	Number := row.Number + 1
	if err := tx.Update("number", Number).Error; err != nil {
		return err
	}
	return nil
}
