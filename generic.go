package filo

import (
	"sync"
)

// GenericStack stack first in last out
// safe for concurrent usage
type GenericStack struct {
	items []interface{}
