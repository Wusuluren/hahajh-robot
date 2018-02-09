package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"hahajh-robot/crawler"
	//"database/sql"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	//"github.com/garyburd/redigo/redis"
	"context"
	"hahajh-robot/storage"
	"sync"
)

type qiubaiItem struct {
	Content  string
	Thumb    string
	ImgUrl   string
	Filepath string
}

var saveItemChan = make(chan *qiubaiItem, 1024)
var mainChan = make(chan bool)

var ctx context.Context
var wg = sync.WaitGroup{}

var strg = storage.NewStorage(storage.FileId)

func main() {
	//db, err := sql.Open("mysql", "root:root@/qiubai")
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//defer db.Close()
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//defer conn.Close()

	config := map[string]string{
		"filepath": "qiubai.log",
	}
	if err := strg.Open(config); err != nil {
		logrus.Fatal(err)
	}
	defer strg.Close()

	qiubaiUrls := []string{
		"https://www.qiushibaike.com/8hr/page/%d/",
		"https://www.qiushibaike.com/hot/page/%d/",
		"https://www.qiushibaike.com/imgrank/page/%d/",
		"https://www.qiushibaike.com/text/page/%d/",
		"https://www.qiushibaike.com/history/page/%d/",
		"https://www.qiushibaike.com/pic/page/%d/",
		"https://www.qiushibaike.com/textnew/page/%d/",
	}
	var ctxCancel context.CancelFunc
	ctx, ctxCancel = context.WithCancel(context.Background())

	go asyncSaveItem()
	for _, url := range qiubaiUrls {
		go asyncCrawlerPages(url)
		wg.Add(1)
	}

	wg.Wait()
	ctxCancel()
	<-mainChan
}

func asyncCrawlerPages(urlPattern string) {
	qiubai := crawler.NewCrawler(crawler.QiubaiId)
	pageNum := 1
	firstContent := ""
	guardContent := ""
	sleepTime := time.Second * 3
	for {
		url := fmt.Sprintf(urlPattern, pageNum)
		logrus.Info(url)
		pageNum += 1
		if pageNum > 13 {
			wg.Done()
			return
			guardContent = firstContent
			pageNum = 1
		}

		items, err := qiubai.Download(url)
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, item := range items {
			if firstContent == "" {
				firstContent = item["content"]
			} else {
				if item["content"] == guardContent {
					guardContent = firstContent
					pageNum = 1
					firstContent = ""
					sleepTime = time.Minute
					wg.Done()
					return
				} else {
					sleepTime = time.Second * 3
				}
			}
			qbItem := &qiubaiItem{}
			qbItem.Content = item["content"]
			qbItem.Thumb = item["thumb"]
			if item["thumb"] != "" {
				ImgUrl := strings.Trim(item["thumb"], "\"")
				filename := ImgUrl[strings.LastIndex(ImgUrl, "/")+1:]
				qbItem.ImgUrl = "http:" + ImgUrl
				qbItem.Filepath = "./pictures/qiushibaike/" + filename
			}
			saveItemChan <- qbItem
			//logrus.Info(qbItem)
		}
		time.Sleep(sleepTime)
	}
}

func asyncSaveItem() {
	var err error
	itemsInterface := make([]interface{}, 64)
	ctr := 0
	for {
		select {
		case item := <-saveItemChan:
			if item.Thumb != "" {
				downloadPicture(item.ImgUrl, item.Filepath)
			}
			//bytes, err := json.Marshal(*item)
			//if err != nil {
			//	logrus.Error(err)
			//	continue
			//}
			//logrus.Println(string(bytes))

			itemsInterface = append(itemsInterface, *item)
			ctr += 1
			if ctr >= 64 {
				err = strg.Save(itemsInterface...)
				if err != nil {
					logrus.Error(err)
				}
				itemsInterface = itemsInterface[0:0]
				ctr = 0
			}
		case <-ctx.Done():
			if ctr > 0 {
				err = strg.Save(itemsInterface...)
				if err != nil {
					logrus.Error(err)
				}
			}
			mainChan <- true
			return
		}
	}
}

func downloadPicture(url, filepath string) {
	_, err := os.Stat(filepath)
	if err == nil {
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	_, err = f.Write(bytes)
	if err != nil {
		logrus.Error(err)
		return
	}
}
