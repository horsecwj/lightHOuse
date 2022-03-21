package util

import (
	"context"
	"fmt"
	"help_center/config"
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/common/api"
	"help_center/spiderbycolly/common/types"
	"log"
	"math"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var ETHInstance *ETHClient
var HECOInstance *ETHClient
var BSCInstance *ETHClient
var ethOnce sync.Once
var hecoOnce sync.Once
var bscOnce sync.Once

const (
	GasLimit            = 21000
	ERC20GasLimit       = 80000
	TransferParamLength = 64

	ERC20TransferMethod = "a9059cbb"
	ERC20BalanceMethod  = "70a08231"
)

type ETHClient struct {
	ctx    context.Context
	client *ethclient.Client
	Config *config.ETHConfig
}

// ETHTransferParam 转账参数
type ETHTransferParam struct {
	Payer       string // 转出地址
	Receiver    string // 接收地址
	Quantity    uint64 // 数量
	GasPrice    uint64 // Gas Price
	DiffDecimal uint64 // 链上精度和本地存储精度差值
	Nonce       uint64 // nonce

	Type     TransferType // 转账类型
	Contract string       // 合约地址

	PriKey string // 私钥
}

func EVMInstance(c types.Chain) *ETHClient {

	if c == types.ChainHECO {

		return InitHECOInstance()
	}

	if c == types.ChainBSC {

		return InitBSCInstance()
	}
	return InitETHInstance()
}

func (p *ETHTransferParam) GasLimit() (gasLimit uint64) {

	gasLimit = GasLimit

	// ERC20 转账需要更多手续费
	if p.Type == TransferTypeERC20 {
		gasLimit = ERC20GasLimit
	}
	return
}

func (p *ETHTransferParam) GenerateData(method string) (data []byte) {

	// ERC20 转账需要更多手续费
	if p.Type == TransferTypeERC20 {
		// 修改为给合约转账
		receiver := p.Receiver
		quantity := p.Quantity

		// 替换转账对象及金额
		p.Receiver = p.Contract
		p.Quantity = 0

		// 构造InputData
		address := strings.Replace(receiver, "0x", "", 1)

		// 转换合约精度
		amount := big.NewInt(int64(quantity))

		decimal := int(p.DiffDecimal)
		amount = amount.Mul(amount, big.NewInt(int64(math.Pow10(decimal))))
		amountText := hexutil.EncodeBig(amount)

		amountText = strings.ReplaceAll(amountText, "0x", "")

		addressData := ZeroFix(TransferParamLength, address)
		amountData := ZeroFix(TransferParamLength, amountText)

		dataText := method + addressData + amountData
		data = ethCommon.Hex2Bytes(dataText)
	}

	return
}

type TransferType string

const (
	TransferTypeETH   TransferType = "ETH"
	TransferTypeERC20 TransferType = "ERC20"
)

// InitETHInstance 初始化单例
func InitETHInstance() *ETHClient {

	ethOnce.Do(func() {
		c := config.ETHConf()
		ETHInstance = createETHUtil(c)
	})

	return ETHInstance
}

// InitHECOInstance 初始化 HECO 单例
func InitHECOInstance() *ETHClient {

	hecoOnce.Do(func() {

		HECOInstance = createETHUtil(config.HECOConf())
	})

	return HECOInstance
}

// InitBSCInstance 初始化 BSC 单例
func InitBSCInstance() *ETHClient {

	bscOnce.Do(func() {

		BSCInstance = createETHUtil(config.BSCConf())
	})

	return BSCInstance
}

// 创建 ETHClient
func createETHUtil(config *config.ETHConfig) *ETHClient {
	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, config.URL)
	if err != nil {

		common.Logger.Error("创建EthereumClient失败:", err)
		return nil
	}

	return &ETHClient{
		ctx:    ctx,
		client: cli,
		Config: config,
	}
}

// LastBlockNumber 查询最新区块
func (c *ETHClient) LastBlockNumber() (number uint, err error) {

	block, err := c.client.BlockByNumber(c.ctx, nil)
	if err != nil {

		return
	}

	number = uint(block.Number().Uint64())
	return
}

