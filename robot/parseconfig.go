package robot

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func parse(filepath string, config interface{}) (interface{}, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}
	return config, err
}

func ParseUrl(filepath string) (map[string]string, error) {
	config := make(map[string]string)
	configIt, err := parse(filepath, config)
	return configIt.(map[string]string), err
}

func ParseAccount(filepath string) ([]*Account, error) {
	config := make([]*Account, 0)
	configIt, err := parse(filepath, &config)
	return *(configIt.(*[]*Account)), err
}
