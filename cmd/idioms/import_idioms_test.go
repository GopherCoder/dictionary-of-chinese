package idioms

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"fmt"
	"testing"
)

func TestFetchExplain(tests *testing.T) {
	tt := []struct {
		url    string
		result model.Idiom
	}{
		{
			url:    "http://www.zd9999.com/cy/htm0/1.htm",
			result: model.Idiom{ID: "1"},
		},
		{
			url:    "http://www.zd9999.com/cy/htm0/2.htm",
			result: model.Idiom{ID: "2"},
		},
	}
	for _, t := range tt {
		fmt.Println(fetchExplain(t.url, &t.result))
	}
}

func TestFetchDataFirstPage(tests *testing.T) {
	fetchDataFirstPage()
}
func TestFetchDataPerPage(tests *testing.T) {
	fetchDataPerPage()
}

func TestAttainIdiomIds(tests *testing.T) {
	db.Start()
	tt := []struct {
		value string
	}{
		{
			value: "idioms:ids",
		},
	}

	for _, t := range tt {
		fmt.Println(attainIdiomIds(t.value))
	}
}
