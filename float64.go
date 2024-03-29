package filo

import (
	"sync"
)

// Float64Stack integer stack first in last out
// safe for concurrent usage
type Float64Stack struct {
	items []float64
	mu    *sync.RWMutex
}

// Push pushes new item to the stack
func (f *Float64Stack) Push(j float64) {
	f.mu.Lock()
	f.items = append(f.items, j)
	f.mu.Unlock()
}

// Pop pops the last string from the stack
func (f *Float64Stack) Pop() float64 {
	f.mu.Lock()
	defer f.mu.Unlock()

	ln := len(f.items)
	if ln == 0 {
		return 0
	}

	tail := f.items[ln-1]
	f.items = f.items[:ln-1]

	return tail
}

// Len gets the number of items pushed
// into the stack
func (f *Float64Stack) Len() int {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return len(f.items)
}

// NewFloat64Stack creates new IntStack
func NewFloat64Stack() *Float64Stack {
	return &Float64Stack{
		mu: &sync.RWMutex{},
	}
}

