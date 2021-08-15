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

	if stack.Len() != 1 {
		t.Error("Len() on 1-value stack is not 1")
	}

	stack.Pop()
	if stack.Len() != 0 {
		t.Error("Len() on drained stack is not zero")
	}

}

func BenchmarkGenericStack(b *testing.B) {
	stack := NewGenericStack()
	var ln int
	var r interface{}
	for i := 0; i < b.N; i++ {
		stack.Push(1)
		ln = stack.Len()
		r = stack.Pop()
	}
	genericBenchLen = ln
	genericBenchResult = r

}

func ExampleGenericStack() {
	stack := NewGenericStack()
	data := []interface{}{
		1,
		"a",
		true,
		nil,
		'b',
		[]string{
			"g",
			"o",
		},
		[]int{
			10,
			20,
			30,
		},
	}

	for _, d := range data {
		stack.Push(d)
	}

	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}

	// Output:
