package main

import (
	"fmt"
)

func main() {
	// Bucle for para mostrar valores de 1 a 10, omitiendo el 5
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("Valor de i: %d\n", i)
	}
}
