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

