package gquery_test

import (
	"fmt"
	"hahajh-robot/util/gquery"
	"testing"
)

func TestReStrCmp(t *testing.T) {
	fmt.Println(gquery.ReStrCmp("abcd", "ab*"))
	fmt.Println(gquery.ReStrCmp("abcdef", "ab*f"))
	fmt.Println(gquery.ReStrCmp("abcdef", "*ef"))
}
