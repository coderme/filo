package filo

import (
	"fmt"
	"testing"
)

var (
	strBenchLen    int
	strBenchResult string
)

func TestStringStack_Len(t *testing.T) {
	stack := NewStringStack()

	if stack.Len() != 0 {
		t.Error("Len() on empty stack is not zero")
	}

	stack.Push("a")
