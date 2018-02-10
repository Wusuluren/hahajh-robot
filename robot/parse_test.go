package robot_test

import (
	"hahajh-robot/robot"
	"hahajh-robot/util/pathhelper"
	"os"
	"testing"
)

func TestParseUrl(t *testing.T) {
	env := os.Getenv("ENVIRONMENT")
	if env != "DEBUG" {
		t.Skip("skipped")
	}

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
	env := os.Getenv("ENVIRONMENT")
	if env != "DEBUG" {
		t.Skip("skipped")
	}

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
