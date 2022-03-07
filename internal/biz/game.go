package biz

import (
	"help_center/internal/data"
	"log"
)

// GetGame 获取游戏列表
func GetGame(d *data.GameQuery, isAdm bool) *JsonFormat {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 10
	}
	//num := 1
	num := d.GameCount()
	if num > 0 {
		list := d.GameSearch(isAdm)
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: nil}
}

// MatchGame 获取游戏列表
func MatchGame(subStr string) *JsonFormat {
	if subStr != "" {
		list, lenList := data.GameMatch(subStr)
		return &JsonFormat{Code: 1, Page: 1, PageSize: lenList, PageNum: 1, ArticleNum: lenList, Data: list}
	}
	return &JsonFormat{Code: 0, Page: 0, PageSize: 0, PageNum: 0, ArticleNum: 0, Data: "当前参数获取到的游戏数量为0"}
}

// AddGame 添加游戏信息
func AddGame(d *data.Game) *BaseJson {
	if data.VerificationGames(d.GameName) == nil {
		return &BaseJson{Code: 0, Data: "该游戏已添加"}
	}
	err := data.GameAdd(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加游戏"}
	}
}

// DelArticle 删除游戏信息
func DelGame(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.GameDelete(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功成功删除游戏"}
	}
}

// ModGame 修改游戏信息
func ModGame(d *data.Game) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.GameUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改游戏"}
	}
}

//GetLikeGame 获取相关游戏
func GetLikeGame(d *data.GameQuery) *BaseJson {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 4
	}
	num := d.GameCount()
	if num > 0 {
		list := d.LikeGame()
		return &BaseJson{Code: 1, Data: list}
	}
	return &BaseJson{Code: 0, Data: nil}
}

//GetGameValue 获取游戏价值
func GetGameValue(d *data.GameQuery) *JsonFormat {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 10
	}
	num := d.GameCount()
	if num > 0 {
		list := d.GameValue()
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: nil}
}
