package word

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func GetWordsByNameHandler(context *gin.Context) {}

func GetWordsByIdHandler(context *gin.Context) {
	var params string
	params = context.Param("id")
	paramsInt, _ := strconv.Atoi(strings.TrimSpace(params))

	numberTotal, _ := redis.Int(db.DB.Do("get", wordGlobalParams.wordIDs))
	if numberTotal != 0 {
		if paramsInt > numberTotal {
			ResponseWord(context, http.StatusBadRequest, fmt.Sprintf("id should in (0 ~ %d)", numberTotal))
			return

		}
	}
	result := hashGetAllWord(params)
	if !zaddSort(result) {
		ResponseWord(context, http.StatusBadRequest, "zsort fail")
		return
	}
	ResponseWord(context, http.StatusOK, result.BasicSerialize())

}

func GetWordsAtRandomHandler(context *gin.Context) {

	number := randSampleNumber()
	fmt.Println(number)
	if number == "-1" {
		ResponseWord(context, http.StatusBadRequest, fmt.Sprintf("id should be exists"))
		return
	}
	ResponseWord(context, http.StatusOK, hashGetAllWord(number).BasicSerialize())

}

func GetRankHandler(context *gin.Context) {
	// step key exists
	// step zsort
	if ok, _ := redis.Bool(db.DB.Do("EXISTS", wordGlobalParams.wordZsort)); !ok {
		ResponseWord(context, http.StatusBadRequest, "zsort fail")
		return
	}

	var number int
	number = 10
	result, err := redis.Strings(db.DB.Do("ZREVRANGE", wordGlobalParams.wordZsort, 0, number-1))
	if err != nil {
		ResponseWord(context, http.StatusBadRequest, "zsort fail")
		return
	}
	var results []*model.Word
	for _, key := range result {
		results = append(results, hashGetAllWord(key))
	}
	ResponseWord(context, http.StatusOK, results)

}
