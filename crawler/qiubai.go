package crawler

import (
	"fmt"
	"github.com/wusuluren/gquery"
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

	var htmlRoot *gquery.HtmlNode
	children := gquery.NewHtml(html).Gquery("body")
	if len(children) > 0 {
		htmlRoot = children[0]
	} else {
		fmt.Println("htmlRoot not found")
		return nil, nil
	}
	articles := htmlRoot.First("div#content").
		First("div.content-block.clearfix").
		First("div#content-left").
		Children("div.article.block.untagged.mb15")
	items := make([]map[string]string, 0)
	for _, article := range articles {
		item := make(map[string]string)
		context := article.First("a.contentHerf").
			First("div.content").
			First("span")

		text := context.First("*").Text()
		text = strings.TrimLeft(text, " \t\r\n")
		text = strings.TrimRight(text, " \t\r\n")
		item["content"] = text

		thumb := article.First("div.thumb").
			First("a").
			First("img")
		thumbStr := ""
		thumbStr = thumb.Attr("src")
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
	}

	return items, nil
}

func newQiubai() *qiubai {
	return &qiubai{}
}
