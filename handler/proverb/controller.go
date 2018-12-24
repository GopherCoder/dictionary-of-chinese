package proverb

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProverbByKeyHandler(context *gin.Context) {}

func GetProverbByIdHandler(context *gin.Context) {
	var params string
	params = context.Param("id")
	total := totalNumber()
	id, _ := strconv.Atoi(params)
	if !isSuitableId(id) {
		ResponseProverb(context, http.StatusInternalServerError, fmt.Sprintf("id should be less than %d", total))
		return
	}
	fullKey := consistKey(params)
	ResponseProverb(context, http.StatusOK, hashGetAllProverbByKey(fullKey))

}

func GetProverbAtRandomHandler(context *gin.Context) {
	var params sampleParams
	if err := context.ShouldBindQuery(&params); err != nil {
		ResponseProverb(context, http.StatusExpectationFailed, err.Error())
		return
	}
	//fmt.Println(params, "params")
	number, _ := strconv.Atoi(params.Number)
	//fmt.Println(params, number)
	results := hashGetAllProverbAtSamples(number)
	if results == nil {
		ResponseProverb(context, http.StatusOK, "record not found")
		return
	}
	ResponseProverb(context, http.StatusOK, results)
}
