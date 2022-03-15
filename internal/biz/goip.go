package biz

import (
	"help_center/internal/data"
	"net/http"
)

func Goip(req *http.Request) {

	ip := req.RemoteAddr
	ip = ip[:len(ip)-6]
	err := data.VerificationIp(ip)
	if err != nil {
		data.Addip(ip)
	}
}

func GetData(d *data.Day) *BaseJson {
	list := d.DataSearch()
	return &BaseJson{Code: 1, Data: list}

}
