package storage


type InMemoryDB map[string]interface{}


func NewInMemoryDB() InMemoryDB {
	MemoryDB := make(map[string]interface{})
	return MemoryDB
}