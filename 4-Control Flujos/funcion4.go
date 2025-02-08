package main

import (
	"fmt"
)

// Función que suma dos valores
func sumar(a int, b int) int {
	return a + b
}

// Función que realiza suma, resta, multiplicación y división
func operaciones(a int, b int) (int, int, int, float64) {
	suma := a + b
	resta := a - b
	multiplicacion := a * b

	// Manejo de división por cero
	var division float64
	if b != 0 {
		division = float64(a) / float64(b)
	} else {
		division = 0
		fmt.Println("Advertencia: División por cero no permitida.")
	}

	return suma, resta, multiplicacion, division
}

func main() {
	var valor1, valor2 int

	// Solicitamos los valores al usuario
	fmt.Print("Introduce el primer valor: ")
	fmt.Scanln(&valor1)

	fmt.Print("Introduce el segundo valor: ")
	fmt.Scanln(&valor2)

	// Llamamos a la función sumar
	resultadoSuma := sumar(valor1, valor2)
	fmt.Printf("La suma de %d y %d es: %d\n", valor1, valor2, resultadoSuma)

	// Llamamos a la función operaciones
	suma, resta, multiplicacion, division := operaciones(valor1, valor2)

	// Mostramos los resultados
	fmt.Printf("Suma: %d\n", suma)
	fmt.Printf("Resta: %d\n", resta)
	fmt.Printf("Multiplicación: %d\n", multiplicacion)
	fmt.Printf("División: %.2f\n", division)
}
