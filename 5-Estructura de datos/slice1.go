package main

import "fmt"

func main() {
	// 1. Crear un slice vacío
	var slice1 []int
	fmt.Println("Slice vacío:", slice1)

	// 2. Crear un slice con valores iniciales
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice con valores iniciales:", slice2)

	// 3. Crear un slice a partir de un array
	array := [5]int{10, 20, 30, 40, 50}
	slice3 := array[1:4] // Desde el índice 1 hasta el 3 (no incluye el 4)
	fmt.Println("Slice creado desde un array:", slice3)

	// 4. Modificar un slice (los slices son referencias a arrays subyacentes)
	slice3[0] = 100
	fmt.Println("Slice modificado:", slice3)
	fmt.Println("Array original también modificado:", array)

	// 5. Añadir elementos a un slice usando append
	slice2 = append(slice2, 6, 7, 8)
	fmt.Println("Slice después de append:", slice2)

	// 6. Crear un slice con make (tamaño y capacidad inicial)
	slice4 := make([]int, 3, 5) // Longitud 3, capacidad 5
	slice4[0] = 9
	slice4[1] = 8
	slice4[2] = 7
	fmt.Println("Slice creado con make:", slice4)
	fmt.Printf("Longitud: %d, Capacidad: %d\n", len(slice4), cap(slice4))

	// 7. Copiar un slice
	slice5 := make([]int, len(slice2))
	copy(slice5, slice2)
	fmt.Println("Slice copiado:", slice5)

	// 8. Eliminar un elemento de un slice (usando append y rebanado)
	index := 2 // Eliminar el elemento en el índice 2
	slice2 = append(slice2[:index], slice2[index+1:]...)
	fmt.Println("Slice después de eliminar un elemento:", slice2)

	// 9. Recorrer un slice con un bucle for
	fmt.Println("Recorriendo el slice:")
	for i, v := range slice2 {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}
}
