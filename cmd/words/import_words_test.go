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
