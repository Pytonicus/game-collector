package main

import (
	"net/http"
	"log"
)

func main(){
	router := NewRouter()

	server := http.ListenAndServe(":8000", router)
	log.Fatal(server)
}