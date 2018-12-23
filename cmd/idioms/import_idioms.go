package idioms

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"dictionary-of-chinese/pkg/helper"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"

	"github.com/PuerkitoBio/goquery"
)

var idiomsParams struct {
	RootURL   string
	FirstURL  string
	SecondURL string
	FormatURL string
	TotalPage int
	IdiomHash string
	IdiomIDs  string
}

func init() {
	idiomsParams = struct {
		RootURL   string
		FirstURL  string
		SecondURL string
		FormatURL string
		TotalPage int
		IdiomHash string
		IdiomIDs  string
	}{
		RootURL:   "http://www.zd9999.com",
		FirstURL:  "http://www.zd9999.com/cy/index.htm",
		SecondURL: "http://www.zd9999.com/cy/index_2.htm",
		FormatURL: "http://www.zd9999.com/cy/index_%d.htm",
		TotalPage: 0,
		IdiomHash: "idiom:hash",
		IdiomIDs:  "idiom:ids",
	}
}

func Start() {
	db.Start()
	fetchDataFirstPage()
	fetchDataPerPage()
	defer db.DB.Close()
}

func urlFormat(page int) string {
	return fmt.Sprintf(idiomsParams.FormatURL, page)
}

func urlConsist(part string) string {
	return idiomsParams.RootURL + part
}

func attainTotalPage() (int, string) {
	// body > div:nth-child(2) > center > table > tbody > tr > td:nth-child(1) > a:nth-child(5)
	response, err := helper.DownloadHtml(idiomsParams.FirstURL)
	responseString := string([]byte(response))
	if err != nil {
		return 0, responseString
	}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(responseString))
	totalPageString, _ := doc.Find("body > div:nth-child(2) > center > table > tbody > tr > td:nth-child(1) > a:nth-child(5)").Attr("href")
	totalPageInt, _ := strconv.Atoi(helper.RegexHandler(totalPageString))
	return totalPageInt, responseString
}

func fetchDataFirstPage() {
	total, response := attainTotalPage()
	idiomsParams.TotalPage = total
	fmt.Println(total)
	results := commonFetch(response)
	if !successImportIdiomsHashPage(results) {
		fmt.Println("1234")
		return
	}

}

var g = func(p int, c chan model.Idioms) {
	url := urlFormat(p)
	response, _ := helper.DownloadHtml(url)
	responseString := string([]byte(response))
	result := commonFetch(responseString)
	c <- result
}

func fetchDataPerPage() {
	rand.Seed(time.Now().Unix())
	// 	time.Sleep(5 * time.Second)
	var c = make(chan model.Idioms)
	for p := 2; p <= idiomsParams.TotalPage; p++ {
		go g(p, c)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		results := <-c
		//fmt.Println(results)
		if !successImportIdiomsHashPage(results) {
			fmt.Println(76545678)
			return
		}
	}

}

func commonFetch(response string) model.Idioms {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(response))

	var results model.Idioms
	doc.Find("body > div:nth-child(3) > center > table > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		if i >= 2 {
			selection.Find("td > a").Each(func(i int, selection *goquery.Selection) {
				url, _ := selection.Attr("href")
				idiom := strings.TrimSpace(selection.Text())
				fmt.Println(url, idiom)
				var one model.Idiom
				one.Name = idiom
				fetchExplain(urlConsist(url), &one)
				results = append(results, one)
			})
		}
	})
	return results
}

func fetchExplain(url string, result *model.Idiom) *model.Idiom {
	response, _ := helper.DownloadHtml(url)
	responseString := string([]byte(response))
	//fmt.Println(responseString)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(responseString))
	//fmt.Println(doc.Html())
	table := doc.Find("body > div:nth-child(2) > center > table:nth-child(2) > tbody > tr > td:nth-child(1) > table:nth-child(2) > tbody > tr > td > table > tbody > tr:nth-child(2) > td > table > tbody > tr")
	result.PinYin = table.Eq(0).Find("td").Eq(1).Text()
	result.Explain = table.Eq(1).Find("td").Eq(1).Text()
	result.Source = table.Eq(2).Find("td").Eq(1).Text()
	result.Example = table.Eq(3).Find("td").Eq(1).Text()
	return result
}

func successImportIdiomsHash(result *model.Idiom) bool {
	var values struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		PinYin  string `json:"pinyin"`
		Explain string `json:"explain"`
		Source  string `json:"from"`
		Example string `json:"example"`
	}
	values.ID = attainIdiomIds(idiomsParams.IdiomIDs)
	keyID := fmt.Sprintf(idiomsParams.IdiomHash+":%d", values.ID)
	fmt.Println(keyID)
	if isKeyExist(keyID) {
		return true
	}
	values.Name = result.Name
	values.PinYin = result.PinYin
	values.Explain = result.Explain
	values.Source = result.Source
	values.Example = result.Example
	fmt.Println(values)
	if _, err := db.DB.Do("HMSET", redis.Args{}.Add(keyID).AddFlat(&values)...); err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

func successImportIdiomsHashPage(results model.Idioms) bool {
	for _, result := range results {
		if !successImportIdiomsHash(&result) {
			return false
		}
	}
	return true
}

func isKeyExist(key string) (ok bool) {
	ok, _ = redis.Bool(db.DB.Do("EXISTS", key))
	return
}

func attainIdiomIds(key string) int {
	var c int
	if !isKeyExist(key) {
		c, _ = redis.Int(db.DB.Do("SET", key, 0))
		return c
	} else {
		c, _ = redis.Int(db.DB.Do("INCR", key))
	}
	return c
}
