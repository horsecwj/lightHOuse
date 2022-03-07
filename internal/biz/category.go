package biz

import (
	"help_center/internal/data"
	"log"
	"strings"
)

// GetCategory 获取分类
func GetCategory(d *data.CategoryQuery) *BaseJson {
	list := d.CategorySearch()
	return &BaseJson{Code: 1, Data: list}
}

// AddCategory 添加分类
func AddCategory(d *data.Category) *BaseJson {
	err := data.CategoryAdd(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加分类"}
	}
}

// DelCategory 删除分类
func DelCategory(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	if data.CategoryCount(d.Id) > 0 {
		return &BaseJson{Code: 0, Data: "存在子分类，请删除子分类或修改子分类的父级分类"}
	}
	err := data.CategoryDelete(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功删除分类"}
	}
}

// ModCategory 修改分类
func ModCategory(d *data.Category) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	if !strings.Contains("cn,en", d.Lang) {
		return &BaseJson{Code: 0, Data: "参数 lang 值应为cn或en"}
	}
	err := data.CategoryUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改分类"}
	}
}
