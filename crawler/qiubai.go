package crawler

import (
	"fmt"
	"hahajh-robot/util/gquery"
	"io/ioutil"
	"net/http"
	"strings"
)

type qiubai struct {
}

func (q *qiubai) Download(url string) ([]map[string]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	html := string(bytes)

	htmlNodeTree := gquery.ParseHtml(html)
	var htmlRoot *gquery.HtmlNode
	for _, node := range htmlNodeTree {
		if node.Label == "html" {
			htmlRoot = node
			break
		}
	}
	if isEmptyNode(htmlRoot) {
		fmt.Println("htmlRoot not found")
		return nil, nil
	}
	articles := htmlRoot.Find("body").
		Find("div.#\"content\"").
		Find("div.\"content-block clearfix\"").
		Find("div.#\"content-left\"").
		Children("div.\"article block untagged mb15 typs_*\"")

	items := make([]map[string]string, 0)
	for _, article := range articles {
		item := make(map[string]string)
		context := article.Find("a.\"contentHerf\"").
			Find("div.\"content\"").
			Find("span")
		item["content"] = getChildrenText(context)

		thumb := article.Find("div.\"thumb\"").
			Find("a").
			Find("img")
		thumbStr := ""
		if value, ok := thumb.Attribute["src"]; ok {
			thumbStr = value
		}
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
	}

	return items, nil
}

func newQiubai() *qiubai {
	return &qiubai{}
}
