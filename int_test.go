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

	stack.Push(1)

	if stack.Len() != 1 {
		t.Error("Len() on 1-value stack is not 1")
	}

