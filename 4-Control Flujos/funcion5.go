package main

import (
	"fmt"
)

// Función que recibe tres valores y devuelve tres resultados
func calcularResultados(a, b, c int) (int, int, float64) {
	suma := a + b + c
	producto := a * b * c
	promedio := float64(suma) / 3.0

	return suma, producto, promedio
}

func main() {
	var valor1, valor2, valor3 int

	// Solicitamos los valores al usuario
	fmt.Print("Introduce el primer valor: ")
	fmt.Scanln(&valor1)

	fmt.Print("Introduce el segundo valor: ")
	fmt.Scanln(&valor2)

	fmt.Print("Introduce el tercer valor: ")
	fmt.Scanln(&valor3)

	// Llamamos a la función calcularResultados
	suma, producto, promedio := calcularResultados(valor1, valor2, valor3)

	// Mostramos los resultados
	fmt.Printf("La suma de %d, %d y %d es: %d\n", valor1, valor2, valor3, suma)
	fmt.Printf("El producto de %d, %d y %d es: %d\n", valor1, valor2, valor3, producto)
	fmt.Printf("El promedio de %d, %d y %d es: %.2f\n", valor1, valor2, valor3, promedio)
}
