package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/wusuluren/hahajh-robot/util/stack"
	"os"
	"strings"
)

const (
	Keyword = iota
	Name
	Args
	Values
)

type ArgInfo struct {
	name  string
	isDir bool
}

type DefineItem struct {
	keyword string
	name    string
	args    []string
	values  []string
	prefix  string
}

var (
	defineMap = make(map[string]*DefineItem)
)

func usage() {
	fmt.Println("usage: gomacro [directory | file]")
	os.Exit(1)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) <= 1 {
		usage()
	}
	args := os.Args[1:]
	argsInfo := make([]*ArgInfo, 0)
	for _, arg := range args {
		//fmt.Println(arg)
		fileInfo, err := os.Stat(arg)
		Panic(err)
		argsInfo = append(argsInfo, &ArgInfo{
			name:  arg,
			isDir: fileInfo.IsDir(),
		})
	}
	for _, argInfo := range argsInfo {
		if argInfo.isDir {

		} else {
			gomacro(argInfo.name)
		}
	}
}

func readWriteLine(reader *bufio.Reader, writer *bufio.Writer) (string, error) {
	bytes, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	line := string(bytes)
	_, err = writer.WriteString(line + "\n")
	Panic(err)
	return line, nil
}

func preReadBytes(reader *bufio.Reader, num int) ([]byte, error) {
	//lines := make([]string, num)
	bytes, err := reader.Peek(num)
	//for i := 0;i < num;i++ {
	//	bytes, _, err := reader.ReadLine()
	//	if err != nil {
	//		return lines, err
	//	}
	//	line := string(bytes)
	//	lines[i] = line
	//}
	return bytes, err
}

func parseDeclare(lines []string) ([]*DefineItem, error) {
	defineItems := make([]*DefineItem, 0)
	lineNum := len(lines)
	lineIdx := 0
	var defineItem *DefineItem
	except := Keyword
	var keyword, name string
	var args = make([]string, 0)
	var values = make([]string, 0)
	stk := stack.NewStack()
	for {
		line := lines[lineIdx]
		//fmt.Println(line)
		//line = strings.Trim(line, "\t ")
		if except == Keyword {
			if len(line) >= 8 && line[0:8] == "#define " {
				if defineItem != nil {
					defineItems = append(defineItems, defineItem)
					keyword = ""
					name = ""
					values = make([]string, 0)
					args = make([]string, 0)
				}
				keyword = "define"
				idx := 8
				last := 8
				for i := 8; i < len(line); i++ { //TODO:strong check
					c := line[i]
					if c == ' ' {
						except = Values
						idx = i + 1
						break
					} else if c == '(' {
						except = Args
						name = line[8:i]
						stk.Push(i + 1)
					} else if c == ')' {
						except = Values
						last = stk.Pop().(int)
						args = append(args, line[last:i])
					} else if c == ',' {
						except = Args
						last = stk.Pop().(int)
						args = append(args, line[last:i])
						stk.Push(i + 1)
					}
				}
				v := ""
				for i := len(line) - 1; i >= idx; i-- {
					if line[i] != ' ' {
						if line[i] == '\\' {
							except = Values
						} else {
							except = Keyword
							v = line[idx : i+1]
							break
						}
					}
				}
				idx = 0
				prefix := ""
				for i, c := range v {
					if c != ' ' || c != '\t' {
						idx = i
						break
					}
				}
				if v[idx:] != "" {
					v = v[idx:]
					prefix = v[0:idx]
					values = append(values, v)
				}
				//fmt.Println("first:", prefix)
				defineItem = &DefineItem{
					keyword: keyword,
					name:    name,
					args:    args,
					values:  values,
					prefix:  prefix,
				}
			} else if strings.Trim(line, "\t ") != "" {
				return nil, errors.New(fmt.Sprintf("except keyword: %s", line))
			}
		} else if except == Values {
			v := ""
			except = Keyword
			for i := len(line) - 1; i >= 0; i-- {
				c := line[i]
				if c != ' ' && c != '\r' && c != '\n' && c != '\t' {
					if c == '\\' {
						except = Values
					} else {
						v = line[0 : i+1]
						break
					}
				}
			}
			//fmt.Println("v:", v)
			if defineItem.prefix != "" {
				v = strings.Trim(v, defineItem.prefix)
			} else {
				idx := 0
				prefix := ""
				for i, c := range v {
					if c != ' ' || c != '\t' {
						idx = i
						break
					}
				}
				if v[idx:] != "" {
					v = v[idx:]
					prefix = v[0:idx]
					defineItem.values = append(defineItem.values, v)
					defineItem.prefix = prefix
					//fmt.Println("next:", prefix)
				}
			}
		}

		lineIdx += 1
		if lineIdx >= lineNum {
			break
		}
	}
	if defineItem != nil {
		defineItems = append(defineItems, defineItem)
	}
	//for _, item := range defineItems {
	//	fmt.Println(item)
	//}
	return defineItems, nil
}

