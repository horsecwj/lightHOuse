package schedule

import (
	"fmt"
	"help_center/config"
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"help_center/spiderbycolly/spiderService/util"
	"log"

	"os"
	"sort"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func initDatabse() {
	viper.AddConfigPath("../../config")
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {

		fmt.Println("配置读取出错: ", err)
		return
	}

	// 监听配置
	viper.OnConfigChange(func(in fsnotify.Event) {

		config.RefreshConf()
	})
	viper.WatchConfig()
	// 尝试初始化数据库连接
	err := database.Init(true)
	// 关闭连接
	//defer database.CloseConn()

	if err != nil {

		common.Logger.Info("数据库初始化失败:", err)
		os.Exit(1)
		return
	}
}

func TestRun(t *testing.T) {
	initDatabse()
	db := database.DB()
	res, _ := util.GetTopGameFiCoinCko()

	db.SaveTopCkoGameFi(res)
}

func TestRun2(t *testing.T) {
	initDatabse()
	db := database.DB()
	err := db.DeleteTopCkoGameFi()
	fmt.Print(err)
}
func TestRun3(t *testing.T) {
	initDatabse()
	db := database.DB()
	res, _ := util.GetArticleBybitArt("")
	sort.Sort(util.BybitArticleSlice(res))
	_ = db.SaveBybitHighLightArt(res)
}

func TestRun4(t *testing.T) {
	//initDatabse()
	//db := database.DB()
	//res,err :=db.GetBybitArt()
	//ress,_  := util.GetArticleBybitArt(res.Link)
	//if err !=nil{
	//	return
	//}
	//sort.Sort( util.BybitArticleSlice(ress))
	//_=db.SaveBybitHighLightArt(ress)
	//fmt.Print(res.Title)
}

func TestRun5(t *testing.T) {
	initDatabse()
	common.InitLogger("server")
	syncCoinGk()
}

func TestRun8(t *testing.T) {
	initDatabse()
	//synCmcGameFi()
	syncCoinGk()
}

//func TestRun7(t *testing.T) {
//	initDatabse()
//	db := database.DB()
//	res, err := db.GetBybitNewlyArtBy()
//	log.Printf(res.Time, err)
//
//	c := colly.NewCollector(
//		colly.MaxDepth(2),
//		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36"),
//	)
//	// 设置抓取频率限制
//	_ = c.Limit(&colly.LimitRule{
//		DomainGlob:  "*",
//		RandomDelay: 3 * time.Second, // 随机延迟
//	})
//
//	c.OnRequest(func(req *colly.Request) {
//		log.Println("Visiting", req.URL)
//	})
//	c.OnError(func(_ *colly.Response, err error) {
//		log.Println("Something went wrong:", err)
//	})
//	c.OnResponse(func(resp *colly.Response) {
//		log.Print(resp.StatusCode)
//	})
//	util.GetArticleBybitDetailSlate(c, "https://learn.bybit.com/trading/what-is-market-correction/")
//}

func TestRun9(t *testing.T) {
	initDatabse()
	_ = database.DB()
	cmcArt()
}

func TestRun10(t *testing.T) {
	initDatabse()
	_ = database.DB()
	bybitHighly()
}

func TestRun11(t *testing.T) {
	initDatabse()
	_ = database.DB()
	bybitNewly()
}

//DeletebitNewlyArtCoin

func TestRun19(t *testing.T) {
	initDatabse()
	db := database.DB()
	err := db.DeletebitNewlyArtCoin("1")
	if err != nil {
		return
	}

}
func TestRun12(t *testing.T) {
	initDatabse()
	_ = database.DB()
	cmcGameFi()
}

func TestRun13(t *testing.T) {
	initDatabse()
	_ = database.DB()
	cgkGameFi()
}

func TestRun6(t *testing.T) {

	c := cron.New()
	err := c.AddFunc("@every 10m", synCmc)
	if err != nil {
		log.Print(err)
		return
	}
	c.Run()
	select {}

}
