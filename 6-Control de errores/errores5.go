package main

import (
	"fmt"
)

func divide(dividendo, divisor int) {
	//Uso defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo, divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")
	}
}

func main() {

	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(100, 4)

}
