package model

import (
	"ADDD_douyin/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

var err error

func InitDb() {

	db, err := gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName))

	if err != nil {
		fmt.Println("连接数据库错误，请检查参数", err)
	}

	// 禁用数据表的复数形式
	db.SingularTable(true)
	// 自动迁移
	db.AutoMigrate()
	//设置连接池中的最大连接数
	db.DB().SetMaxIdleConns(5)

	// 设置数据库中的最大连接数
	db.DB().SetMaxOpenConns(10)

	// 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	defer db.Close()
}
