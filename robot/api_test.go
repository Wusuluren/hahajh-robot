package robot_test

import (
	"hahajh-robot/robot"
	"testing"
)

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestApi(t *testing.T) {
	configUrls, err := robot.ParseUrl("hahajh-url.yml")
	configAccount, err := robot.ParseAccount("hahajh-account.yml")
	checkError(t, err)
	account := configAccount[0]
	t.Log(account.Username, account.Password)
	account.Urls = configUrls

	item := robot.HahajhItem{
		Text:    "test",
		Picture: "test.jpg",
	}
	err = robot.InitAccount(account)
	checkError(t, err)
	err = account.Login()
	checkError(t, err)
	defer account.Logout()
	err = account.Publish(item)
	checkError(t, err)
}
