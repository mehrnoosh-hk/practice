package storage

import (
  "github.com/rs/xid"
)

type InMemoryDB map[string]interface{}


func NewInMemoryDB() InMemoryDB {
	MemoryDB := make(map[string]interface{})
	return MemoryDB
}


func(IMDB InMemoryDB) Create(c interface{}) (string, error){
  id := xid.New()
  IMDB[id] = c
}