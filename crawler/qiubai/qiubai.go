package qiubai

import (
	"fmt"
	"hahajh-robot/util/gquery"
	"io/ioutil"
	"net/http"
	"strings"
)

type Qiubai struct {
}

func (q *Qiubai) Download(url string) ([]map[string]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//f, err := os.Open("test.html")
	//if err != nil {
	//	return nil, err
	//}
	//defer f.Close()
	//bytes, err := ioutil.ReadAll(f)
	//f, err := os.OpenFile("test.html", os.O_CREATE, 0666)
	//if err != nil {
	//	return err
	//}
	//defer f.Close()
	//f.Write(bytes)
	bytes, err := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	//print(html)

	htmlNodeTree := gquery.ParseHtml(html)
	var htmlRoot *gquery.HtmlNode
	for _, node := range htmlNodeTree {
		if node.Label == "html" {
			htmlRoot = node
			break
		}
	}
	if htmlRoot == nil {
		fmt.Println("htmlRoot not found")
		return nil, nil
	}
	articles := htmlRoot.Find("body").
		Find("div.#\"content\"").
		Find("div.\"content-block clearfix\"").
		Find("div.#\"content-left\"").
		Children("div.\"article block untagged mb15 typs_*\"")

		//fmt.Println(len(articles))
	//items := make([]*qiubaiItem, 0)
	items := make([]map[string]string, 0)
	for _, article := range articles {
		item := make(map[string]string)
		context := article.Find("a.\"contentHerf\"").
			Find("div.\"content\"").
			Find("span")
		textArry := make([]string, 0)
		for _, node := range context.Children("") {
			if node.Label == "" {
				textArry = append(textArry, node.Text)
			}
		}
		item["content"] = strings.Trim(strings.Join(textArry, ""), "\t\n\r ")

		thumb := article.Find("div.\"thumb\"").
			Find("a").
			Find("img")
		thumbStr := ""
		if value, ok := thumb.Attribute["src"]; ok {
			thumbStr = value
		}
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
		//fmt.Println(item)
		//fmt.Println("-------------------------")
	}

	return items, nil
}

func New() *Qiubai {
	q := Qiubai{}
	return &q
}
