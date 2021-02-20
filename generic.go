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
