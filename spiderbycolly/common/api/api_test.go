package api

import (
	"fmt"
	"testing"
)

func TestCurrentGasPrice(t *testing.T) {

	gasPrice, err := CurrentGasPrice()
	if err != nil {

		fmt.Println("查询 Gas Price 出错: ", gasPrice)
		return
	}

	fmt.Println("当前 Gas Price:", gasPrice)
}
