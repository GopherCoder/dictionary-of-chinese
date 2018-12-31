package idiom

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

var idiomGlobalParam struct {
	key   string
	ids   string
	zsort string
}

func init() {
	idiomGlobalParam.key = "idiom:hash"
	idiomGlobalParam.ids = "idiom:ids"
	idiomGlobalParam.zsort = "idiom:zsort"
}

func ResponseIdiom(context *gin.Context, code int, response interface{}) {
	context.JSON(
		code, gin.H{"data": response})

}

func consistKey(id string) string {
	return fmt.Sprintf(idiomGlobalParam.key+":%s", id)
}

func isKeyExistInRedis(key string) bool {
	if exist, _ := redis.Bool(db.DB.Do("EXISTS", key)); !exist {
		return !exist
	}
	return true
}

func hashGetAllByKey(key string) *model.Idiom {
	result, err := redis.StringMap(db.DB.Do("HGETALL", key))
	if err != nil {
		fmt.Println("execute command hgetall fail")
		return nil
	}
	var one model.Idiom
	one.ID = result["ID"]
	one.Name = result["Name"]
	one.PinYin = result["PinYin"]
	one.Explain = result["Explain"]
	one.Source = result["Source"]
	one.Example = result["Example"]
	return &one
}

func numberHashGetAll(number int) model.Idioms {
	rand.Seed(time.Now().UnixNano())
	if number == 0 {
		number = 1
	}
	if !isKeyExistInRedis(idiomGlobalParam.ids) {
		return nil
	}
	totalNumber, err := redis.Int(db.DB.Do("GET", idiomGlobalParam.ids))
	if err != nil {
		return nil
	}
	var results model.Idioms
	for i := 0; i < number; i++ {
		key := consistKey(strconv.Itoa(rand.Intn(totalNumber)))
		//fmt.Println(key)
		one := hashGetAllByKey(key)
		results = append(results, *one)
	}
	return results

}

func totalNumber() int {
	if !isKeyExistInRedis(idiomGlobalParam.ids) {
		return -1
	}
	totalNumber, err := redis.Int(db.DB.Do("GET", idiomGlobalParam.ids))
	if err != nil {
		return -1
	}
	return totalNumber
}

func zaddOneRecord(values *model.Idiom) bool {
	if ok, _ := redis.Int(db.DB.Do("ZINCRBY", idiomGlobalParam.zsort, 1, values.ID)); ok == 0 {
		return false
	}
	return true
}
