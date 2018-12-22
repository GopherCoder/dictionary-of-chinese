package router

import (
	"dictionary-of-chinese/handler/idiom"
	"dictionary-of-chinese/handler/proverb"
	"dictionary-of-chinese/handler/word"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) InitRouter(g *gin.Engine, handler ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(handler...)

	g.GET("/health", func(context *gin.Context) {
		context.JSON(
			http.StatusOK, gin.H{
				"ping": "pong",
			})
	})
	v1 := g.Group("/v1/api")
	{
		idiom.Register(v1)
		proverb.Register(v1)
		word.Register(v1)
	}

	g.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK,
			fetchPath(g))
	})
	return g
}

func fetchPath(g *gin.Engine) []string {
	routers := g.Routes()
	var paths []string
	for _, router := range routers {
		paths = append(paths, router.Path)
	}
	return paths
}
