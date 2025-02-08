package main

import "fmt"

func main() {
	// 1. Declarar una variable y un puntero
	var numero int = 42
	var ptr *int // Puntero a un entero

	// 2. Asignar la dirección de la variable al puntero
	ptr = &numero
	fmt.Println("Valor de numero:", numero)
	fmt.Println("Dirección de numero:", &numero)
	fmt.Println("Valor de ptr (dirección almacenada):", ptr)
	fmt.Println("Valor al que apunta ptr:", *ptr)

	// 3. Modificar el valor a través del puntero
	*ptr = 100
	fmt.Println("Nuevo valor de numero (modificado a través del puntero):", numero)

	// 4. Puntero a un puntero
	var ptr2 **int = &ptr
	fmt.Println("Valor de ptr2 (dirección de ptr):", ptr2)
	fmt.Println("Valor al que apunta ptr2 (dirección almacenada en ptr):", *ptr2)
	fmt.Println("Valor al que apunta *ptr2 (valor de numero):", **ptr2)

	// 5. Puntero a un struct
	type Persona struct {
		nombre string
		edad   int
	}

	persona := Persona{nombre: "Juan", edad: 30}
	ptrPersona := &persona
	fmt.Println("Nombre de la persona (accedido a través del puntero):", ptrPersona.nombre)
	fmt.Println("Edad de la persona (accedido a través del puntero):", ptrPersona.edad)

	// 6. Modificar un struct a través de un puntero
	ptrPersona.edad = 35
	fmt.Println("Edad modificada a través del puntero:", persona.edad)

	// 7. Puntero a un slice
	slice := []int{1, 2, 3, 4, 5}
	ptrSlice := &slice
	fmt.Println("Slice original:", *ptrSlice)
	(*ptrSlice)[0] = 10 // Modificar el slice a través del puntero
	fmt.Println("Slice modificado:", *ptrSlice)

	// 8. Puntero a un array
	array := [3]int{10, 20, 30}
	ptrArray := &array
	fmt.Println("Array original:", *ptrArray)
	(*ptrArray)[1] = 200 // Modificar el array a través del puntero
	fmt.Println("Array modificado:", *ptrArray)

	// 9. Puntero nil (puntero no inicializado)
	var ptrNil *int
	if ptrNil == nil {
		fmt.Println("ptrNil es nil (no apunta a ninguna dirección)")
	}

	// 10. Uso de punteros en funciones (paso por referencia)
	valor := 5
	fmt.Println("Valor antes de llamar a la función:", valor)
	incrementar(&valor)
	fmt.Println("Valor después de llamar a la función:", valor)
}

// Función que recibe un puntero y modifica el valor al que apunta
func incrementar(ptr *int) {
	*ptr = *ptr + 1
}
