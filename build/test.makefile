ROOT_PATH = github.com/wusuluren/hahajh-robot

all:
	go test -v  $(ROOT_PATH)/crawler
	go test -v  $(ROOT_PATH)/robot
	go test -v  $(ROOT_PATH)/storage
	go test -v  $(ROOT_PATH)/util/gquery
	go test -v  $(ROOT_PATH)/util/stack
