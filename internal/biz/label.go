package biz

import (
	"help_center/internal/data"
	"log"
)

//AddLabel 添加标签
func AddLabel(d *data.Label) *BaseJson {
	if d.Word == "" {
		return &BaseJson{Code: 0, Data: "标签不能为空"}
	}
	_, err := data.VerificationLabel(d.Word)
	if err == nil {
		return &BaseJson{Code: 1, Data: "该标签已经存在"}
	}
	err = data.LabelAdd(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加标签"}
	}
}

//DelLabel 删除标签
func DelLabel(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.LabelDel(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功删除标签"}
	}
}

//ModLabel 修改标签
func ModLabel(d *data.Label) *BaseJson {
	_, err := data.VerificationLabel(d.Word)
	if err == nil {
		return &BaseJson{Code: 1, Data: "该标签已经存在"}
	}
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err = data.LabelUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改标签"}
	}
}

//GetLabel 获取标签
func GetLabel() *BaseJson {
	list := data.SearchLabel()
	return &BaseJson{Code: 1, Data: list}
}
