package main

import (
	"fmt"
	"time"
)

func main() {
	// Obtener la hora actual
	now := time.Now()
	hour := now.Hour()

	// Determinar el saludo según la hora
	if hour < 12 {
		fmt.Println("¡Buenos días!")
	} else if hour < 21 {
		fmt.Println("¡Buenas tardes!")
	} else {
		fmt.Println("¡Buenas noches!")
	}
}
