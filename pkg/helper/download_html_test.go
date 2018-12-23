package helper

import (
	"fmt"
	"testing"
)

func TestDownloadHtml(t *testing.T) {
	tt := "http://xhy.5156edu.com/html2/xhy.html"
	r, _ := DownloadHtml(tt)
	fmt.Println(string([]byte(r)))
}
