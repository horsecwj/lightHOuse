package util

import (
	"help_center/spiderbycolly/spiderService/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func GetArticleBybitArt(titleStart string) ([]model.BybitArticle, error) {
	// 创建Collector
	//newArtFlag := true
	HighlightArtFlag := true
	c := colly.NewCollector(
		// 设置用户代理
		colly.MaxDepth(2),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36"),
	)
	// 设置抓取频率限制
	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 3 * time.Second, // 随机延迟
	})
	var ArrTopGameFi = make([]model.BybitArticle, 0, 1)
	c.OnRequest(func(req *colly.Request) {
		log.Println("Visiting", req.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnResponse(func(resp *colly.Response) {
		//log.Print(resp.StatusCode)
	})

	c.OnHTML("div[class='vc_column_inner tdi_131  wpb_column vc_column_container tdc-inner-column td-pb-span12'] div[id='tdi_132'] ", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			s := ts.Find("div[class='td_module_flex td_module_flex_1 td_module_wrap td-animation-stack td-meta-info-hide ']")
			for i := range s.Nodes {
				str := s.Find("div[class='td-module-meta-info']").Eq(i)
				Overview := str.Eq(0).Find("div[class='td-excerpt']").Text()
				link, isAlive := str.Eq(0).Find("h3 a").Attr("href")
				if !isAlive {
					continue
				}
				title := str.Eq(0).Find("h3 a").Text()
				if link == titleStart {
					HighlightArtFlag = false
				}
				if HighlightArtFlag {
					res := GetArticleBybitDetailSlate(c, link)
					res.Title = title
					res.OverView = Overview
					res.Link = link

					if len(res.Article) != 0 {
						ArrTopGameFi = append(ArrTopGameFi, res)
					}
				}
			}
		})
	})

	err := c.Visit("https://learn.bybit.com/")
	if err != nil {
		return ArrTopGameFi, err
	}
	c.Wait()
	return ArrTopGameFi, nil
}

func GetNewArticleBybitArt(titleStart string) ([]model.BybitNewlyArticle, error) {

	newArtFlag := true
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36"),
	)
	// 设置抓取频率限制
	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 3 * time.Second, // 随机延迟
	})
	var ArrTopGameFi = make([]model.BybitNewlyArticle, 0, 1)
	c.OnRequest(func(req *colly.Request) {
		log.Println("Visiting", req.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnResponse(func(resp *colly.Response) {
		//log.Print(resp.StatusCode)
	})

	c.OnHTML("div[class='td_block_wrap tdb_loop tdi_153 td-h-effect-up-shadow td_with_ajax_pagination td-pb-border-top td_block_template_2 tdb-category-loop-posts'] div[id='tdi_153'] ", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			s := ts.Find("div[class='tdb_module_loop td_module_wrap td-animation-stack td-meta-info-hide ']")
			for i := range s.Nodes {
				str := s.Find("div[class='td-module-meta-info']").Eq(i)
				title := str.Eq(0).Find("h3 a").Text()
				Overview := str.Eq(0).Find("div[class='td-excerpt']").Text()
				link, isAlive := str.Eq(0).Find("h3 a").Attr("href")
				if !isAlive {
					continue
				}
				if link == titleStart {
					newArtFlag = false
				}
				if newArtFlag {

					res := GetArticleBybitDetailSlate(c, link)
					res.Title = title
					res.OverView = Overview
					res.Link = link

					temp := model.BybitNewlyArticle{Title: res.Title, OverView: res.OverView, Link: res.Link,
						Article: res.Article, Time: res.Time, Timestamp: res.Timestamp, Articletext: res.Articletext, Pic: res.Pic}
					if len(res.Article) != 0 {
						ArrTopGameFi = append(ArrTopGameFi, temp)
					}
				}
			}
		})
	})

	err := c.Visit("https://learn.bybit.com/")
	if err != nil {
		return ArrTopGameFi, err
	}
	c.Wait()
	return ArrTopGameFi, nil

}

func GetArticleBybitDetailSlate(collector *colly.Collector, url string) model.BybitArticle {

	collector = collector.Clone()
	_ = collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})
	time.Sleep(2 * time.Second)
	collector.OnRequest(func(request *colly.Request) {
		log.Println("start visit: ", request.URL.String())
	})
	tempBybitArticle := model.BybitArticle{}
	collector.OnHTML("div[data-td-block-uid='tdi_103'] div[class='tdb-block-inner td-fix-index']", func(elem *colly.HTMLElement) {
		art, err := elem.DOM.Html()
		artText := elem.DOM.Text()
		if err != nil {
			log.Print(err)
		} else {
			tempBybitArticle.Article = art
			tempBybitArticle.Articletext = artText
		}
	})

	collector.OnHTML("div[data-td-block-uid='tdi_70'] div[class='tdb-block-inner td-fix-index']", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			timeStr, boolF := ts.Find("time").Attr("datetime")
			//formatTime,err:=time.Parse("2006-01-02 15:04:05",formatTimeStr)
			if boolF {
				tempBybitArticle.Time = timeStr
			}
			timeStamp, err := RFC3339ToCSTInt64(timeStr)
			if err == nil {
				tempBybitArticle.Timestamp = timeStamp
			}
		})
	})

	collector.OnHTML("div[data-td-block-uid='tdi_102']", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			timeStr := ts.Find("style").Eq(1).Nodes[0].FirstChild.Data
			var ssr2 string
			ssr := strings.Split(timeStr, "background: url('")
			if len(ssr) >= 2 {
				sssr := ssr[1]
				ssssr := strings.Split(sssr, "');")
				if len(ssssr) >= 1 {
					ssr2 = ssssr[0]
				}
			}
			if len(ssr2) == 0 {
				return
			}
			res, err := http.Get(ssr2)
			if err != nil {
				return
			}
			var data []byte
			if res != nil {
				data, err = ioutil.ReadAll(res.Body)
				if err != nil {
					return
				}
				tempBybitArticle.Pic = string(data)
			}
		})
	})

	_ = collector.Visit(url)
	collector.Wait()
	return tempBybitArticle
}
