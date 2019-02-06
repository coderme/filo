package filo

import (
	"fmt"
	"testing"
)

func TestStringStack_Len(t *testing.T) {
	stack := NewStringStack()

	if stack.Len() != 0 {
		t.Error("Len() on empty stack is not zero")
	}

	stack.Push("a")

	if stack.Len() != 1 {
		t.Error("Len() on 1-value stack is not 1")
	}

	stack.Pop()
	if stack.Len() != 0 {
		t.Error("Len() on drained stack is not zero")
	}

}

func BenchmarkStringStack(b *testing.B) {
	stack := NewStringStack()
	for i := 0; i < b.N; i++ {
		stack.Push("a")
		stack.Len()
		stack.Pop()
	}

}

func ExampleStringStack() {
	stack := NewStringStack()
	data := []string{
		"a",
		"b",
		"c",
	}

	for _, d := range data {
		stack.Push(d)
	}

	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}

	// Output:
	// c
	// b
	// a

}
