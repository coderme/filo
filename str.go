package filo

import (
	"sync"
)

// StringStack stack first in last out
// safe for concurrent usage
type StringStack struct {
	items []string
	mu    *sync.RWMutex
}

// Push pushes new item to the stack
func (s *StringStack) Push(j string) {
