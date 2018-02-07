package storage_test

import (
	"hahajh-robot/storage"
	"testing"
)

func TestFileStorage(t *testing.T) {
	storage := storage.NewStorage(storage.FileId)
	config := map[string]string{
		"filepath": "file_test.log",
	}
	err := storage.Open(config)
	if err != nil {
		t.Fatal(err)
	}
	defer storage.Close()
	testData1 := []interface{}{1, 2.0, "3", '4'}
	storage.Save(testData1)
	testData2 := map[string]string{
		"name": "hahajh",
	}
	storage.Save(testData2)
}
