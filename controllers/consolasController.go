package controllers

import (
	"fmt"
	"net/http"
)

func consolasList(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Listado de consolas</h1>")
}