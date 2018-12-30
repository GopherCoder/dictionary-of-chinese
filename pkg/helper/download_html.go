package helper

import (
	"bufio"
	"io"
	"io/ioutil"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
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
		e := determineCharset(resp.Body)
		utf8Content := transform.NewReader(resp.Body, e.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}
}

func determineCharset(i io.Reader) encoding.Encoding {
	resp, err := bufio.NewReader(i).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(resp, "")
	//fmt.Println(e)
	return e
}
