package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//UkdGd2NGSmhaR0Z5Y0dGblpUMHhKbk5uY205MWNEMXRZWGdtWTNWeWNtVnVZM2s5VlZORUptWmxZWFIxY21Wa1BURW1jbUZ1WjJVOVpHRjVKbU5oZEdWbmIzSjVQV2RoYldWekpuTnZjblE5ZFhObGNpWnZjbVJsY2oxa1pYTmpKbXhwYldsMFBUSTI=

type autoGenerated struct {
	Page           int `json:"page"`
	ResultsPerPage int `json:"resultsPerPage"`
	PageCount      int `json:"pageCount"`
	ResultCount    int `json:"resultCount"`
	Dapps          []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		New       bool   `json:"new"`
		Statistic struct {
			Balance                 float64 `json:"balance"`
			BalanceInFiat           float64 `json:"balanceInFiat"`
			TotalBalanceInFiat      float64 `json:"totalBalanceInFiat"`
			Graph                   string  `json:"graph"`
			ExchangeRate            int     `json:"exchangeRate"`
			CurrencyName            string  `json:"currencyName"`
			TransactionCount        int     `json:"transactionCount"`
			UserActivity            int     `json:"userActivity"`
			VolumeInFiat            float64 `json:"volumeInFiat"`
			TotalVolumeInFiat       float64 `json:"totalVolumeInFiat"`
			TotalVolumeChangeInFiat float64 `json:"totalVolumeChangeInFiat"`
			Changes                 struct {
				Dau struct {
					Status string `json:"status"`
					Label  string `json:"label"`
				} `json:"dau"`
				Volume struct {
					Status string `json:"status"`
					Label  string `json:"label"`
				} `json:"volume"`
				Tx struct {
					Status string `json:"status"`
					Label  string `json:"label"`
				} `json:"tx"`
				TokenVolume struct {
					Status string `json:"status"`
					Label  string `json:"label"`
				} `json:"tokenVolume"`
				TotalVolume struct {
					Status string `json:"status"`
					Label  string `json:"label"`
				} `json:"totalVolume"`
				TotalBalance struct {
					Status string `json:"status"`
					Label  string `json:"label"`
				} `json:"totalBalance"`
			} `json:"changes"`
		} `json:"statistic"`
		GodzillaID int `json:"godzillaId"`
		Slugs1     struct {
			Multichain string `json:"multichain"`
			Hive       string `json:"hive"`
			Wax        string `json:"wax"`
		} `json:"slugs,omitempty"`
		Logo            string   `json:"logo"`
		DeepLink        string   `json:"deepLink"`
		MobileFriendly  bool     `json:"mobileFriendly"`
		Featured        bool     `json:"featured"`
		Protocols       []string `json:"protocols"`
		ActiveProtocols []string `json:"activeProtocols"`
		Category        string   `json:"category"`
		Tracked         bool     `json:"tracked"`
		Slugs2          struct {
			BinanceSmartChain string `json:"binance-smart-chain"`
		} `json:"slugs,omitempty"`
		Slugs3 struct {
			Multichain        string `json:"multichain"`
			Wax               string `json:"wax"`
			BinanceSmartChain string `json:"binance-smart-chain"`
		} `json:"slugs,omitempty"`
		Slugs4 struct {
			Wax string `json:"wax"`
		} `json:"slugs,omitempty"`
		Slugs5 struct {
			Multichain string `json:"multichain"`
			Wax        string `json:"wax"`
			Eos        string `json:"eos"`
		} `json:"slugs,omitempty"`
		Slugs6 struct {
			Multichain string `json:"multichain"`
			Ronin      string `json:"ronin"`
			Ethereum   string `json:"ethereum"`
		} `json:"slugs,omitempty"`
		Slugs7 struct {
			Eos string `json:"eos"`
		} `json:"slugs,omitempty"`
		Slugs8 struct {
			Multichain string `json:"multichain"`
			Ethereum   string `json:"ethereum"`
			Polygon    string `json:"polygon"`
		} `json:"slugs,omitempty"`
		Slugs9 struct {
			Ronin string `json:"ronin"`
		} `json:"slugs,omitempty"`
		Slugs10 struct {
			Solana string `json:"solana"`
		} `json:"slugs,omitempty"`
	} `json:"dapps"`
	Ad interface{} `json:"ad"`
}

type DappGameFi struct {
	Coin     string
	CateGory string
	Balance  float64
	Users    int
	Volume   float64
	Activity string
}

func getDappReader() {
	url := "https://dappradar.com/v2/api/dapps?params=DappRadarpage=1&range=day&category=games&sort=user&order=desc&limit=26"
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			return
		}
	}()
	body, _ := ioutil.ReadAll(res.Body)
	//log.Print(body)
	resStruct := new(autoGenerated)
	err = json.Unmarshal(body, resStruct)
	if err != nil {
		return
	}
	var resDapp = make([]DappGameFi, 0, 26)

	for _, item := range resStruct.Dapps {
		//if  item.Category != "games"{
		//	continue
		//}
		resDapp = append(resDapp, DappGameFi{
			Coin:     item.Name,
			CateGory: item.Category,
			Balance:  item.Statistic.BalanceInFiat,
			Users:    item.Statistic.UserActivity,
			Volume:   item.Statistic.VolumeInFiat,
			Activity: item.Statistic.Graph,
		})

		//if len(resDapp) >=10{
		//	break
		//}
	}
	//timeNow := time.Unix(resStruct.Time/1000, 0) //2017-08-30 16:19:19 +0800 CST
	//timeString := timeNow.Format("2006-01-02 15:04:05")
	return
}
