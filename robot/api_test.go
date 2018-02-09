package robot_test

import (
	"hahajh-robot/robot"
	"testing"
)

func TestApi(t *testing.T) {
	configUrls, err := robot.ParseUrl("hahajh-url.yml")
	configAccount, err := robot.ParseAccount("hahajh-account.yml")
	if err != nil {
		t.Fatal(err)
	}
	account := configAccount[0]
	t.Log(account.Username, account.Password)

	item := &robot.HahajhItem{
		Text:    "test",
		Picture: "test.jpg",
	}
	err = robot.InitAccount(configUrls, account)
	if err != nil {
		t.Fatal(err)
	}
	//err = account.Signup()
	//if err != nil {
	//	t.Fatal(err)
	//}
	err = account.Login()
	if err != nil {
		t.Fatal(err)
	}
	defer account.Logout()
	err = account.Publish(item)
	if err != nil {
		t.Fatal(err)
	}
}
