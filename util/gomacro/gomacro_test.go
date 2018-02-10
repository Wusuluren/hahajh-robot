package main_test

// +go macro
/*
#define checkError(err) \
	if err != nil { \
		t.Log(err) \
	}
*/

import (
	"testing"
)

func TestGomacro(t *testing.T) {
	var err error
	// +go macro: checkError(err)
}
