package main

import "fmt"

func main() {
	// Crear una matriz 3x3
	var matriz [3][3]int

	// Rellenar la matriz con valores fijos
	matriz = [3][3]int{
		{4, 4, 4},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Mostrar la matriz por pantalla
	fmt.Println("Matriz 3x3:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", matriz[i][j]) // Usar \t para alinear las columnas
		}
		fmt.Println() // Salto de línea después de cada fila
	}

	// Rellenar la matriz con valores dinamicos
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = i*3 + j + 1 // Rellenar con valores secuenciales
		}
	}

	// Mostrar la matriz por pantalla
	fmt.Println("Matriz 3x3:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", matriz[i][j]) // Usar \t para alinear las columnas
		}
		fmt.Println() // Salto de línea después de cada fila
	}

}
