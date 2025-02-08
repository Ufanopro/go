package main

import (
	"fmt"
)

// Definimos la función saludar que recibe un string y muestra un mensaje personalizado
func saludar(nombre string) {
	fmt.Printf("¡Hola %s!\n", nombre)
}

func main() {
	// Variable para almacenar la entrada del usuario
	var nombre string

	// Solicitamos al usuario que introduzca su nombre
	fmt.Print("Introduce tu nombre: ")
	fmt.Scanln(&nombre)

	// Llamamos a la función saludar pasando el nombre introducido por el usuario
	saludar(nombre)
}
