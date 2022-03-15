package data

import (
	"log"
	"strconv"
	"time"

	"github.com/yangtizi/cz88"
	"gorm.io/gorm"
)

type IpRecord struct {
	Id      int64     `json:"id"`
	Ip      string    `json:"ip"`
	Country string    `json:"country"`
	Created time.Time `json:"created"`
}

type Data struct {
	User    int
	NewUser int
	Country []Region
}

type Region struct {
	Country string `json:"country"`
	Num     int64
}

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
	if d.Day == 0 {
		d.Day = 1
	}
	var row Data
	row.User = UserNum(0)
	newuser := UserNum(d.Day)
	row.NewUser = newuser
	row.Country = CountryData(d.Day)
	return row
}

func CountryData(day int) []Region {
	var (
		row     []IpRecord
		country []Region
		middle  Region
		count   int64
		com     []string
	)
	tx := GetDbCli().Session(&gorm.Session{}).Table("ip_records")
	if day != 0 {
		t1 := time.Now().Year()
		t2 := time.Now().Month()
		t3 := time.Now().Day()
		EndTime := time.Date(t1, t2, t3+1, 0, 0, 0, 0, time.Local)
		StartTime := time.Date(t1, t2, t3-day+1, 0, 0, 0, 0, time.Local)
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
			err := tx.Where("country = ?", row[i].Country).Count(&count).Error
			if err != nil {
				log.Println(err.Error())
			}
		}
		//fmt.Println(row[i].Country, count)
	}
	// for x := range com {
	// 	err := tx.Where("country = ?", com[x]).Count(&count).Error
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}

	// country[x].Country = com[x]
	// country[x].Num = count
	//fmt.Println(com[x], count)
	middle.Country = ""
	middle.Num = 0
	// if country[x].Num > country[x-1].Num {
	// 	middle = country[x]
	// 	country[x-1] = country[x]
	// 	country[x] = middle
	// }
	//}
	return country
}

func UserNum(day int) int {
	var count int64
	tx := GetDbCli().Session(&gorm.Session{}).Table("ip_records")
	if day != 0 {
		t1 := time.Now().Year()
		t2 := time.Now().Month()
		t3 := time.Now().Day()
		EndTime := time.Date(t1, t2, t3+1, 0, 0, 0, 0, time.Local)
		StartTime := time.Date(t1, t2, t3-day+1, 0, 0, 0, 0, time.Local)
		tx = tx.Where("created BETWEEN ? AND ?", StartTime, EndTime)
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
