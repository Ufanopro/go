package main

import "fmt"

func saludar(name string, f func(string)) {
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {

	saludo := func(name string) {
		fmt.Printf("¡Hola, %s!\n", name)
	}
	saludar("Cesar", saludo)

	resultado1 := aplicar(duplicar, 5)  // resultado1 = 10
	resultado2 := aplicar(triplicar, 5) // resultado2 = 15

	fmt.Println(resultado1, resultado2)

}
