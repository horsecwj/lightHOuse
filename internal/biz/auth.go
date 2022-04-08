package biz

import (
	"help_center/internal/conf"
	"help_center/internal/data"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	googleidtokenverifier "github.com/movsb/google-idtoken-verifier"
)

func AdminLogin(d *LoginData) (int, *BaseJson) {
	if d.Name == conf.GetCfg().Admin.User && d.Password == conf.GetCfg().Admin.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = conf.GetCfg().Admin.User
		exp := time.Now().Add(time.Hour * 6).Unix()
		claims["exp"] = exp
		// Generate encoded token and send it as response.
		secret := conf.GetCfg().Jwt.Secret
		tk, err := token.SignedString([]byte(secret)) //密钥
		if err != nil {
			return http.StatusUnauthorized, &BaseJson{Code: 0, Data: err.Error()}
		}
		return http.StatusOK, &BaseJson{Code: 1, Data: tk}
	} else {
		return http.StatusUnauthorized, &BaseJson{Code: 0, Data: "用户名或密码错误"}
	}
}

func ParseGoogleToken(TokenId, code string) *BaseJson {

	clientID := "559756290278-9v1ngbvivap03i80qntgsin48ggmj5pc.apps.googleusercontent.com"
	claims, err := googleidtokenverifier.Verify(TokenId, clientID)
	if err != nil {
		return &BaseJson{Code: 0, Data: "Token error"}
	}
	if _, err := data.VerificationUserLogin(claims.Sub); err != nil {
		err := data.AddUser(claims, code)
		if err != nil {
			return &BaseJson{Code: 0, Data: err}
		}
		if code != "" {
			err := data.UserLoginByCode(code)
			if err != nil {
				log.Println(err)
			}
		}
	}
	row, _ := data.VerificationUserLogin(claims.Sub)

	token, err := data.CreateToken(row.Id)
	if err != nil {
		return &BaseJson{Code: 0, Data: err}
	}
	return &BaseJson{Code: 1, Data: token}
}

func GetUser(Id int64, r *http.Request) *BaseJson {
	user, err := data.UserSearch(Id, r)
	if err != nil {
		return &BaseJson{Code: 0, Data: err}
	}
	return &BaseJson{Code: 1, Data: user}

}

func GetUsers(email string) *BaseJson {
	users, err := data.UsersSearch(email)
	if err != nil {
		return &BaseJson{Code: 0, Data: err}
	}
	return &BaseJson{Code: 1, Data: users}
}

func ModNotes(d *data.Notes) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.NotesUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改备注"}
	}
}
