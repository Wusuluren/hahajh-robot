package robot_test

import (
	"hahajh-robot/robot"
	"hahajh-robot/util/pathhelper"
	"os"
	"testing"
	"time"
)

func TestApi(t *testing.T) {
	env := os.Getenv("ENVIRONMENT")
	if env != "DEBUG" {
		t.Skip("skipped")
	}

	ph, err := pathhelper.NewPathHelper("robot")
	if err != nil {
		t.Fatal(err)
	}
	configUrls, err := robot.ParseUrl(ph.MakeFilePath("test-url.yml"))
	if err != nil {
		t.Fatal(err)
	}
	configAccounts, err := robot.ParseAccount(ph.MakeFilePath("test-account.yml"))
	if err != nil {
		t.Fatal(err)
	}
	for _, account := range configAccounts {
		t.Log(account.Username, account.Password)

		item := &robot.HahajhItem{
			Text:    "test",
			Picture: ph.MakeFilePath("test.jpg"),
		}
		_ = item
		err = robot.InitAccount(configUrls, account)
		if err != nil {
			t.Fatal(err)
		}
		err = account.Login()
		if err != nil {
			err = account.Signup()
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			err = account.Login()
			if err != nil {
				t.Fatal(err)
			}
		}
		time.Sleep(time.Second)
		//err = account.Publish(item)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//time.Sleep(time.Second)
		err = account.Logout()
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
