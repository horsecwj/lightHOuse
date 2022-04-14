package biz

import (
	"help_center/internal/data"
	"net"
	"net/http"
)

func RemoteIp(req *http.Request) {
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
	err := data.VerificationIp(remoteAddr)
	if err != nil {
		data.Addip(remoteAddr)
	}
}

func GetData(d *data.Day) *BaseJson {
	list := d.DataSearch()
	return &BaseJson{Code: 1, Data: list}
}
