package word

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/words/name/:name", GetWordsByNameHandler)
	r.GET("/words/ids/:id", GetWordsByIdHandler)
	r.GET("/words/samples", GetWordsAtRandomHandler)
}
