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
	s.mu.Lock()
	s.items = append(s.items, j)
	s.mu.Unlock()
}

// Pop pops the last string from the stack
func (s *StringStack) Pop() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	ln := len(s.items)
	if ln == 0 {
		return ""
	}

	tail := s.items[ln-1]
	s.items = s.items[:ln-1]

	return tail
}

// Len gets the number of items pushed
// into the stack
func (s *StringStack) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// NewStringStack creates new StringStack
func NewStringStack() *StringStack {
	return &StringStack{
		mu: &sync.RWMutex{},
	}
}

