qiubai_EXE = qiubai.exe
LIB = ../crawler/*.go ../util/*
qiubai_SRC = ../cmd/crawler/qiubai.go

all:$(qiubai_EXE)

$(qiubai_EXE):$(qiubai_SRC) $(LIB)
	go build $(qiubai_SRC)
