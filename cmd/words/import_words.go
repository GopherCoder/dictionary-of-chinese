package words

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/helper"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Start() {}

var wordGlobalParams struct {
	rootURL   string
	firstURL  string
	secondURL string
	formatURL string
	totalPage int
}

func init() {
	wordGlobalParams.rootURL = "http://www.zd9999.com"
	wordGlobalParams.firstURL = "http://www.zd9999.com/ci/index.htm"
	wordGlobalParams.secondURL = "http://www.zd9999.com/ci/index_2.htm"
	wordGlobalParams.formatURL = "http://www.zd9999.com/ci/index_%d.htm"
}

func urlFormat(value string) string {
	return wordGlobalParams.rootURL + value
}

func fetchTotalPage() (model.Words, error) {
	var ok bool
	ok = true
	response, err := helper.DownloadHtml(wordGlobalParams.firstURL, ok)
	if err != nil {
		return nil, err
	}
	responseString := string(response)
	//fmt.Println(responseString)
	return commonHandler(responseString), nil

}

func fetchPerPage(page int) (model.Words, error) {
	var ok bool
	ok = true
	url := fmt.Sprintf(wordGlobalParams.formatURL, page)
	response, _ := helper.DownloadHtml(url, ok)
	responseString := string([]byte(response))
	return commonHandler(responseString), nil
}

func commonHandler(response string) model.Words {
	responseString := strings.NewReader(response)
	doc, _ := goquery.NewDocumentFromReader(responseString)
	var results model.Words

	//body > div:nth-child(2) > center > table > tbody > tr > td:nth-child(1) > a:nth-child(5)
	if wordGlobalParams.totalPage == 0 {
		endPage := doc.Find("body > div:nth-child(2) > center > table > tbody > tr > td:nth-child(1) > a:nth-child(3)")
		fmt.Println(endPage.Text())
		//wordGlobalParams.totalPage, _ = strconv.Atoi(endPage)
	}
	// body > div:nth-child(3) > center > table > tbody > tr:nth-child(2)
	doc.Find("body > div:nth-child(3) > center > table > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(i)
		if i > 0 {
			children := selection.Find("td > a")
			childrenUrl, _ := children.Attr("href")
			childrenContent := strings.TrimSpace(children.Text())
			childrenExplain := childrenResponse(childrenUrl)
			var one model.Word
			one.Name = childrenContent
			one.Explain = childrenExplain
			results = append(results, one)
		}
	})
	return results
}

func childrenResponse(url string) string {
	var ok bool
	ok = true
	response, _ := helper.DownloadHtml(url, ok)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string([]byte(response))))
	explainText := doc.Find("body > div:nth-child(3) > center > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(2) > td").Text()
	newReplacer := strings.NewReplacer(" ", "", "\n", "", "\t", "")
	newExplainText := newReplacer.Replace(explainText)
	return newExplainText
}
