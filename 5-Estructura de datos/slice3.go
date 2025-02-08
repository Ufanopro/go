package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)
	nombres := make([]string, 5, 10)
	nombres[0] = "César"
	fmt.Println(nombres)
	copy(slice2, slice2)
	fmt.Println(slice2)
	fmt.Println(slice1)
}
