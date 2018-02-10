package robot_test

import (
	"hahajh-robot/robot"
	"hahajh-robot/util/pathhelper"
	"testing"
)

func TestParseUrl(t *testing.T) {
	ph, err := pathhelper.NewPathHelper("robot")
	if err != nil {
		t.Fatal(err)
	}
	config, err := robot.ParseUrl(ph.MakeFilePath("test-url.yml"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
}

func TestParseAccount(t *testing.T) {
	ph, err := pathhelper.NewPathHelper("robot")
	if err != nil {
		t.Fatal(err)
	}
	config, err := robot.ParseAccount(ph.MakeFilePath("test-account.yml"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
	for _, item := range config {
		t.Log(item)
	}
}
