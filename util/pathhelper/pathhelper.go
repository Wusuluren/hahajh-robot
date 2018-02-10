package pathhelper

import (
	"fmt"
	"os"
)

type PathHelper struct {
	rootPath    string
	module      string
	pathPattern string
}

func (ph *PathHelper) MakeFilePath(file string) string {
	return fmt.Sprintf("%s/%s", ph.pathPattern, file)
}

func (ph *PathHelper) GetRootPath() string {
	return ph.rootPath
}

func NewPathHelper(module string) (*PathHelper, error) {
	helper := &PathHelper{
		rootPath: "../test-files",
		module:   module,
	}
	helper.pathPattern = fmt.Sprintf("%s/%s", helper.rootPath, helper.module)
	fileInfo, err := os.Stat(helper.pathPattern)
	if err != nil || !fileInfo.IsDir() {
		err := os.MkdirAll(helper.pathPattern, 0777)
		if err != nil {
			return nil, err
		}
	}
	return helper, nil
}
