package schedule

import (
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/database"
	"help_center/spiderbycolly/spiderService/model"
	"help_center/spiderbycolly/spiderService/util"
	"log"
	"sort"
	"time"

	"github.com/robfig/cron"
)

func createArtSpider(c *cron.Cron) (err error) {
	syncBybit()
	err = c.AddFunc("@every 24h", syncBybit)
	if err != nil {
		log.Print(err)
		return
	}
	synCmc()
	err = c.AddFunc("@every 24h", synCmc)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("createArtSpider success ")
	return
}

func synCmc() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
			common.Logger.Info("run time panic: %v", err)
		}
	}()

	log.Print("syncBycmc")
	err := util.Retry(3, 1*time.Second, cmcArt)
	if err != nil {
		common.Logger.Info("更新失败cmcart:", err)
		log.Print(err)
	}
}

func syncBybit() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
			common.Logger.Info("run time panic: %v", err)
		}
	}()
	err := util.Retry(3, 1*time.Second, bybitHighly)
	if err != nil {
		common.Logger.Info("更新失败bybitHighly:", err)
		log.Print(err)
	}
	err = util.Retry(3, 1*time.Second, bybitNewly)

	if err != nil {

		common.Logger.Info("更新失败bybitHighly:", err)

		log.Print(err)
	}

}

func bybitHighly() error {
	var (
		err  error
		resM []model.BybitArticle
	)
	db := database.DB()
	err = db.DeleteBybitArt()
	if err != nil {
		return err
	}

	resM, err = util.GetArticleBybitArt("res.Link")

	if err != nil {
		return err
	}

	sort.Sort(util.BybitArticleSlice(resM))
	err = db.SaveBybitHighLightArt(resM)
	if err != nil {
		return err
	}
	return nil
}

func bybitNewly() error {
	var (
		err  error
		resM []model.BybitNewlyArticle
	)
	db := database.DB()
	res, err := db.GetBybitNewlyArt()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		resM, err = util.GetNewArticleBybitArt("res[0].Link")
	} else {
		resM, err = util.GetNewArticleBybitArt(res[0].Link)
	}
	if err != nil {
		return err
	}
	sort.Sort(util.BybitNewlyArticleSlice(resM))
	err = db.SaveBybitNewlyArt(resM)
	if err != nil {
		return err
	}
	return nil
}

func cmcArt() error {
	var (
		err  error
		resM []model.SlateArticle
	)
	db := database.DB()
	res, err := db.GetSlateArt()

	if err != nil {
		common.Logger.Info("database 获取失败:", err)
		return err
	}

	if len(res) != 0 {
		resM, err = util.GetArticleCryptoSlate(res[0].Link)
	} else {
		resM, err = util.GetArticleCryptoSlate("res[0].Link")
	}
	if err != nil {
		common.Logger.Info(" 爬取失败:", err)
		return err
	}
	sort.Sort(util.SlateArticleSlice(resM))
	var resFinal []model.SlateArticle
	resMap, err := db.GetManySlateArt()
	for _, item := range resM {
		if resMap[item.Link] {
		} else {
			resFinal = append(resFinal, item)
		}
	}
	err = db.SaveSlateArt(resFinal)
	if err != nil {
		log.Print(err)
		common.Logger.Info(" 插入失败:", err)
		return err
	}
	return nil
}
