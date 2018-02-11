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

type GoStat struct {
	stat    string
	vars    []string
	pos     []int
	pattern string
}

type DefineItem struct {
	keyword string
	name    string
	args    []string
	values  []*GoStat
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

func parseMacroDeclare(lines []string) ([]*DefineItem, error) {
	defineItems := make([]*DefineItem, 0)
	lineNum := len(lines)
	lineIdx := 0
	var defineItem *DefineItem
	except := Keyword
	var keyword, name string
	var args = make([]string, 0)
	var values = make([]*GoStat, 0)
	for {
		line := lines[lineIdx]
		if except == Keyword {
			if len(line) >= 8 && line[0:8] == "#define " {
				//fmt.Println(line)
				if defineItem != nil {
					defineItems = append(defineItems, defineItem)
					defineItem = nil
					keyword = ""
					name = ""
					values = make([]*GoStat, 0)
					args = make([]string, 0)
				}
				keyword = "define"
				isStart := false
				stk := stack.NewStack()
				start := 8
				except = Name
				for i := 8; i < len(line); i++ { //TODO:strong check
					c := line[i]
					if isIdentifier(c) {
						if !isStart {
							stk.Push(i)
							isStart = true
						}
					} else {
						//if c == '(' {
						//} else if c == ',' {
						//} else if c == ')' {
						//	except = Values
						//	idx = i+1
						//	break
						//}
						if c == '\\' {
							except = Values
							break
						}
						if !stk.Empty() {
							start = stk.Pop().(int)
						}
						if isStart {
							isStart = false
							identifier := line[start:i]
							//fmt.Println(identifier)
							if except == Name {
								except = Args
								name = identifier
							} else if except == Args {
								args = append(args, identifier)
								if c == ')' { //TODO: check other
									except = Values
								}
							} else if except == Values {
								goStat := &GoStat{
									stat: identifier,
									vars: make([]string, 0),
									pos:  make([]int, 0),
								}
								values = append(values, goStat)
							}
						}
					}
				}
				prefix := ""
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
			//fmt.Println(line)
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
			stat := ""
			if defineItem.prefix != "" {
				stat = strings.Trim(v, defineItem.prefix)
			} else {
				idx := 0
				for i, c := range v {
					if c != ' ' || c != '\t' {
						idx = i
						break
					}
				}
				stat = v[idx:]
				defineItem.prefix = v[0:idx]
			}
			if stat != "" {
				goStat := &GoStat{
					stat: stat,
					vars: make([]string, 0),
					pos:  make([]int, 0),
				}
				defineItem.values = append(defineItem.values, goStat)
			}
			//fmt.Println(defineItem.prefix, stat)
		}
		lineIdx += 1
		if lineIdx >= lineNum {
			break
		}
	}
	if defineItem != nil {
		defineItems = append(defineItems, defineItem)
	}
	parseGoStats(defineItems)
	for _, item := range defineItems {
		fmt.Println(item)
		for _, value := range item.values {
			fmt.Println(value.stat, value.pattern, value.vars, value.pos)
		}
	}
	return defineItems, nil
}

func isIdentifier(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func parseGoStats(defineItems []*DefineItem) {
	for _, defineItem := range defineItems {
		if len(defineItem.args) == 0 {
			continue
		}
		if len(defineItem.values) == 0 {
			continue
		}
		for _, value := range defineItem.values {
			//fmt.Println(value.stat)
			stk := stack.NewStack()
			idx := 0
			isStart := false
			start := 0
			for {
				c := value.stat[idx]
				if isIdentifier(c) {
					if !isStart {
						stk.Push(idx)
						isStart = true
					}
				} else {
					if !stk.Empty() {
						start = stk.Pop().(int)
					}
					if isStart {
						isStart = false
						identifier := value.stat[start:idx]
						isArg := false
						for i, arg := range defineItem.args {
							if arg == identifier {
								value.vars = append(value.vars, identifier)
								value.pos = append(value.pos, i)
								isArg = true
							}
						}
						if isArg {
							value.pattern += "%s"
						} else {
							value.pattern += identifier
						}
						//fmt.Println(isArg, start, identifier, value.pattern, string(c))
					}
					value.pattern += string(c)
				}
				idx += 1
				if idx >= len(value.stat) {
					if !stk.Empty() {
						start = stk.Pop().(int)
					}
					if isStart {
						isStart = false
						identifier := value.stat[start:idx]
						isArg := false
						for i, arg := range defineItem.args {
							if arg == identifier {
								value.vars = append(value.vars, identifier)
								value.pos = append(value.pos, i)
								isArg = true
							}
						}
						if isArg {
							value.pattern += "%s"
						} else {
							value.pattern += identifier
						}
						//fmt.Println(isArg, start, identifier, value.pattern, string(c))
					}
					break
				}
			}
		}
		//fmt.Println(defineItem)
	}
}

func parseMacroStat(line string) (*DefineItem, error) {
	var defineItem *DefineItem
	var name string
	var args = make([]string, 0)
	line = strings.TrimLeft(line, "\t ")
	line = strings.TrimPrefix(line, "// +go macro: ")
	stk := stack.NewStack()
	start := 0
	except := Name
	isStart := false
	for i := 0; i < len(line); i++ { //TODO:strong check
		c := line[i]
		if isIdentifier(c) {
			if !isStart {
				stk.Push(i)
				isStart = true
			}
		} else {
			if !stk.Empty() {
				start = stk.Pop().(int)
			}
			if isStart {
				isStart = false
				identifier := line[start:i]
				//fmt.Println(identifier)
				if except == Name {
					except = Args
					name = identifier
				} else if except == Args {
					args = append(args, identifier)
					if c == ')' { //TODO: check other
						except = Values
					}
				}
			}
		}
	}
	defineItem = &DefineItem{
		name: name,
		args: args,
		//values: values,
	}
	parseGoStats([]*DefineItem{defineItem})
	//fmt.Println(defineItem)
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
	eof := false
	for !eof {
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
			defineItems, err := parseMacroDeclare(macroLines)
			Panic(err)
			for _, defineItem := range defineItems {
				defineMap[defineItem.name] = defineItem
			}
		} else if strings.HasPrefix(strings.TrimLeft(line, "\t "), "// +go macro: ") {
			// macro declare
			defineStat, err := parseMacroStat(line)
			Panic(err)
			if defineItem, ok := defineMap[defineStat.name]; !ok {
				Panic(errors.New(fmt.Sprintf("unknowm name:%s", defineStat.name)))
			} else {
				size := 0
				for _, stat := range defineItem.values {
					size += len(stat.stat) + 1
				}
				bytes, _ := reader.Peek(size) //ignore io.EOF
				statValues := string(bytes)
				macroValues := ""
				for _, value := range defineItem.values {
					if len(defineStat.args) > 0 {
						ifs := make([]interface{}, 0)
						for i := range value.pos {
							arg := defineStat.args[value.pos[i]]
							ifs = append(ifs, arg)
						}
						macroValues += fmt.Sprintf(value.pattern+"\n", ifs...)
					}
				}
				//fmt.Println(len(macroValues), macroValues)
				//fmt.Println(len(statValues), statValues)
				if statValues != macroValues {
					for _, stat := range defineItem.values {
						if len(defineStat.args) > 0 {
							ifs := make([]interface{}, 0)
							for i := range stat.pos {
								arg := defineStat.args[stat.pos[i]]
								ifs = append(ifs, arg)
							}
							_, err = writer.WriteString(fmt.Sprintf(stat.pattern+"\n", ifs...))
						} else {
							_, err = writer.WriteString(stat.stat + "\n")
						}
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
