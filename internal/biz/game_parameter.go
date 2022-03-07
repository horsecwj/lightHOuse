package biz

import (
	"help_center/internal/data"
)

func GetGameParameter(d *data.GameQuery, isAdm bool) *JsonFormat {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 10
	}
	num := d.ParameterCount()
	if num > 0 {
		list := d.SearchGameParameter(isAdm)
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: nil}
}

func GetGameCmk(d *data.GameQuery, losers bool) *JsonFormat {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 10
	}
	num := d.GameCount()
	if num > 0 {
		list := d.SearchGameCmk(losers)
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: nil}
}
