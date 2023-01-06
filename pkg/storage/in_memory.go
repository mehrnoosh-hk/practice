package storage

import (
	"github/practice/pkg/domain/entity"
	"strconv"
)

type InMemoryDB map[string]entity.Contact


func NewInMemoryDB() InMemoryDB {
	MemoryDB := make(map[string]entity.Contact)
	return MemoryDB
}


func(IMDB InMemoryDB) Create(c entity.Contact) (string, error){
  id := strconv.Itoa(len(IMDB) + 1)
  IMDB[id] = c
  return id, nil 
}