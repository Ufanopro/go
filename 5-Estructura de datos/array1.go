package main

import "fmt"

func main() {
	// Crear un array con los días de la semana
	diasSemana := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}

	// Recorrer el array e imprimir cada día con su posición
	for i, dia := range diasSemana {
		fmt.Printf("Posición %d: %s\n", i, dia)
	}
}
