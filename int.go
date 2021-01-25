package filo

import (
	"sync"
)

// IntStack integer stack first in last out
// safe for concurrent usage
type IntStack struct {
