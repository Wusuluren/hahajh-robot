package crawler

import (
	"fmt"
	"hahajh-robot/util/gquery"
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
		if node.Label == "body" {
			htmlRoot = node
			break
		}
	}
	if isEmptyNode(htmlRoot) {
		fmt.Println("htmlRoot not found")
		return nil, nil
	}
	items := make([]map[string]string, 0)

	articles := htmlRoot.Find("div.\"j-content\"").
		Find("div.\"g-bd f-cb\"").
		Find("div.\"g-mn\"").
		Find("div.\"j-r-c\"").
		Eq("div.\"j-r-list\"", 0).
		Find("ul").
		Children("li")
	for _, article := range articles {
		item := make(map[string]string)
		context := article.Find("div.\"j-r-list-c\"").
			Find("div.\"j-r-list-c-desc\"").
			Find("a")
		item["content"] = getChildrenText(context)

		thumb := article.Find("div.\"j-r-list-c\"").
			Find("div.\"j-r-list-c-img\"").
			Find("a").
			Find("img")
		thumbStr := ""
		if value, ok := thumb.Attribute["data-original"]; ok {
			thumbStr = value
		}
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
	}

	articles = htmlRoot.Find("div.\"j-content\"").
		Find("div.\"g-bd f-cb\"").
		Find("div.\"g-mn\"").
		Find("div.\"j-r-c\"").
		Eq("div.\"j-r-list\"", 1).
		Find("ul").
		Children("li")
	for _, article := range articles {
		item := make(map[string]string)
		context := article.Find("div.\"j-r-list-c\"").
			Find("div.\"j-r-list-c-desc\"").
			Find("a")
		item["content"] = getChildrenText(context)

		thumb := article.Find("div.\"j-r-list-c\"").
			Find("div.\"j-r-list-c-img\"").
			Find("a").
			Find("img")
		thumbStr := ""
		if value, ok := thumb.Attribute["data-original"]; ok {
			thumbStr = value
		}
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
	}

	articles = htmlRoot.Find("div.\"j-content\"").
		Find("div.\"g-bd f-cb\"").
		Find("div.\"g-mn\"").
		Find("div.\"j-r-c\"").
		Eq("div.\"j-r-wrst gud-put*\"", 0).
		Find("div.\"j-list\"").
		Children("div.\"j-list-c\"")

	for _, article := range articles {
		item := make(map[string]string)
		context := article.Find("div.\"j-item\"").
			Find("div.\"j-item-des\"").
			Find("a")
		item["content"] = getChildrenText(context)

		thumb := article.Find("div.\"j-item\"").
			Find("a").
			Find("div.\"j-item-img\"").
			Find("img")

		thumbStr := ""
		if value, ok := thumb.Attribute["data-original"]; ok {
			thumbStr = value
		}
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
	}

	articles = htmlRoot.Find("div.\"j-content\"").
		Find("div.\"g-bd f-cb\"").
		Find("div.\"g-mn\"").
		Find("div.\"j-r-c\"").
		Eq("div.\"j-r-wrst gud-put*\"", 1).
		Find("div.\"j-list\"").
		Children("div.\"j-list-c\"")

	for _, article := range articles {
		item := make(map[string]string)
		context := article.Find("div.\"j-item\"").
			Find("div.\"j-item-des\"").
			Find("a")
		item["content"] = getChildrenText(context)

		thumb := article.Find("div.\"j-item\"").
			Find("a").
			Find("div.\"j-item-img\"").
			Find("img")

		thumbStr := ""
		if value, ok := thumb.Attribute["data-original"]; ok {
			thumbStr = value
		}
		item["thumb"] = strings.Trim(thumbStr, "\t\n\r ")

		items = append(items, item)
	}

	return items, nil
}

func newBsbdj() *bsbdj {
	return &bsbdj{}
}
