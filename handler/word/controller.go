package word

import (
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
	ResponseWord(context, http.StatusOK, hashGetAllWord(params).BasicSerialize())

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
