package crwaler

import (
	"hahajh-robot/crwaler/qiubai"
)

type Crawler interface {
	Download() error
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
