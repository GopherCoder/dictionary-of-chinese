package proverb

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/garyburd/redigo/redis"
)

var proverbGlobalParams struct {
	Ids          string
	ProverbHash  string
	ProverbZsort string
}

func init() {
	proverbGlobalParams.Ids = "proverb:ids"
	proverbGlobalParams.ProverbHash = "proverb:hash"
	proverbGlobalParams.ProverbZsort = "proverb:zsort"
}

func isSuitableId(id int) bool {
	if !isKeyExist(proverbGlobalParams.Ids) {
		return false
	}
	result, err := redis.Int(db.DB.Do("GET", proverbGlobalParams.Ids))
	if err != nil {
		return false
	}
	if id < 0 || id > result {
		return false
	}
	return true

}

func isKeyExist(key string) bool {
	ok, err := redis.Bool(db.DB.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return ok
}

func totalNumber() int {
	result, _ := redis.Int(db.DB.Do("GET", proverbGlobalParams.Ids))
	return result
}

func ResponseProverb(context *gin.Context, code int, response interface{}) {
	context.JSON(
		code, gin.H{
			"data": response,
		})
}

func consistKey(id string) string {
	return fmt.Sprintf(proverbGlobalParams.ProverbHash+":%s", id)
}

func hashGetAllProverbByKey(key string) *model.Proverb {
	if !isKeyExist(key) {
		return nil
	}
	result, err := redis.StringMap(db.DB.Do("HGETALL", key))
	if err != nil {
		return nil
	}
	var one model.Proverb
	one.ID = result["ID"]
	one.Riddle = result["Riddle"]
	one.Answer = result["Answer"]
	return &one
}

func hashGetAllProverbAtSamples(number int) model.Proverbs {
	total := totalNumber()
	rand.Seed(time.Now().UnixNano())
	if number < 0 || number > total {
		return nil
	}
	if number == 0 {
		number = 1
	}
	var results model.Proverbs
	for i := 0; i < number; i++ {
		id := rand.Intn(total)
		key := consistKey(strconv.Itoa(id))
		one := hashGetAllProverbByKey(key)
		results = append(results, *one)
	}
	return results
}

func zaddSort(result *model.Proverb) bool {
	if number, _ := redis.Int(db.DB.Do("ZINCRBY", proverbGlobalParams.ProverbZsort, 1, result.ID)); number == 0 {
		return false
	}
	return true
}
