package biz

import (
	"help_center/internal/data"
	"log"
)

// GetCurrency 获取代币列表
func GetCurrency(d *data.ArticleQuery) *JsonFormat {
	list := d.CurrencySearch()
	return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, Data: list}
}

// AddCurrency 添加代币信息
func AddCurrency(d *data.Currency) *BaseJson {
	if d.CurrencyName == "" {
		return &BaseJson{Code: 0, Data: "代币名不能为空"}
	}
	err := data.CurrencyAdd(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加代币"}
	}
}

// DelCurrency 删除代币信息
func DelCurrency(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.CurrencyDelete(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功成功删除文章"}
	}
}

// ModCurrency 修改代币信息
func ModCurrency(d *data.Currency) *BaseJson {
	if d.CurrencyName == "" {
		return &BaseJson{Code: 0, Data: "代币名不能为空"}
	}
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.CurrencyUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改文章"}
	}
}
