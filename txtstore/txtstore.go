package txtstore

import "github.com/google/uuid"

var store = make(map[string]string)

// Write writes into a map a returns the used id
func Write(input string) string {
	id := uuid.New().String()
	store[id] = input
	return id
}

// Read reads from a map with the given id
func Read(id string) string {
	return store[id]
}
