package biz

import (
	"help_center/internal/data"
	"log"
	"net/http"
)

//AddBanner 添加Banner
func AddBanner(d *data.Banner, r *http.Request) *BaseJson {
	if d.Chain == 0 {
		return &BaseJson{Code: 0, Data: "该文章Id不能为空"}
	}
	_, err := data.VerificationTitle(d.Chain)
	if err != nil {
		return &BaseJson{Code: 0, Data: "该文章Id不存在"}
	}
	err = data.BannerAdd(d, r)
	if err != nil {
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加Banner"}
	}
}

//DelBanner 删除标签
func DelBanner(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.BannerDel(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功删除Banner"}
	}
}

//ModBanner 修改Banner
func ModBanner(d *data.Banner) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	if d.Chain != 0 {
		_, err := data.VerificationTitle(d.Chain)
		if err != nil {
			return &BaseJson{Code: 0, Data: "该文章Id不存在"}
		}
	}
	row, _ := data.VerificationBanner()
	if d.Number > len(row) {
		return &BaseJson{Code: 0, Data: "Banner序号超出最大范围"}
	}

	err := data.BannerUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改Banner"}
	}
}

//GetBanner 获取Bannner
func GetBanner() *BaseJson {
	list := data.BannerSearch()
	return &BaseJson{Code: 1, Data: list}
}
