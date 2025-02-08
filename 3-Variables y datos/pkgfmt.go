package main

import "fmt"

func main() {

	// Definición de variables
	name := "Cesar" // Variable tipo string
	age := 45       // Variable tipo int

	// Imprimir el mensaje formateado
	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)

	// Validación adicional (opcional): Verificar valores válidos
	if age < 0 {
		fmt.Println("Error: La edad no puede ser negativa.")
	} else {
		fmt.Printf("La edad ingresada (%d) es válida.\n", age)
	}
}
