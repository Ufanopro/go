package main

import (
	"fmt"
)

// Función que puede causar un panic si el divisor es 0
func dividir(dividendo, divisor int) int {
	if divisor == 0 {
		panic("¡Error: No se puede dividir por cero!") // Provoca un panic
	}
	return dividendo / divisor
}

func main() {
	// Llamada a una función que puede causar un panic
	// Usamos recover para manejar el panic
	defer func() {
		if r := recover(); r != nil {
			// recover() captura el valor del panic
			fmt.Println("Recuperado del panic:", r)
		}
	}()

	// Ejemplo de uso de la función dividir
	fmt.Println("Iniciando el programa...")
	resultado := dividir(10, 0)                         // Esto causará un panic
	fmt.Println("Resultado de la división:", resultado) // Esta línea no se ejecutará

	// Esta línea no se alcanzará si ocurre un panic
	fmt.Println("Fin del programa")
}
