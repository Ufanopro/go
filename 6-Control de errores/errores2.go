package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}
	return dividendo / divisor, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error en la división:", err)
		return
	}

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error en la conversión:", err)
		return
	}

	fmt.Println("Número:", num)
	fmt.Println("Resultado:", result)
}
