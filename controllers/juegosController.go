package controllers

import (
	"fmt"
	"net/http"
)

func juegosList(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Listado de juegos</h1>")
}