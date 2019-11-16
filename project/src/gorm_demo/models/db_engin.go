package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func Init(){
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)
	
	dbType = "mysql"
	dbName = "go_demo"
	user = "root"
	password = "Qq1234567"
	host = "localhost"
	tablePrefix = "cs"
	
	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Fatal("数据库连接失败...")
	}else {
		log.Println("数据库连接成功...")
	}
	
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer DB.Close()
}

