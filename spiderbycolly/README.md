# help_center/spiderbycolly article topGameFi ETH transcation lock 


##项目说明
~~~ds
├── cmd
│   ├── ethereum.go
│   ├── root.go
│   ├── server.go
│   └── help_center/spiderbycollyCmd.go
├── common
│   ├── api
│   │   ├── api.go
│   │   └── api_test.go
│   ├── logger.go
│   └── types
│       └── types.go
├── config
│   ├── config.go
│   └── config.yaml
├── database
│   ├── address.go
│   ├── address_test.go
│   ├── base.go
│   ├── batch
│   │   └── batch.go
│   ├── block.go
│   ├── bybitArt.go
│   ├── bybitArtNewly.go
│   ├── ckoGameFi.go
│   ├── clickhouse.go
│   ├── cmcGameFi.go
│   ├── ethTrans.go
│   ├── schedule.go
│   ├── slateArt.go
│   ├── store
│   │   └── model.go
│   ├── symbol.go
│   ├── transaction.go
│   └── transfer.go
├── go.mod
├── go.sum
├── logs
├── main.go
└── help_center/spiderbycollyService
├── ethereum
│   ├── ethereum.go
│   ├── schedule
│   │   ├── logs
│   │   ├── schedule.go
│   │   ├── transaction.go
│   │   ├── transaction_test.go
│   │   └── transfer.go
│   └── util
│       ├── ethereum.go
│       ├── ethereum_test.go
│       └── myEth.go
├── model
│   └── TopGameFi.go
├── schedule
│   ├── arthelp_center/spiderbycolly.go
│   ├── arthelp_center/spiderbycolly_test.go
│   ├── logs
│   │   └── server_2022-02-15.log
│   ├── schedule.go
│   └── topGameFi.go
├── server
│   ├── controller
│   │   ├── controller.go
│   │   └── topGameFi.go
│   ├── server.go
│   └── service
│       └── topGameFi.go
├── help_center/spiderbycolly.go
└── util
├── TopGameFiKingData.go
├── TopGameFi_test.go
├── bybitArt.go
├── cyrptoSlateArt.go
├── dappreader.go
├── gameWayArt.go
├── retryFunc.go
├── sortObj.go
├── timeParse.go
├── topCmcGameFi.go
└── topCoinGecko.go

* 配置文件 ：./config.yaml/
  -- Mysql数据库配置：database

* run spider article and topGameFi conmand: 
    - go run main.go spiderCmd
