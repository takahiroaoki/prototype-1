package main

import (
	"backend/handler"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", handler.NewMainHandler())
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("resources/static"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListernAndServe: ", err)
	}
}
