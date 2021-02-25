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
