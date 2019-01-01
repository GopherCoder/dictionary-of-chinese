package idiom

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/garyburd/redigo/redis"

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
	if !zaddOneRecord(result) {
		ResponseIdiom(context, http.StatusBadRequest, "zincrby fail")
		return
	}
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

func GetIdiomsRankHandler(context *gin.Context) {
	// 根据搜索的 id, 维护一个固定长度的 zset
	// step one : sorted set exists or not
	// step two : sorted set range by score

	if ok, _ := redis.Bool(db.DB.Do("EXISTS", idiomGlobalParam.zsort)); !ok {
		ResponseIdiom(context, http.StatusBadRequest, "not exists rank")
		return
	}
	var rankNumber int
	rankNumber = 10
	result, err := redis.Strings(db.DB.Do("ZREVRANGE", idiomGlobalParam.zsort, 0, rankNumber-1))
	if err != nil {
		ResponseIdiom(context, http.StatusBadRequest, err)
		return
	}
	var results []*model.Idiom
	for _, key := range result {
		values := hashGetAllByKey(fmt.Sprintf(idiomGlobalParam.key+":%s", key))
		results = append(results, values)
	}
	ResponseIdiom(context, http.StatusOK, &results)
}
