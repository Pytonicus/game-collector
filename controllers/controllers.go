package controllers 

import (
	"fmt"
	"net/http"
	"game-collector/settings"
)

var session = settings.GetSession()

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Página principal de la Api</h1>")
}