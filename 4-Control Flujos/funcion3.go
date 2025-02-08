package main

import "fmt"

// Función que recibe dos enteros y devuelve su suma
func sumar(a int, b int) int {
	return a + b
}

func main() {
	// Definimos dos valores de ejemplo
	valor1 := 5
	valor2 := 3

	// Llamamos a la función sumar y almacenamos el resultado
	resultado := sumar(valor1, valor2)

	// Mostramos el resultado
	fmt.Printf("La suma de %d y %d es: %d\n", valor1, valor2, resultado)
}
