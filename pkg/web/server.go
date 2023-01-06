package web

import (
	"log"
	"net/http"
)

type WebPhoneBook struct {
	InfoLog *log.Logger
	ErrorLog *log.Logger
	ServerMux *http.ServeMux
	Host string
}


func (app *WebPhoneBook) ServerStartup() {
	srv := http.Server{
		Addr: app.Host,
		Handler: app.ServerMux,
	}
	app.InfoLog.Printf("Server Starts on localhost:%v", srv.Addr)
	app.ErrorLog.Fatal(srv.ListenAndServe())
}

func (app *WebPhoneBook) RoutersSetup() {
	app.ServerMux = http.NewServeMux()
	app.ServerMux.HandleFunc("/", HomeHandler)
}