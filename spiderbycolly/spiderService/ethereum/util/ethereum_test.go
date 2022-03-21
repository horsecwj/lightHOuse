package util

import (
	"fmt"
	"help_center/config"
	"log"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestEthereumUtil_GetBalanceAtContract(t *testing.T) {

	ethUtil := createETHUtil(&config.ETHConfig{})

	zeroText := func(length int) string {

		return strings.Join(make([]string, length, length), "0") + "0"
	}

	method := "0x27e235e3"
	address := strings.ToLower("0x1E794280B67548b5d893988336c1b1b395Beb9d2")

	method = strings.ReplaceAll(method, "0x", "")
	address = strings.ReplaceAll(address, "0x", "")
	addressData := zeroText(64-len(address)) + address
	dataText := method + addressData
	data := common.Hex2Bytes(dataText)

	balanceHex, err := ethUtil.GetBalanceAtContract("0xdAC17F958D2ee523a2206206994597C13D831ec7", data)
	if err != nil {

		fmt.Println("查询余额失败:", err)
		return
	}

	balanceHex = strings.ReplaceAll(balanceHex, "0x", "")

	balance, err := hexutil.DecodeUint64(balanceHex)
	fmt.Println("balance of:", address, " is:", balance)
}

func TestEthereumUtil_GetTransaction(t *testing.T) {

	ethUtil := createETHUtil(&config.ETHConfig{})
	tx, _, err := ethUtil.GetTransaction("0x62963869d376ea61c0f511a2c5ba232494d35d94543ba83f796b6fc1870f8582")
	if err != nil {

		fmt.Println("查询交易出错 -> ", err)
		return
	}

	fmt.Println(tx.Hash().Hex())
	sender, err := ethUtil.GetSender(tx, "0xec5c1011cf9930bf704422c6bd76d5c9b94668df513a3c63b380a697247b2b75", 1)
	if err != nil {

		fmt.Println("查询发送者出错 -> ", err)
		return
	}

	fmt.Println("sender -> ", sender.Hex())
}

func TestInitETHInstance(t *testing.T) {
	transEth()

}

//以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，
//在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18。
//因为我们正在处理大数，我们得导入原生的Gomath和math/big包。这是您做的转换。
func TestInitHECOInstance(t *testing.T) {
	bal1 := GetBalance("0x51822AbC7de1dDA7b8C480C3C92FEbaCF2A3A65c")

	bal2 := GetBalance("0xBd47812fcE4C41C877ECc1583530B5Ff8DE75Fc5")
	log.Print(bal1, bal2)
}

func TestZeroFix(t *testing.T) {
	GetPrivateKey()
}
func TestInitETHInstance2(t *testing.T) {
	erc20Trans()
}

func TestInitETHInstance3(t *testing.T) {
	log.Print(len("1000000000000000000"))
}
func TestInitETHInstance4(t *testing.T) {
	normalErc20Trans()
}
