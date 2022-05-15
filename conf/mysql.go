package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func initMySql(err error) {
	dsn := "root:root@(172.21.16.17:3306)/douyin?charsetset=utf8mb4&parseTime=True&loc=local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	fmt.Println("sucess link to mysql")
	defer DB.Close()

}
