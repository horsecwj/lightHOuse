package util

import (
	"log"
	"testing"
)

func TestSendRequest(t *testing.T) {
	//GetTopGameFiCoinMarket()
}
func TestGetArticleBybit(t *testing.T) {
	//res := GetArticleBybitArt()
	GetTopGameKingData()
}

func TestGetArticleBybitArt(t *testing.T) {
	getDappReader()
}
func TestRun10(t *testing.T) {
	//Feb. 22, 2022 at 6:00 am UTC
	//Feb. 15, 2022 at 1:30 pm UTC
	//Feb. 22, 2022 at 12:40 pm UTC
	//Mar. 1, 2022 at 2:30 am UTC
	timeInt, err := timeParse("Mar. 1, 2022 at 2:30 am UTC")
	log.Print(timeInt, err)
}

func TestGetArticleBybitDetailSlate(t *testing.T) {
	RFC3339Str := "2020-11-08T08:18:46+08:00"
	//RFC3339Str :=  2022-02-07T09:17:17+08:00
	cst, err := RFC3339ToCSTLayout(RFC3339Str)
	if err != nil {
		log.Println(err)
	}
	log.Println(cst)
}

func TestGetNewArticleBybitArt(t *testing.T) {
	RFC3339Str := "2020-11-08T08:18:46+08:00"
	RFC3339Str = "2022-02-18T13:19:09+08:00"
	//RFC3339Str :=  2022-02-07T09:17:17+08:00
	cst, err := RFC3339ToCSTInt64(RFC3339Str)
	if err != nil {
		log.Println(err)
	}
	log.Println(cst)

}
