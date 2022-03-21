package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type GasResponse struct {
	Slow    GasData `json:"slow"`
	Normal  GasData `json:"normal"`
	Fast    GasData `json:"fast"`
	Instant GasData `json:"instant"`
}

type GasData struct {
	GWei float64 `json:"gwei"`
	Usd  float64 `json:"usd"`
}

// 查询当前GasPrice
func CurrentGasPrice() (gasPrice uint64, err error) {

	req, err := http.NewRequest("GET", "https://ethgas.watch/api/gas", nil)
	if err != nil {

		return
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {

		return
	}

	// 读取相应结果
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return
	}
	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("通过ethgas获取gas失败:%v", string(respBody))
		return
	}

	var result GasResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {

		return
	}

	gasPrice = uint64(result.Normal.GWei) * uint64(math.Pow10(9))
	return
}