// GetBlockByNumber 查询区块信息
func (c *ETHClient) GetBlockByNumber(number int64) (*ethTypes.Block, error) {

	return c.client.BlockByNumber(c.ctx, big.NewInt(number))
}

// GetSender 查询交易发送地址
func (c *ETHClient) GetSender(tx *ethTypes.Transaction, hash string, index uint64) (ethCommon.Address, error) {

	return c.client.TransactionSender(c.ctx, tx, ethCommon.HexToHash(hash), uint(index))
}

// GetTransaction 查询交易
func (c *ETHClient) GetTransaction(txHash string) (tx *ethTypes.Transaction, isPending bool, err error) {

	hash := ethCommon.HexToHash(txHash)
	return c.client.TransactionByHash(c.ctx, hash)
}

// GetTransactionReceipt 查询交易Header
func (c *ETHClient) GetTransactionReceipt(txHash string) (*ethTypes.Receipt, error) {

	hash := ethCommon.HexToHash(txHash)
	return c.client.TransactionReceipt(c.ctx, hash)
}

// GetCurrentGas 查询当前Gas Wei单位
func (c *ETHClient) GetCurrentGas() (gas uint64, err error) {

	// 先通过星火矿池的接口查询
	if c.Config.ApiGetGasPrice {
		if gas, err = api.CurrentGasPrice(); err != nil {
			common.Logger.Error("API 查询 Gas Price 出错:", err)
			var g *big.Int
			g, err = c.client.SuggestGasPrice(c.ctx)
			if err != nil {
				return
			}
			gas = g.Uint64()
		}
	} else {

		// 直接调用 RPC 接口查询
		var g *big.Int
		g, err = c.client.SuggestGasPrice(c.ctx)
		if err != nil {
			return
		}
		gas = g.Uint64()
	}

	return
}

// EstimateGas 估算Gas limit
func (c *ETHClient) EstimateGasLimit(param *ETHTransferParam) (limit uint64, err error) {
	// 直接转 ETH， 固定Limit
	if param.Type == TransferTypeETH {

		limit = GasLimit
		return
	}

	data := param.GenerateData(ERC20BalanceMethod)

	fromAddr := ethCommon.HexToAddress(param.Payer)
	toAddr := ethCommon.HexToAddress(param.Contract)

	limit, err = c.client.EstimateGas(c.ctx, ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Value: big.NewInt(0),
		Data:  data,
	})

	if limit < 50000 {

		limit = limit * 2
	}
	return
}

// GetNonce 查询Nonce
func (c *ETHClient) GetNonce(address string) (nonce uint64, err error) {

	addr := ethCommon.HexToAddress(address)
	nonce, err = c.client.PendingNonceAt(c.ctx, addr)

	return
}

// GetCode 查询地址是不是合约
func (c *ETHClient) GetCode(address string) ([]byte, error) {

	addr := ethCommon.HexToAddress(address)
	return c.client.CodeAt(c.ctx, addr, nil)
}

// GetBalance 查询以太坊余额
func (c *ETHClient) GetBalance(address string) (balance *big.Int, err error) {

	addr := ethCommon.HexToAddress(address)

	return c.client.BalanceAt(c.ctx, addr, nil)
}

// GetERC20Balance 查询
func (c *ETHClient) GetERC20Balance(contract, method, address string) (balance *big.Int, err error) {
	address = strings.ReplaceAll(address, "0x", "")
	method = strings.ReplaceAll(method, "0x", "")

	common.Logger.Info("查询ERC20:", contract, " 方法:", method, " 地址:", address)

	// 组装参数
	addressData := ZeroFix(TransferParamLength, address)

	dataText := method + addressData
	data := ethCommon.Hex2Bytes(dataText)

	balanceHex, err := c.GetBalanceAtContract(contract, data)
	if err != nil {

		common.Logger.Info("查询ERC20出错:", err)
		return
	}

	balanceHex = strings.Replace(balanceHex, "0x", "", 1)

	// 去掉中间的0 后加上 0x
	balanceHex = "0x" + ZeroTrim(balanceHex)

	if strings.EqualFold(balanceHex, "0x") {

		return big.NewInt(0), nil
	}

	return hexutil.DecodeBig(balanceHex)
}

