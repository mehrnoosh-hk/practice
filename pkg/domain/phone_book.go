package domain

import (
	ST "github/practice/pkg/storage"
	"log"
)

type Contact struct {
	Name string
	Phone string
}


func (c *Contact) AddContact(s ST.Storgae) (string, error){
	id, err := s.Create(c)

	if err != nil {
		log.Printf("Error occured during creating new contact %q", err)
		return "", err
	}

	return id, nil
}


func (c *Contact) UpdateContact (){}


func (c *Contact) RemoveContact (){}