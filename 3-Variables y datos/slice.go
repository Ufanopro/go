package main

import "fmt"

func main() {
	// Crear un slice con valores iniciales
	numeros := []int{10, 20, 30, 40, 50}

	// Imprimir el slice
	fmt.Println("Slice inicial:", numeros)

	// Acceder a un elemento específico
	fmt.Println("Elemento en la posición 2:", numeros[2]) // Imprime 30

	// Modificar un valor en el slice
	numeros[1] = 25
	fmt.Println("Slice después de modificar un elemento:", numeros)

	// Agregar un elemento al final
	numeros = append(numeros, 60)
	fmt.Println("Slice después de agregar un elemento:", numeros)

	// Crear un slice de otro slice (sub-slice)
	subSlice := numeros[1:4] // Obtiene los elementos 20, 25, 30
	fmt.Println("Sub-slice:", subSlice)

	// Longitud y capacidad del slice
	fmt.Println("Longitud del slice:", len(numeros))
	fmt.Println("Capacidad del slice:", cap(numeros))
}
