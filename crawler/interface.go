package crawler

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
		crawler = newQiubai()
	case BsbdjId:
		crawler = newBsbdj()
	}
	return crawler
}
