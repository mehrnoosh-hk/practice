package services

import "github/practice/pkg/domain/entity"

type Storage interface {
  Create(c entity.Contact) (string, error)

}

func CreateContact(c entity.Contact, s Storage) (string,  error) {
  id, err := s.Create(c)
  return id, err
}