package database

import (
	"fmt"
	"help_center/config"
	"help_center/spiderbycolly/spiderService/model"

	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *DBConn
var once sync.Once

type DBConn struct {
	*gorm.DB
}

func Init(initKeyDB bool) error {

	normalDB := DB()
	if normalDB == nil {
		log.Print(fmt.Errorf("初始化数据库失败"))
		return fmt.Errorf("初始化数据库失败")
	}

	return nil
}

type TransactionHandle func(tx *DBConn) (interface{}, error)

func (db *DBConn) Transaction(handler TransactionHandle) (interface{}, error) {

	tx := &DBConn{db.Begin()}

	result, err := handler(tx)

	if err != nil {

		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return result, nil
}

func create(config *config.DatabaseConfig) *DBConn {

	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true"

	params := []interface{}{config.Username, config.Password, config.Host, config.Port, config.Database}
	uri := fmt.Sprintf(format, params...)
	db, err := gorm.Open("mysql", uri)

	if err != nil {

		return nil
	}

	db.LogMode(config.LogMode)
	db.SingularTable(true)

	// 表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {

		return "" + defaultTableName
	}

	//db.SetLogger(server_log.SQLLogger{})
	return &DBConn{db}
}

// 获取数据库实例
func DB() *DBConn {
	once.Do(func() {

		dbConfig := config.DBConf()
		if dbConfig == nil {

			return
		}

		db = create(dbConfig)

	})

	return db
}
func ckDB() *DBConn {
	once.Do(func() {

		dbConfig := config.CKDBConf()
		if dbConfig == nil {

			return
		}

	})

	return db
}

// 获取原始连接对象
func OriginDB() *gorm.DB {

	return DB().DB
}

// 关闭数据库连接
func CloseConn() {

	err := db.Close()
	if err != nil {

		fmt.Println(err)
		return
	}

}

// 创建 & 更新表结构
func AutoMigrate() {
	var tables []interface{}
	tables = append(tables, &model.TopCkoGameFi{}, &model.TopCmkGameFi{}, &model.TopCmkGameFiLosers{})

	tables = append(tables, &model.BybitArticle{}, &model.SlateArticle{})
	tables = append(tables, &MyBlockInfo{}, &TransInfo{})

	tables = append(tables, &model.BybitNewlyArticle{}, &BcwBlockNumber{})

	//tables = append(tables,&BlockInfo{})

	// 创建表结构
	for _, table := range tables {
		createTable(table)
	}
}

// 创建表
func createTable(table interface{}) {
	db := DB()
	if !db.HasTable(table) {

		if err := db.CreateTable(table).Error; err != nil {

			fmt.Print(err)
		}
	}
	db.AutoMigrate(table)
}
