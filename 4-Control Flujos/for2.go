package main

import (
	"fmt"
)

func main() {
	// Bucle for para mostrar valores de 1 a 10 y detenerse cuando i sea 5
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Se ha alcanzado el valor 5, deteniendo el bucle.")
			break
		}
		fmt.Printf("Valor de i: %d\n", i)
	}
}
