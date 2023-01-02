package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World"))
	})
	log.Println("Server listen on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}