package main

import (
	"dictionary-of-chinese/router"

	"github.com/gin-gonic/gin"
)

func main() {

	routers := router.Router{}
	g := gin.Default()
	routers.InitRouter(g)
	g.Run(":8089")

}
