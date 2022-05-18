package conf

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database(connString string) {
	if db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		PrepareStmt:            true, //执行任何 SQL 时都创建并缓存预编译语句
		SkipDefaultTransaction: true, //禁用默认事务功能
	}); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(20)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Second * 30)
		DB = db
		migration() //迁移表结构
	}

}
