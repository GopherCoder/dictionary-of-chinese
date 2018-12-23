package helper

import "strings"

func StringHandler(value string) string {
	//                        共有 281 页
	valueList := strings.Split(strings.TrimSpace(value), " ")
	return valueList[1]
}