// GetBalanceAtContract 查询ERC20 代币余额
func (c *ETHClient) GetBalanceAtContract(contract string, data []byte) (balance string, err error) {

	address := ethCommon.HexToAddress(contract)
	msg := ethereum.CallMsg{
		From:     address,
		To:       &address,
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     data,
	}

	resp, err := c.client.CallContract(c.ctx, msg, nil)
	if err != nil {

		return
	}

	return hexutil.Encode(resp), nil
}

// Transfer 发起ETH & ERC20 代币转账
func (c *ETHClient) Transfer(p *ETHTransferParam) (txHash string, err error) {
	// 归结手续费
	priKey, err := crypto.HexToECDSA(p.PriKey)
	if err != nil {

		err = fmt.Errorf("私钥不正确:%s", err)
		return
	}

	// 若没有指定 Nonce，则进行查询
	nonce := p.Nonce
	if nonce == 0 {
		nonce, err = c.GetNonce(p.Payer)
		if err != nil {

			err = fmt.Errorf("获取Nonce错误:%s", err)
			return
		}
	}

	// 若有指定 Gas，则使用指定值，否则查询当前
	gas := p.GasPrice
	if gas == 0 {
		gas, err = c.GetCurrentGas()

		if err != nil {

			err = fmt.Errorf("获取Gas错误:%s", err)
			return
		}
	}

	common.Logger.Info("当前 Gas Price: ", gas)

	// 根据参数信息获取GasLimit跟InputData
	gasLimit := p.GasLimit()
	data := p.GenerateData(ERC20TransferMethod)
	receiver := ethCommon.HexToAddress(p.Receiver)
	// common.HexToAddress("0x7481fd1FC439e043112EC02BD32d143bf1a5d73C")
	amount := big.NewInt(int64(p.Quantity))
	decimal := int(p.DiffDecimal)
	amount = amount.Mul(amount, big.NewInt(int64(math.Pow10(decimal))))
	revAdr := receiver.Hex()
	log.Print(revAdr)
	// 构造交易
	tx := ethTypes.NewTx(&ethTypes.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(int64(gas)),
		Gas:      gasLimit,
		To:       &receiver,
		Value:    amount,
		Data:     data,
	})

	// 签名交易
	signer := ethTypes.MakeSigner(params.MainnetChainConfig, nil)
	signedTx, err := ethTypes.SignTx(tx, signer, priKey)
	if err != nil {

		err = fmt.Errorf("签名交易失败: %s", err)
		return
	}

	// 发送交易
	err = c.client.SendTransaction(c.ctx, signedTx)
	if err != nil {

		return
	}

	txHash = signedTx.Hash().Hex()
	return
}

// 测算当前 ERC20 的手续费
func (c *ETHClient) GetERC20Gas(contract, wallet, addr string, quantity uint64) (gas uint64, err error) {

	// 先查询 Gas Price
	price, err := c.GetCurrentGas()
	if err != nil {

		return
	}

	// 再查询执行 Limit
	limit, err := c.EstimateGasLimit(&ETHTransferParam{
		Payer:    wallet,
		Receiver: addr,
		Quantity: quantity,
		Type:     TransferTypeERC20,
		Contract: contract,
	})
	if err != nil {

		return
	}

	gas = price * limit
	return
}

func IsValidAddress(addr string) bool {
	return ethCommon.IsHexAddress(addr)
}

// ZeroTrim 去除 hex string 多余的 0
func ZeroTrim(text string) string {

	if strings.HasPrefix(text, "0") {

		return ZeroTrim(strings.Replace(text, "0", "", 1))
	}

	return text
}

// ZeroFix EVM Input Data 补充0以达到长度
func ZeroFix(length int, raw string) string {

	return strings.Join(make([]string, length-len(raw), length-len(raw)), "0") + "0" + raw
}
