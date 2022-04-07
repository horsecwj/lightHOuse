package data

import (
	"errors"
	"math/rand"
	"time"
)

type code struct {
	base    string // 进制的包含字符, string类型
	decimal uint64 // 进制长度
	pad     string // 补位字符,若生成的code小于最小长度,则补位+随机字符, 补位字符不能在进制字符中
	len     int    // code最小长度
}

func TestInviteCode(Id int64) (string, error) {
	inviteCode := code{
		base:    "HVE8S2DZX9C7P5IK3MJUAR4WYLTN6BGQ",
		decimal: 32,
		pad:     "F",
		len:     6,
	}
	// 初始化检查
	if res, err := inviteCode.initCheck(); !res {
		return "", err
	}
	code := inviteCode.idToCode(uint64(Id))
	return code, nil
}

// id转code
func (c *code) idToCode(id uint64) string {
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % c.decimal
		id = id / c.decimal
		res += string(c.base[mod])
	}
	resLen := len(res)
	if resLen < c.len {
		res += c.pad
		for i := 0; i < c.len-resLen-1; i++ {
			rand.Seed(time.Now().UnixNano())
			res += string(c.base[rand.Intn(int(c.decimal))])
		}
	}
	return res
}

func (c *code) initCheck() (bool, error) {
	lenBase := len(c.base)
	// 检查进制字符
	if c.base == "" {
		return false, errors.New("base string is nil or empty")
	}
	// 检查长度是否符合
	if uint64(lenBase) != c.decimal {
		return false, errors.New("base length and len not match")
	}
	return true, errors.New("")
}
