package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	googleidtokenverifier "github.com/movsb/google-idtoken-verifier"
	"gorm.io/gorm"
)

func AddUser(claims *googleidtokenverifier.ClaimSet, from string) error {
	now := time.Now()
	Id := now.UnixMilli()
	code, err := TestInviteCode(Id)
	if err != nil {
		return err
	}
	user := &UserLogin{
		Id:      Id,
		Email:   claims.Email,
		Subject: claims.Sub,
		Code:    code,
		Number:  0,
		From:    from,
	}
	tx := GetDbCli().Session(&gorm.Session{}).Table("user_logins")
	return tx.Create(&user).Error
}

func CreateToken(Id int64) (string, error) {
	now := time.Now()
	expiredAt := now.Add(60 * 60 * 24 * 7 * time.Second)
	claims := AuthClaims{
		ID:        uuid.NewString(),
		IssuedAt:  now.Unix(),
		ExpiresAt: expiredAt.Unix(),
		Issuer:    "lighthouse",
		Subject:   strconv.FormatUint(uint64(Id), 10),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		log.Println(err.Error())
	}
	return token, nil
}

func (ac AuthClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := time.Now().Unix()
	if now > ac.ExpiresAt {
		vErr.Errors |= jwt.ValidationErrorExpired
	}
	if now < ac.IssuedAt {
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}
	if vErr.Errors == 0 {
		return nil
	}
	return vErr
}

func ParseToken(ctx context.Context, ts string) (*AuthClaims, error) {
	claims := &AuthClaims{}
	token, err := jwt.ParseWithClaims(ts, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func UserSearch(Id int64, r *http.Request) (*UserData, error) {
	user := &UserData{}
	err := GetDbCli().Session(&gorm.Session{}).Table("user_logins").Where("id = ?", Id).First(user).Error
	if err != nil {
		return nil, err
	}
	user.Code = fmt.Sprintf("http://%s/introduce?code=%s", r.Host, user.Code)
	return user, nil
}

func UsersSearch(email string) ([]UsersData, error) {
	var users []UsersData
	tx := GetDbCli().Session(&gorm.Session{}).Table("user_logins")
	if email != "" {
		tx = tx.Where("email = ?", email)
	}
	err := tx.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NotesUpdate(d *Notes) error {
	tx := GetDbCli().Session(&gorm.Session{}).Table("user_logins")
	return tx.Where("id = ?", d.Id).Update("notes", d.Notes).Error
}
