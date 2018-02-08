package robot_test

import (
	"hahajh-robot/robot"
	"testing"
)

func TestParseUrl(t *testing.T) {
	config, err := robot.ParseUrl("hahajh-url.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
}

func TestParseAccount(t *testing.T) {
	config, err := robot.ParseAccount("hahajh-account.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
	for _, item := range config {
		t.Log(item)
	}
}
