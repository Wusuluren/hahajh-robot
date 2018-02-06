package main

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"hahajh-robot/crawler"
	//"database/sql"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	//"github.com/garyburd/redigo/redis"
	"context"
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

var sess *mgo.Session
var collect *mgo.Collection

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
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logrus.Fatal(err)
	}
	defer sess.Close()
	collect = sess.DB("hahajh-robot").C("qiubai")

	var ctxCancel context.CancelFunc
	ctx, ctxCancel = context.WithCancel(context.Background())

	go asyncSaveItem()

	go func() {
		qiubai := crawler.NewCrawler(crawler.QiubaiId)
		pageNum := 1
		firstContent := ""
		guardContent := ""
		sleepTime := time.Second * 3
		for {
			url := fmt.Sprintf("https://www.qiushibaike.com/8hr/page/%d/", pageNum)
			logrus.Info(url)
			pageNum += 1
			if pageNum > 13 {
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
						ctxCancel()
						return
					} else {
						sleepTime = time.Second * 3
					}
				}
				qbItem := &qiubaiItem{}
				qbItem.Content = item["content"]
				qbItem.Thumb = item["thumb"]
				ImgUrl := strings.Trim(item["thumb"], "\"")
				filename := ImgUrl[strings.LastIndex(ImgUrl, "/")+1:]
				qbItem.ImgUrl = "http:" + ImgUrl
				qbItem.Filepath = "./pictures/qiushibaike/" + filename
				saveItemChan <- qbItem
				logrus.Info(qbItem)
			}
			time.Sleep(sleepTime)
		}
	}()

	<-mainChan
}

func asyncSaveItem() {
	for {
		select {
		case item := <-saveItemChan:
			if item.Thumb != "" {
				downloadPicture(item.ImgUrl, item.Filepath)
			}
			bytes, err := json.Marshal(*item)
			if err != nil {
				logrus.Error(err)
				continue
			}
			logrus.Println(string(bytes))

			err = collect.Insert(item)
			if err != nil {
				logrus.Error(err)
			}
		case <-ctx.Done():
			mainChan <- true
			return
		}
	}
}

func downloadPicture(url, Filepath string) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()
	f, err := os.OpenFile(Filepath, os.O_CREATE|os.O_WRONLY, 0666)
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
