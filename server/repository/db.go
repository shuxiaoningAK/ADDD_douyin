package repository

import (
	"fmt"
	"github.com/shuxiaoningAK/ADDD_DOUYIN/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormMysql(mysqlConfig config.Mysql) {
	var err error
	if DB, err = gorm.Open(mysql.Open(mysqlConfig.Dsn()), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("Fatal error connect to mysql: %s \n", err))
	}
}
