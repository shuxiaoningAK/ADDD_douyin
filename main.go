package main

import (
	"ADDD_douyin/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.InitRouter(r)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
