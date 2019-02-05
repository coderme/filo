package filo

import (
	"sync"
)

// StringStack stack first in last out
// safe for concurrent usage
type StringStack struct {
	items []string
	mu    *sync.Mutex
}

// Push pushes new item to the stack
func (s *StringStack) Push(i string) {
	s.mu.Lock()
	s.items = append(s.items, i)
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

// NewStringStack creates new StringStack
func NewStringStack() *StringStack {
	return &StringStack{
		mu: &sync.Mutex{},
	}
}
