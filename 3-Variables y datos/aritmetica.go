package main

import "fmt"

func main() {

	a := 10
	b := 3
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3
	fmt.Println(a % b) // 1

	//incremento y decremento
	a = 10
	a++
	fmt.Println(a) // 11
	b = 5
	b--
	fmt.Println(b) // 4

	//operadores en asignacion
	a = 10
	a += 5
	fmt.Println(a) // 15
	b = 20
	b -= 3
	fmt.Println(b) // 17
	c := 7
	c *= 3
	fmt.Println(c) // 21
	d := 100
	d /= 5
	fmt.Println(d) // 20
	e := 15
	e %= 4
	fmt.Println(e) // 3
}
