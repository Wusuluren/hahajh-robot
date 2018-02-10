package crawler_test

import (
	"github.com/wusuluren/hahajh-robot/crawler"
	"os"
	"testing"
)

func TestBsbdjDownload(t *testing.T) {
	env := os.Getenv("ENVIRONMENT")
	if env != "DEBUG" {
		t.Skip("skipped")
	}

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