func parseStat(line string) (*DefineItem, error) {
	var defineItem *DefineItem
	var name string
	var args = make([]string, 0)
	stk := stack.NewStack()
	line = strings.TrimLeft(line, "\t ")
	line = strings.TrimPrefix(line, "// +go macro: ")
	last := 0
	except := Name
	for i := 0; i < len(line); i++ { //TODO:strong check
		c := line[i]
		if except == Values {
			if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
				Panic(errors.New("bad statement"))
			}
		} else {
			if c == '(' {
				except = Args
				name = line[0:i]
				stk.Push(i + 1)
			} else if c == ')' {
				except = Values
				last = stk.Pop().(int)
				args = append(args, line[last:i])
			} else if c == ',' {
				except = Args
				last = stk.Pop().(int)
				args = append(args, line[last:i])
				stk.Push(i + 1)
			}
		}
	}
	defineItem = &DefineItem{
		name: name,
		args: args,
	}
	return defineItem, nil
}

func gomacro(filepath string) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	Panic(err)
	reader := bufio.NewReader(file)
	file2, err := os.OpenFile(filepath+"2", os.O_WRONLY|os.O_CREATE, 0666)
	Panic(err)
	writer := bufio.NewWriter(file2)
	var line string
	for {
		line, err = readWriteLine(reader, writer)
		if err != nil {
			break
		}
		if strings.HasPrefix(line, "// +go macro") {
			// macro define
			line, err = readWriteLine(reader, writer)
			if err != nil {
				break
			}
			if strings.Trim(line, "\t\r\n ") != "/*" {
				Panic(errors.New("parse"))
			}
			macroLines := make([]string, 0, 4)
			for {
				line, err = readWriteLine(reader, writer)
				if err != nil {
					break
				}
				if strings.Trim(line, "\t\r\n ") == "*/" {
					break
				}
				macroLines = append(macroLines, line)
			}
			defineItems, err := parseDeclare(macroLines)
			Panic(err)
			for _, defineItem := range defineItems {
				defineMap[defineItem.name] = defineItem
			}
		} else if strings.HasPrefix(strings.TrimLeft(line, "\t "), "// +go macro: ") {
			// macro declare
			defineStat, err := parseStat(line)
			Panic(err)
			//fmt.Println(defineStat)
			if defineItem, ok := defineMap[defineStat.name]; !ok {
				Panic(errors.New(fmt.Sprintf("unknowm name:%s", defineStat.name)))
			} else {
				size := 0
				for _, v := range defineItem.values {
					size += len(v) + 1
				}
				bytes, err := preReadBytes(reader, size)
				Panic(err)
				if string(bytes) != strings.Join(defineItem.values, "\n")+"\n" {
					for _, v := range defineItem.values {
						_, err = writer.WriteString(v + "\n")
						Panic(err)
					}
				}
			}
		}
	}
	err = file.Close()
	Panic(err)
	err = writer.Flush()
	Panic(err)
	err = file2.Close()
	Panic(err)
	err = os.Remove(filepath)
	Panic(err)
	err = os.Rename(filepath+"2", filepath)
	Panic(err)
}
