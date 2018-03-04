package crawler

import (
	"github.com/wusuluren/gquery"
	"io/ioutil"
	"net/http"
	"strings"
)

type bsbdj struct {
}

func (b *bsbdj) Download(url string) ([]map[string]string, error) {
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
	headerNode := gquery.NewHtml(html).Gquery("body").Eq(0).
		First("div.j-content").
		First("div.g-bd.f-cb").
		First("div.g-mn").
		First("div.j-r-c")
	items := make([]map[string]string, 0)
	for ctr := 0; ctr < 2; ctr++ {
		articles := headerNode.Children("div.j-r-list").Eq(ctr).
			First("ul").
			Children("li")
		for _, article := range articles {
			item := make(map[string]string)
			context := article.First("div.j-r-list-c").
				First("div.j-r-list-c-desc").
				First("a")
			text := context.First("*").Text()
			text = strings.TrimLeft(text, " \t\r\n")
			text = strings.TrimRight(text, " \t\r\n")
			item["content"] = text

			thumb := article.First("div.j-r-list-c").
				First("div.j-r-list-c-img").
				First("a").
				First("img")
			thumbStr := ""
			thumbStr = thumb.Attr("data-original")
			item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

			items = append(items, item)
		}
	}
	for ctr := 0; ctr < 2; ctr++ {
		articles := headerNode.Children("div.j-r-wrst").Eq(ctr).
			First("div.j-list").
			Children("div.j-list-c")
		for _, article := range articles {
			item := make(map[string]string)
			context := article.First("div.j-item").
				First("div.j-item-des").
				First("a")
			text := context.First("*").Text()
			text = strings.TrimLeft(text, " \t\r\n")
			text = strings.TrimRight(text, " \t\r\n")
			item["content"] = text

			thumb := article.First("div.j-item").
				First("a").
				First("div.j-item-img").
				First("img")

			thumbStr := ""
			thumbStr = thumb.Attr("data-original")
			item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

			items = append(items, item)
		}
	}
	return items, nil
}

func newBsbdj() *bsbdj {
	return &bsbdj{}
}
