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
func ConsolasList(w http.ResponseWriter, r *http.Request){
	var results []models.Consola

	err := session.DB("game-collector").C("consolas").Find(nil).All(&results)

	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Resultados", results)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}


func Consola(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	consola_id := params["id"]
	
	if !bson.IsObjectIdHex(consola_id){
		w.WriteHeader(404)
		return 
	}

	oid := bson.ObjectIdHex(consola_id)

	results := models.Consola{}

	err := session.DB("game-collector").C("consolas").FindId(oid).One(&results)

	if err != nil{
		w.WriteHeader(404)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}


func ConsolaCreate(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var consola_data models.Consola
	err := decoder.Decode(&consola_data)

	if err != nil{
		panic(err)
	}

	defer r.Body.Close()

	err = session.DB("game-collector").C("consolas").Insert(consola_data)

	if err != nil{
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(consola_data)
}


func ConsolaUpdate(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	consola_id := params["id"]

	if !bson.IsObjectIdHex(consola_id){
		w.WriteHeader(404)
		return 
	}

	decoder := json.NewDecoder(r.Body)

	oid := bson.ObjectIdHex(consola_id)

	var consola_data models.Consola 

	err := decoder.Decode(&consola_data)

	if err != nil{
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id":oid}
	change := bson.M{"$set": consola_data}
	err = session.DB("game-collector").C("consolas").Update(document, change)

	if err != nil{
		w.WriteHeader(404)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(consola_data)
}


func ConsolaDelete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	consola_id := params["id"]

	if !bson.IsObjectIdHex(consola_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(consola_id)
	err := session.DB("game-collector").C("consolas").RemoveId(oid)

	if err != nil{
		panic(err)
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Se ha eliminado la consola")
}