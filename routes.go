package main 

import (
	"net/http"
	"github.com/gorilla/mux"
	"game-collector/controllers"
)

type Route struct{
	Name string
	Method string
	Pattern string 
	HandlerFunc http.HandlerFunc 
}

type Routes []Route 

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes{
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{"index", "GET", "/", controllers.Index},
	// Rutas consolas:
	Route{"consolas", "GET", "/consolas", controllers.ConsolasList},
	Route{"consola", "GET", "/consola/{id}", controllers.Consola},
	Route{"consola", "POST", "/consola", controllers.ConsolaCreate},
	Route{"consola", "PUT", "/consola/{id}", controllers.ConsolaUpdate},
	Route{"consola", "DELETE", "/consola/{id}", controllers.ConsolaDelete},
	// Rutas juegos:
	Route{"juegos", "GET", "/juegos", controllers.JuegosList},
	Route{"juego", "GET", "/juego/{id}", controllers.Juego},
	Route{"juego", "POST", "/juego", controllers.JuegoCreate},
	Route{"juego", "PUT", "/juego/{id}", controllers.JuegoUpdate},
	Route{"juego", "DELETE", "/juego/{id}", controllers.JuegoDelete}}