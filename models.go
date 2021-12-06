package main 

type Consola struct {
	Modelo string 	`"json:modelo"`
	Marca string `"json:marca"`
	Foto string `"json:foto"`
}

type Consolas []Consola

type Juego struct {
	Titulo string `"json:titulo"`
	Recopilatorio string `"json:titulo"`
	Imagen string `"json:imagen"`
}

type Juegos []Juego