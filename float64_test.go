package filo

import (
	"fmt"
	"testing"
)

var (
	floate64BenchLen   int
	float64BenchResult float64
)

func TestFloat64Stack_Len(t *testing.T) {
	stack := NewFloat64Stack()

	if stack.Len() != 0 {
		t.Error("Len() on empty stack is not zero")
	}

	stack.Push(1.3)

	if stack.Len() != 1 {
		t.Error("Len() on 1-value stack is not 1")
	}

	stack.Pop()
	if stack.Len() != 0 {
		t.Error("Len() on drained stack is not zero")
	}

}

func BenchmarkFloat64Stack(b *testing.B) {
	stack := NewFloat64Stack()
	var (
		ln int
		r  float64
	)
	for i := 0; i < b.N; i++ {
		stack.Push(1.3)
		ln = stack.Len()
		r = stack.Pop()
	}
	floate64BenchLen = ln
	float64BenchResult = r
}
