package filo

import (
	"fmt"
	"testing"
)

var (
	intBenchLen    int
	intBenchResult int
)

func TestIntStack_Len(t *testing.T) {
	stack := NewIntStack()

	if stack.Len() != 0 {
		t.Error("Len() on empty stack is not zero")
	}
