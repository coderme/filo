package filo

import (
	"sync"
)

// IntStack integer stack first in last out
// safe for concurrent usage
type IntStack struct {
	items []int
	mu    *sync.RWMutex
}

// Push pushes new item to the stack
func (i *IntStack) Push(j int) {
	i.mu.Lock()
	i.items = append(i.items, j)
	i.mu.Unlock()
}

// Pop pops the last string from the stack
func (i *IntStack) Pop() int {
	i.mu.Lock()
	defer i.mu.Unlock()

	ln := len(i.items)
	if ln == 0 {
		return 0
	}

	tail := i.items[ln-1]
	i.items = i.items[:ln-1]
