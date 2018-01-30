package restr

import (
	"fmt"
	"testing"
)

func TestReStrCmp(t *testing.T) {
	fmt.Println(ReStrCmp("abcd", "ab*"))
	fmt.Println(ReStrCmp("abcdef", "ab*f"))
	fmt.Println(ReStrCmp("abcdef", "*ef"))
}
