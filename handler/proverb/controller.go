package proverb

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/garyburd/redigo/redis"

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
	result := hashGetAllProverbByKey(fullKey)
	if !zaddSort(result) {
		ResponseProverb(context, http.StatusBadRequest, "zsort fail")
		return
	}
	ResponseProverb(context, http.StatusOK, result)

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

func GetRankProverbHandler(context *gin.Context) {
	// step one is key exists or not
	// step two range zsort

	if ok, _ := redis.Bool(db.DB.Do("EXISTS", proverbGlobalParams.ProverbZsort)); !ok {
		ResponseProverb(context, http.StatusBadRequest, "zsort fail")
		return
	}
	var number int
	number = 10
	result, err := redis.Strings(db.DB.Do("ZREVRANGE", proverbGlobalParams.ProverbZsort, 0, number-1, "WITHSCORES"))
	if err != nil {
		ResponseProverb(context, http.StatusBadRequest, "zsort fail")
		return
	}
	fmt.Println(result, "kjhgfdfghjk")
	var results []*model.Proverb
	for _, key := range result {
		value := hashGetAllProverbByKey(fmt.Sprintf(proverbGlobalParams.ProverbHash+":%s", key))
		results = append(results, value)
	}
	ResponseProverb(context, http.StatusOK, results)
}
