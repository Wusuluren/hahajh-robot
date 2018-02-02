package crawler_test

import (
	"hahajh-robot/crawler"
	"testing"
)

func TestBsbdjDownload(t *testing.T) {
	crawler := crawler.NewCrawler(crawler.BsbdjId)
	url := "http://www.budejie.com/pic/"
	items, err := crawler.Download(url)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		t.Log(item)
	}
}
