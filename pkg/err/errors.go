package errDictionary

import "fmt"

type CodeErr struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (c CodeErr) Error() string {
	return fmt.Sprintf("Code: %d, Detail: %s", c.Code, c.Detail)
}

type CodeErrs []CodeErr

/*
 */
