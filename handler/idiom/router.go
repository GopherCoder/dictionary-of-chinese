package idiom

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/idioms/:name", GetIdiomsByNameHandler)
	r.GET("/idioms/:id", GetIdiomsByIdHandler)
	r.GET("/idioms/samples", GetIdiomsAtRandomHandler)
}
