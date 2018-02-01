package qiubai

import (
	"testing"
)

func TestDownload(t *testing.T) {
	crawler := &Qiubai{}
	url := "https://www.qiushibaike.com/"
	items, err := crawler.Download(url)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		t.Log(item)
	}
}
