package crawler

import (
	"hahajh-robot/crawler/bsbdj"
	"hahajh-robot/crawler/qiubai"
)

type Crawler interface {
	Download(url string) ([]map[string]string, error)
}

const (
	QiubaiId = iota
	BsbdjId
)

func NewCrawler(id int) Crawler {
	var crawler Crawler
	switch id {
	case QiubaiId:
		crawler = qiubai.New()
	case BsbdjId:
		crawler = bsbdj.New()
	}
	return crawler
}
