package bsbdj

import (
	"testing"
)

func TestDownload(t *testing.T) {
	crawler := &Bsbdj{}
	url := "http://www.budejie.com/pic/"
	items, err := crawler.Download(url)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		t.Log(item)
	}
}
