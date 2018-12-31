package word

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"

	"github.com/gin-gonic/gin"
)

var wordGlobalParams struct {
	wordIDs   string
	wordHash  string
	wordZsort string
}

func init() {
	wordGlobalParams.wordIDs = "word:ids"
	wordGlobalParams.wordHash = "word:hash:"
	wordGlobalParams.wordZsort = "word:zsort"
}

func ResponseWord(context *gin.Context, code int, value interface{}) {
	context.JSON(
		code, gin.H{
			"data": value,
		},
	)
}

func hashFormat(id string) string {
	return fmt.Sprintf(wordGlobalParams.wordHash+"%s", id)
}

func hashGetAllWord(id string) *model.Word {
	var hashId string
	if strings.Contains(id, "word") {
		hashId = id
	} else {
		hashId = hashFormat(id)

	}

	result, err := redis.StringMap(db.DB.Do("HGETALL", hashId))
	if err != nil {
		return nil
	}
	var one model.Word
	resultId, _ := strconv.Atoi(result["ID"])
	one.ID = resultId
	if value, ok := result["Value"]; ok {
		valueList := strings.Split(value, "~")
		one.Name, one.Explain = valueList[0], valueList[1]
	}
	return &one

}

func randSampleNumber() string {
	result, _ := redis.Int(db.DB.Do("DBSIZE"))
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(result)
	if ok, _ := redis.Bool(db.DB.Do("EXISTS", hashFormat(strconv.Itoa(randNumber)))); !ok {
		return "-1"
	}
	return hashFormat(strconv.Itoa(randNumber))
}

func zaddSort(result *model.Word) bool {
	if number, _ := redis.Int(db.DB.Do("ZINCRBY", wordGlobalParams.wordZsort, 1, result.ID)); number == 0 {
		return false
	}
	return true

}
