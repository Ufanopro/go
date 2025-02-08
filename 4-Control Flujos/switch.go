package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Obtener el sistema operativo en el que se ejecuta el programa
	os := runtime.GOOS

	// Determinar el sistema operativo con un switch
	switch os {
	case "windows":
		fmt.Println("Estás ejecutando Go en Windows.")
	case "darwin":
		fmt.Println("Estás ejecutando Go en macOS.")
	case "linux":
		fmt.Println("Estás ejecutando Go en Linux.")
	default:
		fmt.Printf("Sistema operativo no reconocido: %s\n", os)
	}
}
