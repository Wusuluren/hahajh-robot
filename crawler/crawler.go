package crawler

import (
	"hahajh-robot/crawler/qiubai"
)

type Crawler interface {
	Download(url string) ([]map[string]string, error)
}

const (
	QiubaiId = iota
)

func NewCrawler(id int) Crawler {
	var crawler Crawler
	switch id {
	case QiubaiId:
		crawler = qiubai.New()
	}
	return crawler
}
