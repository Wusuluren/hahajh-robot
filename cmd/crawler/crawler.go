package main

import (
	"fmt"
	"hahajh-robot/crwaler"
)

func main() {
	qiubai := crwaler.NewCrawler(crwaler.QiubaiId)
	items, err := qiubai.Download()
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
