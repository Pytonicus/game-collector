package controllers 

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Página principal de la Api</h1>")
}