package idiom

import (
	"dictionary-of-chinese/pkg/db"
	"fmt"
	"testing"
)

func TestHashGetAllByKey(tests *testing.T) {
	db.Start()
	defer db.DB.Close()
	//idiom:hash:12629
	//idiom:hash:25522
	//idiom:hash:22547
	tt := []struct {
		key string
	}{
		{
			key: "idiom:hash:12629",
		},
		{
			key: "idiom:hash:25522",
		},
		{
			key: "idiom:hash:22547",
		},
	}
	for _, t := range tt {
		fmt.Println(hashGetAllByKey(t.key))
	}
}

func TestNumberHashGetAll(tests *testing.T) {
	db.Start()
	defer db.DB.Close()
	tt := []struct {
		number int
	}{
		{
			number: 10,
		},
		{
			number: 5,
		},
	}
	for _, t := range tt {
		fmt.Println(numberHashGetAll(t.number))
	}
}
