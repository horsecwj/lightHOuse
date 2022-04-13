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

	c.OnHTML("div[class='highlight-section'] div[class='ant-row'] div[class='ant-col highlight-section-post ant-col-xs-18 ant-col-md-8'] ", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			link, isAlive := ts.Find("a[class='no-style']").Attr("href")
			if !isAlive {
				return
			}
			link = "https://learn.bybit.com" + link
			if link == titleStart {
				HighlightArtFlag = false
			}
			if HighlightArtFlag {
				oT, _ := ts.Find("div[class='post-card-description']").Attr("title")
				tempT, _ := ts.Find("div[class='post-card-thumbnail'] div").Attr("style")
				timeStr := tempT
				var ssr2 string
				ssr := strings.Split(timeStr, "background-image:url('")
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
				ress, err := http.Get(ssr2)
				if err != nil {
					return
				}
				var data []byte
				res := GetArticleBybitDetailSlate(c, link)
				res.Title = "title"
				res.OverView = "Overview"
				res.Link = link
				temp := model.BybitArticle{Title: res.Title, OverView: res.OverView, Link: res.Link,
					Article: res.Article, Time: res.Time, Timestamp: res.Timestamp, Articletext: res.Articletext, Pic: res.Pic}
				if ress != nil {
					data, err = ioutil.ReadAll(ress.Body)
					if err != nil {
						return
					}
					temp.Pic = string(data)
				}
				if len(oT) != 0 {
					temp.OverView = string(data)
				}
				if len(res.Article) != 0 {
					ArrTopGameFi = append(ArrTopGameFi, temp)
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
		log.Print(resp.StatusCode)
	})
	c.OnHTML("div[class='highlight-section']", func(element *colly.HTMLElement) {
	})
	c.OnHTML("div[class='latest-posts-section-posts'] div[class='ant-row'] div[class='ant-col latest-posts-section-post ant-col-xs-24 ant-col-md-12'] ", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			link, isAlive := ts.Find("a[class='no-style']").Attr("href")
			if !isAlive {
				return
			}
			link = "https://learn.bybit.com" + link
			if link == titleStart {
				newArtFlag = false
			}
			if newArtFlag {
				oT, _ := ts.Find("div[class='post-card-description']").Attr("title")
				tempT, _ := ts.Find("div[class='post-card-thumbnail'] div").Attr("style")
				timeStr := tempT
				var ssr2 string
				ssr := strings.Split(timeStr, "background-image:url('")
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
				ress, err := http.Get(ssr2)
				if err != nil {
					return
				}
				var data []byte
				res := GetArticleBybitDetailSlate(c, link)
				res.Title = "title"
				res.OverView = "Overview"
				res.Link = link
				temp := model.BybitNewlyArticle{Title: res.Title, OverView: res.OverView, Link: res.Link,
					Article: res.Article, Time: res.Time, Timestamp: res.Timestamp, Articletext: res.Articletext, Pic: res.Pic}
				if ress != nil {
					data, err = ioutil.ReadAll(ress.Body)
					if err != nil {
						return
					}
					temp.Pic = string(data)
				}
				if len(oT) != 0 {
					temp.OverView = string(data)
				}
				if len(res.Article) != 0 {
					ArrTopGameFi = append(ArrTopGameFi, temp)
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
	collector.OnHTML("div[class='post-detail-content']", func(elem *colly.HTMLElement) {
		art, err := elem.DOM.Html()
		artText := elem.DOM.Text()
		if err != nil {
			log.Print(err)
		} else {
			tempBybitArticle.Article = art
			tempBybitArticle.Articletext = artText
		}
	})

	collector.OnHTML("div[class='ant-col post-content ant-col-xs-24 ant-col-md-16']", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			title := ts.Find("h1[class='post-detail-title']").Text()
			tempBybitArticle.Title = title
		})
	})

	_ = collector.Visit(url)
	collector.Wait()
	return tempBybitArticle
}
