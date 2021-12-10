package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"game-collector/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// para que una función sea pública usar mayúsculas:
func JuegosList(w http.ResponseWriter, r *http.Request){
	var results []models.Juego

	err := session.DB("game-collector").C("juegos").Find(nil).All(&results)

	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Resultados", results)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}


func Juego(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	juego_id := params["id"]
	
	if !bson.IsObjectIdHex(juego_id){
		w.WriteHeader(404)
		return 
	}

	oid := bson.ObjectIdHex(juego_id)

	results := models.Juego{}

	err := session.DB("game-collector").C("juegos").FindId(oid).One(&results)

	if err != nil{
		w.WriteHeader(404)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}


func JuegoCreate(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var juego_data models.Juego
	err := decoder.Decode(&juego_data)

	if err != nil{
		panic(err)
	}

	defer r.Body.Close()

	err = session.DB("game-collector").C("juegos").Insert(juego_data)

	if err != nil{
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(juego_data)
}


func JuegoUpdate(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	juego_id := params["id"]

	if !bson.IsObjectIdHex(juego_id){
		w.WriteHeader(404)
		return 
	}

	decoder := json.NewDecoder(r.Body)

	oid := bson.ObjectIdHex(juego_id)

	var juego_data models.Juego 

	err := decoder.Decode(&juego_data)

	if err != nil{
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id":oid}
	change := bson.M{"$set": juego_data}
	err = session.DB("game-collector").C("juegos").Update(document, change)

	if err != nil{
		w.WriteHeader(404)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(juego_data)
}


func JuegoDelete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	juego_id := params["id"]

	if !bson.IsObjectIdHex(juego_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(juego_id)
	err := session.DB("game-collector").C("juegos").RemoveId(oid)

	if err != nil{
		panic(err)
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Se ha eliminado la juego")
}