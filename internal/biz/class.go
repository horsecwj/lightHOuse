package biz

import (
	"help_center/internal/data"
	"log"

	"gorm.io/gorm"
)

//添加游戏类型
func AddClass(d *data.Class) *BaseJson {
	_, err := data.VerificationClass(d.Class)
	if err == nil {
		return &BaseJson{Code: 0, Data: "该类型已存在"}
	}
	err = data.ClassAdd(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加类型"}
	}
}

//获取游戏类型
func GetClass() *BaseJson {
	var list = make([]data.Class, 0, 20)
	tx := data.GetDbCli().Session(&gorm.Session{}).Table("classes").Order("id")
	err := tx.Find(&list).Error
	if err != nil {
		log.Println(err.Error())
	}

	return &BaseJson{Code: 1, Data: list}
}

//删除游戏类型
func DelClass(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.ClassDel(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功删除标签"}
	}
}
