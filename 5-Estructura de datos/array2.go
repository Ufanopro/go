package main

import "fmt"

func main() {
	// Array dinamico con ...
	var a = [...]int{10, 20, 30, 40, 50}
	a[0] = 100
	a[1] = 200
	// Recorremos array
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	//Para imprimir indice y valor
	for index, value := range a {
		fmt.Printf("Índice: %d, Valor: %d\n", index, value)
	}
	//Para imprimir solo valor con barra baja _
	for _, value := range a {
		fmt.Printf("Valor: %d\n", value)
	}
}
