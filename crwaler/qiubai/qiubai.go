package qiubai

import (
	//"net/http"
	"fmt"
	"hahajh-robot/util/gquery"
	"io/ioutil"
	"os"
)

type Qiubai struct {
}

type qiubaiItem struct {
	content string
	thumb   string
}

func (q *Qiubai) Download() error {
	//resp, err := http.Get("https://www.qiushibaike.com/")
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()
	f, err := os.Open("test.html")
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	//f, err := os.OpenFile("test.html", os.O_CREATE, 0666)
	//if err != nil {
	//	return err
	//}
	//defer f.Close()
	//f.Write(bytes)
	html := string(bytes)
	//print(html)
	//htmltree.ParseHtml(html)
	//items := parseHtml(html)
	//println(len(items))

	htmlNodeTree := gquery.ParseHtml(html)
	var htmlRoot *gquery.HtmlNode
	for _, node := range htmlNodeTree {
		if node.Label == "html" {
			htmlRoot = node
			break
		}
	}
	if htmlRoot == nil {
		return nil
	}
	a := htmlRoot.Find("body")
	if a != nil {
		fmt.Println(a)
	}
	a = a.Find("div.#\"content\"")
	if a != nil {
		fmt.Println(a)
	}
	a = a.Find("div.\"content-block clearfix\"")
	if a != nil {
		fmt.Println(a)
	}
	a = a.Find("div.#\"content-left\"")
	if a != nil {
		fmt.Println(a)
	}
	a = a.Find("div.\"article block untagged mb15 typs_long\"")
	if a != nil {
		fmt.Println(a)
	}
	a = a.Find("div.\"content\"")
	if a != nil {
		fmt.Println(a)
	}
	a = a.Find("span")
	if a != nil {
		fmt.Println(a)
	}
	for _, node := range a.Children("") {
		if node.Label == "" {
			fmt.Println(node.Text)
		}
	}
	return nil
}

func New() *Qiubai {
	q := Qiubai{}
	return &q
}
