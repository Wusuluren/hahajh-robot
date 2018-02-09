package main

import (
	"github.com/Sirupsen/logrus"
	"hahajh-robot/robot"
	"hahajh-robot/storage"
	"sync"
	"time"
)

type qiubaiItem struct {
	Content  string
	Thumb    string
	ImgUrl   string
	Filepath string
}

var strg storage.Storage
var wg = sync.WaitGroup{}

func main() {
	configUrls, err := robot.ParseUrl("hahajh-url.yml")
	if err != nil {
		logrus.Fatal(err)
	}
	configAccounts, err := robot.ParseAccount("hahajh-account.yml")
	if err != nil {
		logrus.Fatal(err)
	}

	strg = storage.NewStorage(storage.FileId)
	config := map[string]string{
		"filepath": "qiubai.log",
	}
	err = strg.Open(config)
	if err != nil {
		logrus.Fatal(err)
	}
	defer strg.Close()

	err = robot.InitAccount(configUrls, configAccounts...)
	if err != nil {
		logrus.Fatal(err)
	}
	for _, account := range configAccounts {
		wg.Add(1)
		go robotApi(account)
	}
	wg.Wait()
}

func robotApi(account *robot.Account) {
	var err error
	err = account.Login()
	if err != nil {
		logrus.Error(err)
		err = account.Signup()
		if err != nil {
			logrus.Fatal(err)
		}
		time.Sleep(time.Second)
		err = account.Login()
		if err != nil {
			logrus.Fatal(err)
		}
	}
	defer account.Logout()
	logrus.Info(account.Username)

	item := &qiubaiItem{}
	hahajhItem := &robot.HahajhItem{}
	for {
		err = strg.Next(item)
		if err != nil {
			logrus.Error(err)
			wg.Done()
			return
		}
		logrus.Info(item)
		hahajhItem.Text = item.Content
		if len(item.Thumb) > 0 {
			continue //for test, don't publish image
			hahajhItem.Picture = item.Filepath
		} else {
			hahajhItem.Picture = ""
		}
		err = account.Publish(hahajhItem)
		if err != nil {
			logrus.Error(err)
			time.Sleep(time.Second * 5)
		}
		time.Sleep(time.Second)
	}
}
