package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// 数据库配置信息
const (
	userName = "root"
	password = "Sxn@163!"
	ip       = "bj-cynosdbmysql-grp-pvwrb8uw.sql.tencentcdb.com"
	port     = "24295"
	dbName   = "douyin"
)

func InitDB() {

	db, err := gorm.Open("mysql", "root:Sxn@163!@(bj-cynosdbmysql-grp-pvwrb8uw.sql.tencentcdb.com:24295)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	fmt.Println("connect success")
	defer db.Close()

}
