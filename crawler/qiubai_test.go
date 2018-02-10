package crawler_test

import (
	"github.com/wusuluren/hahajh-robot/crawler"
	"testing"
)

func TestQiubaiDownload(t *testing.T) {
	crawler := crawler.NewCrawler(crawler.QiubaiId)
	url := "https://www.qiushibaike.com/"
	items, err := crawler.Download(url)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		t.Log(item)
	}
}
