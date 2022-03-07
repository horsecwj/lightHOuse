package biz

import (
	"help_center/internal/data"
	"log"
)

func AddChain(d *data.Chain) *BaseJson {
	_, err := data.VerificationChain(d.Name)
	if err == nil {
		return &BaseJson{Code: 0, Data: "该链名已存在"}
	}
	err = data.ChainAdd(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加链"}
	}
}

func DelChain(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.ChainDel(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功删除链"}
	}
}

func ModChain(d *data.Chain) *BaseJson {
	chain, err := data.VerificationChain(d.Name)
	if d.Name != chain.Name {
		if err == nil {
			return &BaseJson{Code: 0, Data: "该链名已存在"}
		}
	}
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err = data.ChainUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改标签"}
	}
}

func GetChain() *BaseJson {
	list := data.CategoryChain()
	return &BaseJson{Code: 1, Data: list}
}
