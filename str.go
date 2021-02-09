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
