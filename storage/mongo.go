package storage

import (
	"gopkg.in/mgo.v2"
)

type mongoStorage struct {
	sess    *mgo.Session
	collect *mgo.Collection
}

func (ms *mongoStorage) Open(config map[string]string) error {
	var err error
	if _, ok := config["address"]; !ok {
		return configError
	}
	addr := config["address"]
	if _, ok := config["database"]; !ok {
		return configError
	}
	database := config["database"]
	if _, ok := config["collection"]; !ok {
		return configError
	}
	collection := config["collection"]
	ms.sess, err = mgo.Dial(addr)
	if err != nil {
		return err
	}
	ms.collect = ms.sess.DB(database).C(collection)
	return nil
}

func (ms *mongoStorage) Close() {
	if ms.sess != nil {
		ms.sess.Close()
	}
}

func (ms *mongoStorage) Save(items ...interface{}) error {
	return ms.collect.Insert(items...)
}

func newMongoStorage() Storage {
	return &mongoStorage{}
}
