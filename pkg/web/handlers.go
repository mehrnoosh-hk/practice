package web

import (
	"encoding/json"
	"github/practice/pkg/domain/entity"
	"github/practice/pkg/domain/services"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World <3"))
}

func (app WebPhoneBook) CreateContact(w http.ResponseWriter, r *http.Request) {
	var c entity.Contact
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	id, err := services.CreateContact(c, app.Storage)
  w.Write([]byte(id))
}
