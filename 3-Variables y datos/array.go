package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Semilla aleatoria para obtener diferentes resultados cada vez
	rand.Seed(time.Now().UnixNano())

	// Crear un array con 5 componentes
	var numeros [5]int

	// Llenar el array con números aleatorios entre 1 y 100
	for i := 0; i < len(numeros); i++ {
		numeros[i] = rand.Intn(100) + 1 // rand.Intn(100) genera un número entre 0 y 99, sumamos 1 para que sea entre 1 y 100
	}

	// Imprimir el contenido del array
	fmt.Println("Contenido del array de números aleatorios:")
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
