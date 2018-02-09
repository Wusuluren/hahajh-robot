package storage

type Storage interface {
	Open(config map[string]string) error
	Save(items ...interface{}) error
	Next(item interface{}) error
	NextN(items ...interface{}) error
	Close() error
}

const (
	FileId = iota
	MongoId
	MysqlId
	RedisId
)

func NewStorage(id int) Storage {
	var storage Storage
	switch id {
	case FileId:
		storage = newFileStorage()
	case MongoId:
		storage = newMongoStorage()
	}
	return storage
}
