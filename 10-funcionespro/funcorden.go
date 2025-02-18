package main

import "fmt"

func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

func main() {
	result := double(addOne, 3) // Devuelve 7, ya que addOne(6) = 7
	fmt.Println(result)
}
