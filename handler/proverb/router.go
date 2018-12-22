package proverb

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/proverb/:key", GetProverbByKeyHandler)
	r.GET("/proverb/:id", GetProverbByIdHandler)
	r.GET("proverb/samples", GetProverbAtRandomHandler)
}
