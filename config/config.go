package config

import (
	"fmt"
	"help_center/spiderbycolly/common/types"
	"sync"

	"github.com/spf13/viper"
)

// ETHConfig ETH配置信息
type ETHConfig struct {
	URL            string      // 节点地址
	Chain          types.Chain // 链类型
	ApiGetGasPrice bool        // 使用 api查询当前 Gas
	ConfirmHeight  uint64      // 不可逆高度
	AddrLimit      uint64      // 单次生成地址
	MinGasPrice    uint64
	MaxGasPrice    uint64
	ERC20GasLimit  uint64
}

func ETHConf() *ETHConfig {

	return APPConf().Ethereum
}

var (
	appConfig  APPConfig
	configOnce sync.Once
)

func HECOConf() *ETHConfig {

	return APPConf().HECO
}

func BSCConf() *ETHConfig {

	return APPConf().BSC
}

//Database:
//  Host: "127.0.0.1"
//  Port: "3306"
//  LogMode: true
//  Username: "root"
//  Password: "123456"
//  Database: "ssss"

type APPConfig struct {
	Notice     bool            `toml:"notice"` // 是否发送充值通知（即发现交易立即通知，不等待不可逆)
	Debug      bool            `toml:"debug"`
	Database   *DatabaseConfig `toml:"database"`
	Server     *ServerConfig   `toml:"server"`
	Ethereum   *ETHConfig      `toml:"ethereum"`
	HECO       *ETHConfig      // HECO 节点配置
	BSC        *ETHConfig      // BSC 节点配置
	ClickHouse *ClickHouseConfig
}

// 数据库配置信息
type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	LogMode  bool   `toml:"log_mode"`
}

//ClickHouse:
//  Host: 127.0.0.1
//  Port: 9000
//  UserName: default
//  Password: 123456
//  Database: ethDatas

// ck数据库配置信息
type ClickHouseConfig struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	Database     string `toml:"database"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	ReadTimeout  string `toml:"read_timeout"`
	WriteTimeout string `toml:"write_timeout"`
	BlockSize    string `toml:"block_size"`
	DeBug        string `toml:"de_bug"`
	FlushPeriod  string `toml:"flush_period"`
}

// Server配置信息
type ServerConfig struct {
	Port        string `toml:"port"`
	SummedCheck bool   `toml:"summed_check"`
	VerifyHost  string `toml:"verify_host"`
}

// 项目配置
func APPConf() *APPConfig {

	configOnce.Do(func() {

		if err := viper.Unmarshal(&appConfig); err != nil {

			fmt.Println("读取配置出错:", err)
		}
	})
	//fmt.Println(appConfig, "123123123")
	return &appConfig
}

// 数据库配置
func DBConf() *DatabaseConfig {

	return APPConf().Database
}

// 数据库配置
func CKDBConf() *ClickHouseConfig {

	return APPConf().ClickHouse
}

func RefreshConf() {

	configOnce = sync.Once{}
}
