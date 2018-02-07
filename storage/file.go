package storage

import (
	"encoding/json"
	"os"
	"strings"
)

type fileStorage struct {
	file *os.File
}

func (fs *fileStorage) Open(config map[string]string) error {
	var err error
	if _, ok := config["filepath"]; !ok {
		return configError
	}
	filepath := config["filepath"]
	fs.file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (fs *fileStorage) Close() {
	if fs.file != nil {
		fs.file.Close()
	}
}

func (fs *fileStorage) Save(items ...interface{}) error {
	itemStr := make([]string, 0, 16)
	for _, item := range items {
		bytes, err := json.Marshal(item)
		if err != nil {
			return err
		}
		itemStr = append(itemStr, string(bytes))
	}
	_, err := fs.file.WriteString(strings.Join(itemStr, "\n") + "\n")
	return err
}

func newFileStorage() Storage {
	return &fileStorage{}
}
