package main

import (
	"douyin/conf"
	"douyin/routes"
	"fmt"
)

func main() {
	//读入配置
	conf.Init()
	//请求转发路由
	r := routes.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("服务器启动发生错误,信息如下:", err) //TODO 后续可将服务端产生错误写入日志文件中
	}
}
