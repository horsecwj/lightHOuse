package util

import (
	"github.com/spf13/cast"
	"strings"
	"time"
)

var InternalSortMap = make(map[string]string)

func SegInternalMap() {
	InternalSortMap["Jan"] = "01"
	InternalSortMap["Feb"] = "02"
	InternalSortMap["Mar"] = "03"
	InternalSortMap["Apr"] = "04"
	InternalSortMap["May"] = "05"
	InternalSortMap["Jun"] = "06"
	InternalSortMap["Jul"] = "07"
	InternalSortMap["Aug"] = "08"
	InternalSortMap["Sep"] = "09"
	InternalSortMap["Oct"] = "10"
	InternalSortMap["Nov"] = "11"
	InternalSortMap["Dec"] = "12"
	return
}

func timeParse(timeStr string) (int64, error) {
	//
	temp := strings.Split(timeStr, ".")
	month := temp[0]
	SegInternalMap()
	monthInt := InternalSortMap[month]
	temp2 := strings.Split(temp[1], ",")
	day := temp2[0]
	day = strings.TrimSpace(day)
	temp3 := strings.Split(temp2[1], "at")
	year := temp3[0]
	year = strings.TrimSpace(year)
	temp4 := strings.Split(temp3[1], " ")
	hour := temp4[1]
	hourMm := temp4[2]
	var timeHour string
	if hourMm == "pm" {
		hourTmp := strings.Split(hour, ":")
		timeHour = cast.ToString(cast.ToInt(hourTmp[0])+11) + ":" + hourTmp[1] + ":00"
	} else {
		hourTmp := strings.Split(hour, ":")
		timeHour = cast.ToString(cast.ToInt(hourTmp[0])) + ":" + hourTmp[1] + ":00"
	}
	if len(day) == 1 {
		day = "0" + day
	}
	timeDay := year + "-" + monthInt + "-" + day
	timeStrFin := timeDay + " " + timeHour
	the_time, err := time.Parse("2006-01-02 15:04:05", timeStrFin)
	if err != nil {
		return 0, err
	}

	return the_time.Unix(), err
}
func RFC3339ToCSTInt64(value string) (int64, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return 0, err
	}
	return ts.Unix(), nil
}

var (
	cst *time.Location
)

// CSTLayout China Standard Time Layout
const CSTLayout = "2006-01-02 15:04:05"

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}
}

// RFC3339ToCSTLayout convert rfc3339 value to china standard time layout
func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	return ts.In(cst).Format(CSTLayout), nil
}
