package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//inicializa la semilla y genera un numero entre 1 y 100
	rand.New(rand.NewSource(time.Now().UnixNano()))
	target := rand.Intn(100) + 1
	attempts := 10

	fmt.Println("Adivina el número entre 1 y 100. Tienes 10 intentos.")

	scanner := bufio.NewScanner(os.Stdin)
	//bucle para los 10 intentos
	for i := 1; i <= attempts; i++ {
		fmt.Printf("Intento %d: ", i)
		scanner.Scan()
		input := scanner.Text()
		guess, err := strconv.Atoi(input)

		if err != nil || guess < 1 || guess > 100 {
			fmt.Println("Por favor, introduce un número válido entre 1 y 100.")
			continue
		}

		if guess < target {
			fmt.Println("El número es mayor.")
		} else if guess > target {
			fmt.Println("El número es menor.")
		} else {
			fmt.Printf("¡Felicidades! Adivinaste el número %d en %d intentos.\n", target, i)
			return
		}
	}

	fmt.Printf("Lo siento, has agotado tus intentos. El número era %d.\n", target)
}
