package helper

import (
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func DownloadHtml(url string) ([]byte, error) {
	request := gorequest.New()
	resp, _, _ := request.Get(url).
		Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36").
		End()
	utf8Content := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	return ioutil.ReadAll(utf8Content)
}
