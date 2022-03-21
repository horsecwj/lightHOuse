package util

import (
	"help_center/spiderbycolly/spiderService/model"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func GetArticleCryptoSlate(titleStart string) ([]model.SlateArticle, error) {
	// 创建Collector
	artFlag := true
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

	var ArrTopGameFi = make([]model.SlateArticle, 0, 1)
	c.OnRequest(func(req *colly.Request) {
		log.Println("Visiting", req.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnResponse(func(resp *colly.Response) {
		log.Print(resp.StatusCode)
	})
	//posts clearfix trending
	c.OnHTML("div[class='posts'] div[class='list-post clearfix ']", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, s *goquery.Selection) {

			link, _ := s.Find("a").Attr("href")
			if link == titleStart {
				artFlag = false
			}

			if artFlag {
				res := GetArticleCryptoDetailSlate(c, link)
				res.Link = link
				if len(res.Article) != 0 {
					ArrTopGameFi = append(ArrTopGameFi, res)
				}
			}

		})
	})
	// 查找下一页
	c.OnHTML("a[class='nextpostslink']", func(element *colly.HTMLElement) {
		href, found := element.DOM.Attr("href")
		// 如果有下一页，则继续访问
		if found && artFlag {
			_ = element.Request.Visit(element.Request.AbsoluteURL(href))
		}
	})
	err := c.Visit("https://cryptoslate.com/news/")
	if err != nil {
		return nil, err
	}
	c.Wait()
	return ArrTopGameFi, err
}

func GetArticleCryptoDetailSlate(collector *colly.Collector, url string) model.SlateArticle {

	collector = collector.Clone()
	_ = collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 3 * time.Second,
	})
	time.Sleep(4 * time.Second)
	collector.OnRequest(func(request *colly.Request) {
		log.Println("start visit: ", request.URL.String())
	})
	tempBybitArticle := model.SlateArticle{}
	collector.OnHTML("div[id='main']", func(elem *colly.HTMLElement) {

		elem.DOM.Each(func(_ int, ts *goquery.Selection) {
			s := ts.Find("div").Eq(0)
			title := s.Find("div[class='post-header article clearfix'] div[class='title clearfix ']").Find("h1").Text()
			overView := s.Find("p[class='post-subheading']").Text()
			timeStr := s.Find("div[class='post-meta clearfix'] div span[class='post-date']").Text()

			art, err := s.Find("div[class='post-box clearfix'] article").Html()
			arttext := s.Find("div[class='post-box clearfix'] article").Text()

			if err != nil {
				log.Print(err)
			} else {
				tempBybitArticle = model.SlateArticle{Title: title, OverView: overView, Article: art, Time: timeStr, Articletext: arttext}
			}
			timestamp, err := timeParse(timeStr)
			if err == nil {
				tempBybitArticle.Timestamp = timestamp
			}
		})
	})
	_ = collector.Visit(url)
	collector.Wait()
	return tempBybitArticle
}
