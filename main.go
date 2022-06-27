package main

import (
	"log"

	"github.com/Mariosay/twittor/bd"
	"github.com/Mariosay/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a la BD")
	}
	handlers.Manejadores()
}
