package filo

import (
	"fmt"
	"testing"
)

var (
	genericBenchLen    int
	genericBenchResult interface{}
)

func TestGenericStack_Len(t *testing.T) {
	stack := NewGenericStack()

	if stack.Len() != 0 {
		t.Error("Len() on empty stack is not zero")
	}

	stack.Push(1)

