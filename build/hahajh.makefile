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
