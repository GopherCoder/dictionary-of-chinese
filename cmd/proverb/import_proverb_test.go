package proverb

import (
	"fmt"
	"testing"
)

func TestAttainTotalPages(test *testing.T) {
	fmt.Println(attainTotalPages())
	fmt.Println(proverbParams.TotalPage)
}

func TestFetchDataRootPage(test *testing.T) {
	fetchDataRootPage()
}
