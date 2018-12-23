package helper

import (
	"regexp"
	"strings"
)

func StringHandler(value string) string {
	//                        共有 281 页
	valueList := strings.Split(strings.TrimSpace(value), " ")
	return valueList[1]
}

func RegexHandler(value string) string {
	// index_198.htm
	regexPattern := `\d{1,}`
	re, _ := regexp.Compile(regexPattern)
	return re.FindString(value)
}
