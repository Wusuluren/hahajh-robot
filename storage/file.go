package storage

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type fileStorage struct {
	file   *os.File
	reader *bufio.Reader
	writer *bufio.Writer
	mutex  *sync.Mutex
}

func (fs *fileStorage) Open(config map[string]string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	var err error
	if _, ok := config["filepath"]; !ok {
		return configError
	}
	filepath := config["filepath"]
	fs.file, err = os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	fs.reader = bufio.NewReader(fs.file)
	fs.writer = bufio.NewWriter(fs.file)
	return nil
}

func (fs *fileStorage) Close() error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	var err error
	if fs.file != nil {
		err = fs.writer.Flush()
		err = fs.file.Close()
	}
	return err
}

func (fs *fileStorage) Save(items ...interface{}) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	for _, item := range items {
		bytes, err := json.Marshal(item)
		if err != nil {
			return err
		}
		_, err = fs.writer.WriteString(string(bytes) + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *fileStorage) Next(items ...interface{}) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	for _, item := range items {
		line, _, err := fs.reader.ReadLine()
		if err != nil {
			return err
		}
		err = json.Unmarshal(line, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func newFileStorage() Storage {
	return &fileStorage{
		mutex: new(sync.Mutex),
	}
}
