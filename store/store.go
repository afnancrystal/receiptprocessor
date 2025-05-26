package store

import "sync"

// In-memory map store of receipt ID -> points
var Receipts = make(map[string]int)
var Mu sync.Mutex
