package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	initRounter()

	r.run() // listen and serve on 0.0.0.0:8080
}
