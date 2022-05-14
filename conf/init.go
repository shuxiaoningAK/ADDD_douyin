package conf

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)       //去除生成表后的s
	db.DB().SetMaxIdleConns(20)  //设置连接池中空闲连接的最大数量
	db.DB().SetMaxOpenConns(100) //设置打开数据库连接的最大数量
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration() //迁移表结构
}
