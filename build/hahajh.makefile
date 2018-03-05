ROOT_PATH = github.com/wusuluren/hahajh-robot

qiubai_EXE = qiubai.exe
qiubai_SRC = ../cmd/crawler/qiubai.go

hahajh_EXE = hahajh.exe
hahajh_SRC = ../cmd/robot/hahajh.go

LIB = ../crawler/*.go ../util/* ../robot/*.go ../storage/*


all:$(qiubai_EXE) $(hahajh_EXE)

$(qiubai_EXE):$(qiubai_SRC) $(LIB)
	go build $(qiubai_SRC)

$(hahajh_EXE):$(hahajh_SRC) $(LIB)
	go build $(hahajh_SRC)

build:
	go build $(qiubai_SRC)
	go build $(hahajh_SRC)

clean:
	test -f $(qiubai_EXE) && rm $(qiubai_EXE)
	test -f $(hahajh_EXE) && rm $(hahajh_EXE)

test:
	go test -v  $(ROOT_PATH)/crawler
	go test -v  $(ROOT_PATH)/robot
	go test -v  $(ROOT_PATH)/storage
	go test -v  $(ROOT_PATH)/util/stack
