package settings

import (
	"gopkg.in/mgo.v2"
)

// Prepara la conexión a la base de datos:
func GetSession() *mgo.Session{
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil{
		panic(err)
	}

	return session
}
