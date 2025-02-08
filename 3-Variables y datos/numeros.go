package main

import "fmt"

func main() {
	// Tipos enteros
	var entero int = 42          // int (tamaño depende del sistema)
	var entero32 int32 = 100000  // int32
	var entero64 int64 = 1000000 // int64

	// Tipos de punto flotante
	var flotante32 float32 = 3.14     // float32
	var flotante64 float64 = 3.141592 // float64

	// Tipos complejos
	var complejo64 complex64 = 1 + 2i   // complex64
	var complejo128 complex128 = 3 + 4i // complex128

	// Imprimir los valores y sus tipos
	fmt.Println("Entero (int):", entero)
	fmt.Println("Entero 32-bit (int32):", entero32)
	fmt.Println("Entero 64-bit (int64):", entero64)

	fmt.Println("Flotante 32-bit (float32):", flotante32)
	fmt.Println("Flotante 64-bit (float64):", flotante64)

	fmt.Println("Número complejo 64-bit (complex64):", complejo64)
	fmt.Println("Número complejo 128-bit (complex128):", complejo128)

	// Mostrar los tipos de las variables
	fmt.Printf("Tipo de 'entero': %T\n", entero)
	fmt.Printf("Tipo de 'entero32': %T\n", entero32)
	fmt.Printf("Tipo de 'entero64': %T\n", entero64)
	fmt.Printf("Tipo de 'flotante32': %T\n", flotante32)
	fmt.Printf("Tipo de 'flotante64': %T\n", flotante64)
	fmt.Printf("Tipo de 'complejo64': %T\n", complejo64)
	fmt.Printf("Tipo de 'complejo128': %T\n", complejo128)
}
