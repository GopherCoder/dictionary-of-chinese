package proverb

import (
	"dictionary-of-chinese/model"
	"dictionary-of-chinese/pkg/db"
	"dictionary-of-chinese/pkg/err"
	"dictionary-of-chinese/pkg/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"

	"github.com/PuerkitoBio/goquery"
)

var proverbParams struct {
	RootUrl     string
	Second      string
	Format      string
	TotalPage   int
	ProverbIDs  string
	ProverbHash string
}

func init() {
	proverbParams.RootUrl = "http://xhy.5156edu.com/html2/xhy.html"
	proverbParams.Second = "http://xhy.5156edu.com/html2/xhy_2.html"
	proverbParams.Format = "http://xhy.5156edu.com/html2/xhy_%s.html"
	proverbParams.ProverbIDs = "proverb:ids"
	proverbParams.ProverbHash = "proverb:hash"
}

func Start() {
	db.Start()
	var results = make(chan model.Proverbs)
	go func() {
		results <- fetchDataRootPage()
	}()
	successImportHashPage(<-results)

	var resultss = make(chan model.Proverbs)
	for p := 2; p < proverbParams.TotalPage; p++ {
		fmt.Println(p, "page")
		go func(p int) {
			resultss <- fetchDataPerPage(p)
		}(p)
		successImportHashPage(<-resultss)
	}
	defer db.DB.Close()
}

// 格式化
func urlFormat(page int) string {
	toString := strconv.Itoa(page)
	return fmt.Sprintf(proverbParams.Format, toString)
}

// 获取所有页码
func attainTotalPages() string {
	response, err := helper.DownloadHtml(proverbParams.RootUrl)
	if err != nil {
		err := errDictionary.CodeErr{
			Code:   401,
			Detail: "can not get reponse",
		}
		fmt.Println(err)
		return "-1"
	}
	responseString := string([]byte(response))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(responseString))
	text := doc.Find("body > div:nth-child(3) > center > table:nth-child(1) > tbody > tr > td:nth-child(1)").Text()
	textInt, _ := strconv.Atoi(helper.StringHandler(text))
	proverbParams.TotalPage = textInt
	return responseString
}

// 导入数据: page one
func fetchDataRootPage() model.Proverbs {
	response := attainTotalPages()
	results := commonFetch(response)
	return results

	//fmt.Println(results)
	// proverb:ids : string , incr
	// proverb:hash HMSET
}

// page: per page except one
func fetchDataPerPage(p int) model.Proverbs {
	url := urlFormat(p)
	response, _ := helper.DownloadHtml(url)
	responseString := string([]byte(response))
	results := commonFetch(responseString)
	return results
}

// parse html data
func commonFetch(response string) model.Proverbs {

	var result model.Proverbs
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(response))
	//body > div:nth-child(3) > center > table:nth-child(2) > tbody > tr:nth-child(2)
	doc.Find("body > div:nth-child(3) > center > table:nth-child(2) > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			riddle := selection.Find("td").Eq(0)
			answer := selection.Find("td").Eq(1)
			var one model.Proverb
			one.Riddle = strings.TrimSpace(riddle.Text())
			one.Answer = strings.TrimSpace(answer.Text())
			fmt.Println(one)
			result = append(result, one)
		}
	})
	return result
}

// proverb:ids
func count() (int, error) {
	var (
		c   int
		err error
	)
	if ok, _ := redis.Bool(db.DB.Do("EXISTS", proverbParams.ProverbIDs)); !ok {
		c, err = redis.Int(db.DB.Do("SET", proverbParams.ProverbIDs, 0))
	} else {
		c, err = redis.Int(db.DB.Do("INCR", proverbParams.ProverbIDs))
	}
	return c, err
}

// proverb:hash HMSET by one model.Proverb
func successImportHash(id int, result *model.Proverb) bool {
	var hash struct {
		ID     string `json:"id"`
		Riddle string `json:"riddle"`
		Answer string `json:"answer"`
	}
	result.ID = strconv.Itoa(id)
	if ok, _ := redis.Bool(db.DB.Do("HEXISTS", proverbParams.ProverbHash, id)); ok {
		return true
	}
	fmt.Println(result)
	hash.ID = result.ID
	hash.Riddle = result.Riddle
	hash.Answer = result.Answer
	if _, err := db.DB.Do("HMSET", redis.Args{}.Add(fmt.Sprintf(proverbParams.ProverbHash+":%d", id)).AddFlat(&hash)...); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// proverb:hash HMEST by page
func successImportHashPage(results model.Proverbs) {
	for _, result := range results {
		id, _ := count()
		if !successImportHash(id, &result) {
			continue
		}
	}
}
