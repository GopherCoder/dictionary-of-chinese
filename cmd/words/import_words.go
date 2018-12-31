package words

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"dictionary-of-chinese/pkg/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"

	"github.com/PuerkitoBio/goquery"
)

func Start() {
	db.Start()
	fmt.Println(redis.Bool(db.DB.Do("EXISTS", "proverb:ids")))
	var results = make(chan model.Words)
	go func() {
		result, err := fetchTotalPage()
		if err != nil {
			return
		}
		results <- result

	}()
	importOneWordHash(results)
	// todo: goroutine

	var resultsAll = make(chan model.Words)
	for p := 2; p <= wordGlobalParams.totalPage; p++ {
		fmt.Println("Page:", p)
		go func(p int) {
			results, err := fetchPerPage(p)
			if err != nil {
				return
			}
			resultsAll <- results
		}(p)
		importOneWordHash(resultsAll)
	}

	if _, err := importNumber(wordGlobalParams.wordCount); err != nil {
		return
	}
}

var wordGlobalParams struct {
	rootURL   string
	firstURL  string
	secondURL string
	formatURL string
	totalPage int
	wordIds   string
	wordHash  string
	wordCount int
}

func init() {
	wordGlobalParams.rootURL = "http://www.zd9999.com"
	wordGlobalParams.firstURL = "http://www.zd9999.com/ci/index.htm"
	wordGlobalParams.secondURL = "http://www.zd9999.com/ci/index_2.htm"
	wordGlobalParams.formatURL = "http://www.zd9999.com/ci/index_%d.htm"

	wordGlobalParams.wordIds = "word:ids"
	wordGlobalParams.wordHash = "word:hash:"
	wordGlobalParams.wordCount = 0
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
	result := commonHandler(responseString)
	return result, nil

}

func fetchPerPage(page int) (model.Words, error) {
	var ok bool
	ok = true
	url := fmt.Sprintf(wordGlobalParams.formatURL, page)
	response, _ := helper.DownloadHtml(url, ok)
	responseString := string([]byte(response))
	words := commonHandler(responseString)
	return words, nil
}

// fixme : fixed
func commonHandler(response string) model.Words {
	responseString := strings.NewReader(response)
	doc, _ := goquery.NewDocumentFromReader(responseString)
	var results model.Words
	if wordGlobalParams.totalPage == 0 {
		endPage, _ := doc.Find("body a").Eq(2).Attr("href")
		total := helper.RegexHandler(endPage)
		wordGlobalParams.totalPage, _ = strconv.Atoi(helper.RegexHandler(total))
	}
	doc.Find("div").Eq(0).Find("table > tbody > tr td").Each(func(i int, selection *goquery.Selection) {
		children := selection.Find("a")
		childrenUrl, _ := children.Attr("href")
		childrenContent := strings.TrimSpace(children.Text())
		childrenExplain := childrenResponse(urlFormat(childrenUrl))
		var one model.Word
		one.ID = wordGlobalParams.wordCount
		one.Name = childrenContent
		one.Explain = childrenExplain
		results = append(results, one)
		wordGlobalParams.wordCount += 1
		fmt.Println(one)
	})
	return results
}

func childrenResponse(url string) string {
	var ok bool
	ok = true
	response, _ := helper.DownloadHtml(url, ok)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string([]byte(response))))
	//fmt.Println(doc.Html())
	explainText := doc.Find("div").Eq(0).Find("tr").Eq(0).
		Find("td tbody tr td").Eq(1).Text()
	newReplacer := strings.NewReplacer(" ", "", "\n", "", "\t", "")
	newExplainText := newReplacer.Replace(explainText)
	return newExplainText
}

// operate data into redis

/*
step one:
- hash
- max length of hash is 5000
hash:word:number key: id, value: name~explain
- words:count string count all data of words

*/

func importOneWordHash(words chan model.Words) bool {
	var hashValue struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	}

	for _, word := range <-words {
		hashValue.ID = word.ID
		hashValue.Value = fmt.Sprintf(word.Name + "~" + word.Explain)
		fmt.Println(hashValue, fmt.Sprintf(wordGlobalParams.wordHash+"%d", word.ID))
		if response, err := db.DB.Do("HMSET", redis.Args{}.Add(fmt.Sprintf(wordGlobalParams.wordHash+"%d", word.ID)).AddFlat(&hashValue)...); err == nil {
			fmt.Println("result:", response)
			//return false
		} else {
			return false
		}
	}
	return true
}

func divNumber(number int) int {
	if number == 0 {
		return 0
	}
	return number / 5000
}

func importNumber(number int) (bool, error) {
	return redis.Bool(db.DB.Do("SET", wordGlobalParams.wordIds, number))
}
