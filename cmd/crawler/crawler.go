package main

import (
	"fmt"
	"hahajh-robot/crwaler"
)

func main() {
	qiubai := crwaler.NewCrawler(crwaler.QiubaiId)
	err := qiubai.Download()
	if err != nil {
		fmt.Println(err)
	}
}
