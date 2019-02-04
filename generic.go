package filo

import (
	"sync"
)

// GenericStack stack first in last out
// safe for concurrent usage
type GenericStack struct {
	items []interface{}
	mu    *sync.Mutex
}

// Push pushes new item to the stack
func (g *GenericStack) Push(i interface{}) {
	g.mu.Lock()
	g.items = append(g.items, i)
	g.mu.Unlock()
}

// Pop pops the last string from the stack
func (g *GenericStack) Pop() interface{} {
	g.mu.Lock()
	defer g.mu.Unlock()
	ln := len(g.items)
	if ln == 0 {
		return nil
	}
	tail := g.items[ln-1]
	g.items = g.items[:ln-1]
	return tail
}

// NewGenericStack creates new GenericStack
func NewGenericStack() *GenericStack {
	return &GenericStack{
		mu: &sync.Mutex{},
	}
}
