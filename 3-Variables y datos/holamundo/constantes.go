package main

import "fmt"

const pi = 3.141516
const (
	Lunes = iota + 1
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)

func main() {
	fmt.Println("Hola, Mundo!")
	fmt.Println("Valor de pi", pi)
	fmt.Println(Viernes)
}
