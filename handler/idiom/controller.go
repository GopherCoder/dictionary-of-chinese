package idiom

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdiomsByNameHandler(context *gin.Context) {

}

func GetIdiomsByIdHandler(context *gin.Context) {
	var params string
	params = context.Param("id")
	id, _ := strconv.Atoi(params)
	totalNumber := totalNumber()
	if id < 0 || id > totalNumber {
		ResponseIdiom(
			context, http.StatusExpectationFailed,
			fmt.Sprintf("id should be less than %d", totalNumber))
		return
	}
	key := consistKey(params)
	fmt.Println(key, "key")
	if exists := isKeyExistInRedis(key); !exists {
		ResponseIdiom(context, http.StatusInternalServerError, fmt.Sprintf("record: %s not found", key))
		return
	}
	result := hashGetAllByKey(key)
	//fmt.Println(result, "result")
	ResponseIdiom(context, http.StatusOK, result)
}

func GetIdiomsAtRandomHandler(context *gin.Context) {

	var params sampleParams
	if err := context.ShouldBindQuery(&params); err != nil {
		ResponseIdiom(context, http.StatusExpectationFailed, fmt.Sprintf("params fail"))
		return
	}
	fmt.Println(params, "params")
	number, _ := strconv.Atoi(params.Number)
	ResponseIdiom(context, http.StatusOK, numberHashGetAll(number))

}
