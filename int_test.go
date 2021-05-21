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

	stack.Pop()
	if stack.Len() != 0 {
		t.Error("Len() on drained stack is not zero")
	}

}

func BenchmarkIntStack(b *testing.B) {
	stack := NewIntStack()
	var r, ln int
	for i := 0; i < b.N; i++ {
		stack.Push(1)
		ln = stack.Len()
		r = stack.Pop()
	}
