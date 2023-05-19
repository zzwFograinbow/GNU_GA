package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"Gin_study_new/pkg/setting"
)

var db *gorm.DB

//type Model struct {
//	ID int `json:"id"`
//}

type Rec struct {
	ReID          int    `grom:"column:re_id"`
	PatID         string `grom:"column:pat_id"`
	UserID        string `grom:"column:use_id"`
	Type          string `grom:"column:type"`
	Status        int    `grom:"column:status"`
	Create_time   string `grom:"column:create_time"`
	Done_time     string `grom:"column:done_time"`
	AlgoServer_ip string `grom:"column:algo_server_ip"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string { //这个defaultTableName是根据操作数据库时候结构体名字来自动推断的，还给小写了真是谢谢
		return tablePrefix + defaultTableName
	}
	/*
		defaultTableName是作为参数传递给gorm.DefaultTableNameHandler函数的。在GORM中，defaultTableName表示GORM根据模型的结构体名自动推断出的表名，它是一个字符串类型的变量。
		当您使用GORM创建模型时，如果您没有为该模型定义一个表名，GORM将会使用默认表名处理程序推断出一个表名，这个表名就是defaultTableName。默认表名处理程序会将结构体名称转换为小写字母，并使用下划线分隔单词来生成表名。
	*/

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
