EXE =../cmd/crawler/crawler.exe
LIB = ../crawler/*.go ../util/*
SRC = ../cmd/crawler/crawler.go

all:$(EXE)

$(EXE):$(SRC) $(LIB)
	go build $(SRC)
