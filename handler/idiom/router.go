package idiom

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/idioms/name/:name", GetIdiomsByNameHandler)
	r.GET("/idioms/ids/:id", GetIdiomsByIdHandler)
	r.GET("/idioms/samples", GetIdiomsAtRandomHandler)
	r.GET("/idioms/rank", GetIdiomsRankHandler)
}
