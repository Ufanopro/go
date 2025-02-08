package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		defer fmt.Println(3)
		fmt.Println(1)
		fmt.Println(2)
	*/
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Cuando acabe el flujo realiza esta funcion
	defer file.Close()

	_, err = file.Write([]byte("Hola, Cesar Diaz"))
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

}
