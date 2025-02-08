package main

import (
	"fmt"
	"log"
)

func main() {
	// Pide al usuario que ingrese su apellido
	var apellido string
	fmt.Print("Por favor, ingresa tu apellido: ")
	_, err := fmt.Scanln(&apellido)
	if err != nil {
		log.Fatal(err)
	}

	// Saluda al usuario usando su apellido
	fmt.Printf("¡Hola, %s!\n", apellido)
}
