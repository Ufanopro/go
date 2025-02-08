package main

import (
	"fmt"
)

// Definimos un struct llamado "Persona"
type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

// Método para imprimir los detalles de una Persona
func (p Persona) ImprimirDetalles() {
	fmt.Printf("Nombre: %s\n", p.Nombre)
	fmt.Printf("Edad: %d\n", p.Edad)
	fmt.Printf("Email: %s\n", p.Email)
}

func main() {
	// Creamos una instancia de Persona
	persona1 := Persona{
		Nombre: "Juan Pérez",
		Edad:   30,
		Email:  "juan.perez@example.com",
	}

	// Imprimimos los detalles de la persona
	fmt.Println("Detalles de la persona 1:")
	persona1.ImprimirDetalles()

	// Creamos otra instancia de Persona
	persona2 := Persona{
		Nombre: "María Gómez",
		Edad:   25,
		Email:  "maria.gomez@example.com",
	}

	// Imprimimos los detalles de la segunda persona
	fmt.Println("\nDetalles de la persona 2:")
	persona2.ImprimirDetalles()

	// Accedemos a los campos del struct directamente
	fmt.Println("\nAccediendo a los campos directamente:")
	fmt.Println("Nombre de persona1:", persona1.Nombre)
	fmt.Println("Edad de persona2:", persona2.Edad)
}
