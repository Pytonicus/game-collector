package models 
import "gopkg.in/mgo.v2/bson"

type Consola struct {
	Modelo string 	`"json:modelo"`
	Marca string 	`"json:marca"`
	Foto string 	`"json:foto"`
}

type Consolas []Consola


type Juego struct {
	Titulo string 			`"json:titulo"`
	Recopilatorio string 	`"json:recopilatorio"`
	Imagen string 			`"json:imagen"`
	Consola bson.ObjectId	`"json:consola"`
}

type Juegos []Juego


type User struct {
	Name string `"json:name"`
	Password string `"json:password"`
}

type Users []User