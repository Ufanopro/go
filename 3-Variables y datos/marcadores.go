package main

import "fmt"

func main() {
	entero := 42
	flotante := 3.14
	cadena := "Hola, Go!"
	booleano := true
	caracter := 'G'

	fmt.Printf("Entero en base 10: %d\n", entero)
	fmt.Printf("Entero en binario: %b\n", entero)
	fmt.Printf("Flotante (decimal): %f\n", flotante)
	fmt.Printf("Flotante (científica): %e\n", flotante)
	fmt.Printf("Cadena: %s\n", cadena)
	fmt.Printf("Booleano: %t\n", booleano)
	fmt.Printf("Carácter Unicode: %c\n", caracter)
	fmt.Printf("Dirección de memoria: %p\n", &entero)
	fmt.Printf("Tipo del valor: %T\n", flotante)
}
