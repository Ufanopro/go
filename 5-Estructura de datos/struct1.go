package main

import (
	"fmt"
	"time"
)

// Definir el struct Persona
type Persona struct {
	nombre          string
	apellido        string
	fechaNacimiento time.Time
	telefono        string
	email           string
}

// Método para calcular la edad de la persona
func (p Persona) CalcularEdad() int {
	hoy := time.Now()
	edad := hoy.Year() - p.fechaNacimiento.Year()

	// Ajustar si aún no ha pasado el cumpleaños este año
	if hoy.YearDay() < p.fechaNacimiento.YearDay() {
		edad--
	}
	return edad
}

// Método para imprimir los detalles de la persona
func (p Persona) ImprimirDetalles() {
	fmt.Println("Nombre:", p.nombre)
	fmt.Println("Apellido:", p.apellido)
	fmt.Println("Fecha de Nacimiento:", p.fechaNacimiento.Format("2006-01-02"))
	fmt.Println("Teléfono:", p.telefono)
	fmt.Println("Email:", p.email)
	fmt.Println("Edad:", p.CalcularEdad(), "años")
}

func main() {
	// Crear una instancia de Persona
	persona1 := Persona{
		nombre:          "Juan",
		apellido:        "Pérez",
		fechaNacimiento: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
		telefono:        "555-1234",
		email:           "juan.perez@example.com",
	}

	// 1. Acceder a los campos del struct
	fmt.Println("Nombre:", persona1.nombre)
	fmt.Println("Apellido:", persona1.apellido)

	// 2. Modificar un campo del struct
	persona1.telefono = "555-5678"
	fmt.Println("Teléfono actualizado:", persona1.telefono)

	// 3. Llamar a un método del struct
	persona1.ImprimirDetalles()

	// 4. Crear un struct anónimo
	persona2 := struct {
		nombre   string
		apellido string
	}{
		nombre:   "Ana",
		apellido: "Gómez",
	}
	fmt.Println("Persona anónima:", persona2)

	// 5. Comparar dos structs
	persona3 := Persona{
		nombre:          "Juan",
		apellido:        "Pérez",
		fechaNacimiento: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
		telefono:        "555-1234",
		email:           "juan.perez@example.com",
	}

	if persona1 == persona3 {
		fmt.Println("persona1 y persona3 son iguales")
	} else {
		fmt.Println("persona1 y persona3 son diferentes")
	}

	// 6. Usar un puntero a un struct
	personaPtr := &persona1
	fmt.Println("Accediendo a través de un puntero:", personaPtr.nombre)

	// 7. Modificar un struct a través de un puntero
	personaPtr.telefono = "555-9999"
	fmt.Println("Teléfono modificado a través de un puntero:", persona1.telefono)
}
