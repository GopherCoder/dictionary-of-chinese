package words

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChildrenResponse(tests *testing.T) {
	Convey("get explain for one word", tests, func() {
		fmt.Println(childrenResponse("http://www.zd9999.com/ci/htm31/312801.htm"))
		fmt.Println(childrenResponse("http://www.zd9999.com/ci/htm31/312812.htm"))
		fmt.Println(childrenResponse("http://www.zd9999.com/ci/htm31/312813.htm"))
	})
}

func TestFetchFetchTotalPage(tests *testing.T) {
	Convey("get total page and first page", tests, func() {
		fmt.Println(fetchTotalPage())
		fmt.Println(wordGlobalParams.totalPage)
	})
}

func TestFetchFetchPerPage(tests *testing.T) {
	Convey("fetch every page ", tests, func() {
		fmt.Println(fetchPerPage(2))
		fmt.Println(wordGlobalParams.totalPage)
	})
}

func TestDivNumber(tests *testing.T) {
	Convey("div 5000", tests, func() {
		fmt.Println(divNumber(1))     // 0
		fmt.Println(divNumber(100))   // 0
		fmt.Println(divNumber(5001))  // 1
		fmt.Println(divNumber(10003)) // 2
		fmt.Println(divNumber(4999))  // 0
	})
}
