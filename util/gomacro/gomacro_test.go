package main_test

// +go macro
/*
#define checkError(err) \
	if err != nil { \
		t.Log(err) \
	}
#define max(a, b, c) \
	if a > b { \
		c = a \
	} else { \
		c = b \
	}
*/

import (
	"testing"
)

func TestGomacro(t *testing.T) {
	var err error
	_ = err
	// +go macro: checkError(err)
	if err != nil {
		t.Log(err)
	}
	if err != nil {
		t.Log(err)
	}
	var c int
	_ = c
	// +go macro: max(1, 2, c)
	if 1 > 2 {
		c = 1
	} else {
		c = 2
	}
	if 1 > 2 {
		c = 1
	} else {
		c = 2
	}
}
