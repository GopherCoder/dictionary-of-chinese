package proverb

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/proverb/keys/:key", GetProverbByKeyHandler)
	r.GET("/proverb/ids/:id", GetProverbByIdHandler)
	r.GET("proverb/samples", GetProverbAtRandomHandler)
}
