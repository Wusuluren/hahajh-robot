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
	err = storage.Save(testData2)
	if err != nil {
		t.Fatal(err)
	}

	testData1_1 := []interface{}{}
	err = storage.Next(&testData1_1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(testData1_1)
	testData2_1 := map[string]string{}
	err = storage.Next(&testData2_1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(testData2_1)
}
