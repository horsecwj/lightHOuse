package util

import (
	"help_center/spiderbycolly/spiderService/model"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func GetTopGameFiCoinMarket() ([]*model.TopCmkGameFi, error) {
	var ArrTopGameFi = make([]*model.TopCmkGameFi, 0, 30)
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:96.0) Gecko/20100101 Firefox/96.0"),
	)
	c.OnRequest(func(req *colly.Request) {
		log.Println("Visiting", req.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnResponse(func(resp *colly.Response) {

	})
	//c.OnHTML("div[class='uikit-col-md-8 uikit-col-sm-16']", func(element *colly.HTMLElement) {
	//
	//	element.DOM.Each(func(_ int, s *goquery.Selection) {
	//		str := s.Find("h3")
	//		log.Print(str)
	//	})
	//
	//})
	c.OnHTML("table tbody tr", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, s *goquery.Selection) {
			str := s.Find("td")
			name := str.Eq(1).Find("a div div p[color='text']").Text()
			name2 := str.Eq(1).Find("a div div p[color='text3']").Text()
			link, _ := str.Eq(1).Find("a div img").Attr("src")
			if len(ArrTopGameFi) <= 100 {
				res, _ := http.Get(link)
				data, _ := ioutil.ReadAll(res.Body)
				tplData := model.TopCmkGameFi{
					ID:        len(ArrTopGameFi) + 1,
					Coin:      name2,
					GameFi:    name,
					CoinPic:   string(data),
					Price:     str.Eq(2).Find("span").Text(),
					OneDay:    str.Eq(3).Find("span").Text(),
					DayVolume: str.Eq(4).Text(),
				}
				ArrTopGameFi = append(ArrTopGameFi, &tplData)
			}

		})
	})

	err := c.Visit("https://coinmarketcap.com/gainers-losers/")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	c.Wait()
	return ArrTopGameFi, err
}
