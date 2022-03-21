package database

import (
	"fmt"
	"help_center/config"
	batchinsert "help_center/spiderbycolly/database/batch"

	"github.com/spf13/cast"

	"log"
	"sync"
)

var (
	tableMap   sync.Map
	chOnceDo   sync.Once
	chOnceDoBc sync.Once
	options    []batchinsert.Option
	OsExit     = make(chan struct{}, 1)
	WGroup     sync.WaitGroup
	Bc         *batchinsert.BatchInsert
	Tc         *batchinsert.BatchInsert
)

func getOption() []batchinsert.Option {
	bs := config.CKDBConf().BlockSize
	host := fmt.Sprintf("%s:%s", config.CKDBConf().Host, config.CKDBConf().Port)
	un := config.CKDBConf().Username     //  config.Config.GetString("ClickHouse.UserName")
	pwd := config.CKDBConf().Password    //config.Config.GetString("ClickHouse.Password")
	db := config.CKDBConf().Database     //config.Config.GetString("ClickHouse.Database")
	wt := config.CKDBConf().WriteTimeout // config.Config.GetInt("ClickHouse.WriteTimeout")
	rt := config.CKDBConf().ReadTimeout  //config.Config.GetInt("ClickHouse.ReadTimeout")
	debug := config.CKDBConf().DeBug     //config.Config.GetBool("ClickHouse.DeBug")
	fp := config.CKDBConf().FlushPeriod  // config.Config.GetDuration("ClickHouse.FlushPeriod")

	options = append(options,
		batchinsert.WithHost(host),
		batchinsert.WithUserInfo(un, pwd),
		batchinsert.WithDatabase(db),
		batchinsert.WithWriteTimeOut(cast.ToInt(wt)),
		batchinsert.WithReadTimeout(cast.ToInt(rt)),
		batchinsert.WithBlockSize(cast.ToInt(bs)),
		batchinsert.WithDebug(cast.ToBool(debug)),
		batchinsert.WithFlushPeriod(cast.ToDuration(fp)),
	)
	return options
}

func chOnceDbTc(tableName string) *batchinsert.BatchInsert {
	chOnceDo.Do(func() {
		options = getOption()
		var err error
		sqlStr := fmt.Sprintf("INSERT INTO %s (Hash, BlockHash, Nonce, BlockNumber, TransactionIndex,FromAddr,ToAddr,Value,Gas,GasPrice,BlockTimestamp,Data) "+
			"                                      VALUES (?,  ?,       ?,      ?,          ?,              ?,        ?,     ?,    ?, ?,        ?,           ?)", tableName)
		Tc, err = batchinsert.New(sqlStr, options...)
		if err != nil {
			return
		}
	})
	return Tc
}

func chOnceDbBc(tableName string) *batchinsert.BatchInsert {
	chOnceDoBc.Do(func() {
		options = getOption()
		//irr := Bc.Insert( a.Number, a.Hash, a.Timestamp, a.TxSize,a.Difficulty)
		var err error
		sqlStr := fmt.Sprintf("INSERT INTO %s ( Number, Hash, Timestamp, TxSize,Difficulty) "+
			"                                      VALUES (?,  ?,       ?,      ?,          ?)", tableName)
		Bc, err = batchinsert.New(sqlStr, options...)
		if err != nil {
			return
		}
	})

	return Bc
}

func CreatTransTable(tableName string) error {
	db := chOnceDbTc(tableName)
	table := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			Hash FixedString(100),
				BlockHash FixedString(100),
			Nonce UInt64,
	BlockNumber UInt64,
TransactionIndex UInt8,
FromAddr FixedString(100),
ToAddr FixedString(100),
	Value  UInt64,
	Gas UInt64,
GasPrice UInt64,
BlockTimestamp UInt64,
Data  FixedString(11150)
		) ENGINE = MergeTree()
          PARTITION BY toDate(BlockTimestamp)
          PRIMARY KEY Hash
          SETTINGS index_granularity = 8192;`, tableName)
	tmp := Tc
	log.Print(tmp)
	_, trr := db.DB().Exec(table) //
	if trr != nil {
		log.Println(trr.Error())
		return trr
	}
	return trr
}

func CreatBlockTable(tableName string) error {
	db := chOnceDbBc(tableName)
	table := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			Hash FixedString(100),
			Number UInt64,
			Timestamp UInt64,
			TxSize UInt8,
			Difficulty UInt64
		) ENGINE = MergeTree()
          PARTITION BY toDate(Timestamp)
          PRIMARY KEY Hash
          SETTINGS index_granularity = 8192;`, tableName)

	_, trr := db.DB().Exec(table)
	if trr != nil {
		log.Println(trr.Error())
		return trr
	}
	return trr
}

func (a *TransInfo) InsertTransTable(tableName string) {

	var ch chan TransInfo
	ch = make(chan TransInfo, 10)
	go ExecC(tableName, ch)
	ch <- *a
}

func (a *MyBlockInfo) InsertTransTable(tableName string) {

	var ch chan MyBlockInfo
	ch = make(chan MyBlockInfo, 10)
	go ExecCBlock(tableName, ch)
	ch <- *a
}

func ExecC(tableName string, args chan TransInfo) {

	for {
		select {
		case <-OsExit:
			WGroup.Add(1)
			brr := Bc.Close()
			if brr != nil {
				log.Println(brr.Error())
			}
			tableMap.Delete(tableName)
			WGroup.Done()
			break
		default:
			a := <-args
			irr := Tc.Insert(a.Hash, a.BlockHash, a.Nonce, a.BlockNumber, a.TransactionIndex, a.FromAddr, a.ToAddr, a.Value, a.Gas,
				a.GasPrice, a.BlockTimestamp, a.Data)

			if irr != nil {
				log.Println(irr.Error())
				break
			}
		}
	}
}

func ExecCBlock(tableName string, args chan MyBlockInfo) {

	for {
		select {
		case <-OsExit:
			WGroup.Add(1)
			brr := Bc.Close()
			if brr != nil {
				log.Println(brr.Error())
			}
			tableMap.Delete(tableName)
			WGroup.Done()
			break
		default:
			a := <-args
			irr := Bc.Insert(a.Number, a.Hash, a.Timestamp, a.TxSize, a.Difficulty)
			if irr != nil {
				log.Println(irr.Error())
				break
			}
		}
	}
}
