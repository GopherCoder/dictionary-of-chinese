package helper

import (
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func DownloadHtml(url string, args ...interface{}) ([]byte, error) {
	request := gorequest.New()
	resp, _, _ := request.Get(url).
		Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36").
		End()
	if len(args) == 0 {
		utf8Content := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)

	} else {
		return ioutil.ReadAll(resp.Body)
		//utf8Content := transform.NewReader(resp.Body, simplifiedchinese.HZGB2312.NewDecoder())
		//return ioutil.ReadAll(utf8Content)
	}
}
