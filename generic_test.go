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

