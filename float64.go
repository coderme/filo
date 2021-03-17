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
