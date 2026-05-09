package core

import "sync"

var (
	nextID = 1
	mu     sync.Mutex
)

func generateID() int {
	mu.Lock()
	defer mu.Unlock()
	id := nextID
	nextID++
	return id
}
