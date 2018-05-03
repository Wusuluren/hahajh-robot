package crawler

import (
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
	if err != nil {
		return nil, err
	}
	html := string(bytes)
	articles := gquery.NewHtml(html).Gquery("body").Eq(0).
		First("div#content").
		First("div.content-block.clearfix").
		First("div#content-left").
		Children("div.article.block.untagged.mb15")
	items := make([]map[string]string, 0)
	for _, article := range articles {
		item := make(map[string]string)
		context := article.First("a.contentHerf").
			First("div.content").
			First("span")
		if context.Failed() { //for some weird reason
			context = article.First("a.'contentHerf'").
				First("div.content").
				First("span")
		}
		text := ""
		children := context.Children("*")
		if len(children) > 0 {
			for _, node := range children {
				text += node.Text()
			}
		} else {
			text = context.Text()
		}
		text = strings.TrimLeft(text, " \t\r\n")
		text = strings.TrimRight(text, " \t\r\n")
		item["content"] = text

		thumb := article.First("div.thumb").
			First("a").
			First("img")
		thumbStr := thumb.Attr("src")
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
	}

	return items, nil
}

func newQiubai() *qiubai {
	return &qiubai{}
}
