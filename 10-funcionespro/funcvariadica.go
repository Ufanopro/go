package main

import "fmt"

func suma(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	return total

}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {

	fmt.Println(suma(12, 45, 78, 56))
	fmt.Println(suma(10, 20, 3, 56))
	imprimirDatos("Hola", 28, true, 3.14)

}
