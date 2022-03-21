package data

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/yangtizi/cz88"
	"gorm.io/gorm"
)

var China = []string{
	"河北", "山西", "辽宁", "吉林", "黑龙江", "江苏", "浙江", "安徽", "福建", "江西", "山东", "河南", "湖北", "湖南", "广东", "海南", "四川", "贵州", "云南", "陕西", "甘肃", "青海", "内蒙古", "宁夏", "广西", "西藏", "新疆", "北京", "上海", "天津", "重庆",
}

//添加ip及地区
func Addip(ip string) error {
	country := cz88.GetAddressShort(ip)
	t := time.Now()
	var row IpRecord
	row.Id = t.UnixMilli()
	row.Ip = ip
	for i := 0; i < len(China); i++ {
		if country == China[i] {
			row.Country = "中国"
		}
	}
	if country == "香港" {
		row.Country = "中国香港"
	}
	if country == "澳门" {
		row.Country = "中国澳门"
	}
	if country == "台湾" {
		row.Country = "中国台湾"
	}
	row.Created = t
	tx := GetDbCli().Session(&gorm.Session{})
	return tx.Table("ip_records").Create(&row).Error
}

func (d *Day) DataSearch() interface{} {
	var row Data
	row.User = UserNum(nil, false)
	newuser := UserNum(d, true)
	row.NewUser = newuser
	row.Country = CountryData(d)
	return row
}

func CountryData(d *Day) []Region {
	var (
		row     []IpRecord
		country Region
		Country []Region
		middle  Region
		count   int64
		com     []string
	)
	tx := GetDbCli().Session(&gorm.Session{}).Table("ip_records")
	if d.StartTime != 0 && d.EndTime != 0 {
		StartTime := time.Unix(d.StartTime, 0)
		EndTime := time.Unix(d.EndTime, 0)
		tx = tx.Where("created BETWEEN ? AND ?", StartTime, EndTime)
	}
	err := tx.Find(&row).Error
	if err != nil {
		log.Println(err.Error())
	}
	for i := range row {
		flag := true
		for j := range com {
			if row[i].Country == com[j] {
				flag = false
			}
		}
		if flag {
			com = append(com, row[i].Country)
		}
	}
	for x := range com {
		ty := GetDbCli().Session(&gorm.Session{}).Table("ip_records")
		if d.StartTime != 0 && d.EndTime != 0 {
			StartTime := time.Unix(d.StartTime, 0)
			EndTime := time.Unix(d.EndTime, 0)
			ty = ty.Where("created BETWEEN ? AND ?", StartTime, EndTime)
		}
		err := ty.Where("country = ?", com[x]).Count(&count).Error
		if err != nil {
			fmt.Println(err.Error())
		}
		country.Country = com[x]
		country.Num = count
		Country = append(Country, country)
	}
	for y := 0; y < len(Country)-1; y++ {
		if Country[y+1].Num > Country[y].Num {
			middle = Country[y]
			Country[y] = Country[y+1]
			Country[y+1] = middle
		}
	}
	return Country
}

func UserNum(d *Day, new bool) int {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("ip_records")
	if new {
		if d.StartTime != 0 && d.EndTime != 0 {
			StartTime := time.Unix(d.StartTime, 0)
			EndTime := time.Unix(d.EndTime, 0)
			tx = tx.Where("created BETWEEN ? AND ?", StartTime, EndTime)
		}
	}
	err := tx.Count(&count).Error
	if err != nil {
		log.Println(err.Error())
	}
	strCount := strconv.FormatInt(count, 10)
	intCount, err := strconv.Atoi(strCount)
	if err != nil {
		log.Println(err.Error())
	}
	return intCount
}

//验证ip是否存在数据库
func VerificationIp(ip string) error {
	row := &IpRecord{}
	err := db.Where("ip = ?", ip).First(&row).Error
	return err
}
