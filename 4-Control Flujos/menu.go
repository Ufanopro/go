package main

import (
	"fmt"
)

// Función para mostrar el menú
func mostrarMenu() {
	fmt.Println("----- Menú -----")
	fmt.Println("1. Crear")
	fmt.Println("2. Modificar")
	fmt.Println("3. Borrar")
	fmt.Println("4. Listar")
	fmt.Println("5. Salir")
	fmt.Print("Selecciona una opción (1-5): ")
}

func main() {
	var opcion int

	for {
		// Mostrar el menú
		mostrarMenu()

		// Leer la opción del usuario
		fmt.Scanln(&opcion)

		// Procesar la opción seleccionada
		switch opcion {
		case 1:
			fmt.Println("Has seleccionado: Crear")
			// Aquí puedes añadir la lógica para "Crear"
		case 2:
			fmt.Println("Has seleccionado: Modificar")
			// Aquí puedes añadir la lógica para "Modificar"
		case 3:
			fmt.Println("Has seleccionado: Borrar")
			// Aquí puedes añadir la lógica para "Borrar"
		case 4:
			fmt.Println("Has seleccionado: Listar")
			// Aquí puedes añadir la lógica para "Listar"
		case 5:
			fmt.Println("Saliendo del programa...")
			return // Salir del programa
		default:
			fmt.Println("Opción no válida. Por favor, selecciona una opción del 1 al 5.")
		}

		// Separador para mejor legibilidad
		fmt.Println("-----------------------------")
	}
}
