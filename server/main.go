package main

import (
	"fmt"
	"github.com/shuxiaoningAK/ADDD_DOUYIN/server/config"
	"github.com/shuxiaoningAK/ADDD_DOUYIN/server/router"
)

func main() {
	config.InitConfig()
	//repository.InitGormMysql(config.CONFIG.Mysql)
	r := router.InitRouters()
	if err := r.Run(); err != nil {
		panic(fmt.Errorf("%s \n", err))
	}
}
