hahajh_EXE = hahajh.exe
LIB = ../robot/*.go ../storage/*
hahajh_SRC = ../cmd/robot/hahajh.go

all:$(hahajh_EXE)

$(hahajh_EXE):$(hahajh_SRC) $(LIB)
	go build $(hahajh_SRC)
