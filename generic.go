package filo

import (
	"sync"
)

// GenericStack stack first in last out
// safe for concurrent usage
type GenericStack struct {
	items []interface{}
	mu    *sync.RWMutex
}

// Push pushes new interface{} to the stack
func (g *GenericStack) Push(j interface{}) {
	g.mu.Lock()
	g.items = append(g.items, j)
	g.mu.Unlock()
}

// Pop pops the last interface{} from the stack
func (g *GenericStack) Pop() interface{} {
	g.mu.Lock()
	defer g.mu.Unlock()
