package main

import "fmt"

func main() {
	// Crear un map con los colores y sus valores hexadecimales
	colores := map[string]string{
		"rojo":     "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
		"amarillo": "#FFFF00",
		"blanco":   "#FFFFFF",
		"negro":    "#000000",
	}

	// Imprimir el map
	for color, hex := range colores {
		fmt.Printf("%s: %s\n", color, hex)
	}

	//Imprime un valor del map
	fmt.Println(colores["rojo"])

	//añadir un nuevo valor
	colores["rosa"] = "#FFC0CB"

	//Imprime un valor del map
	fmt.Println(colores["rosa"])

	// Imprimir el map
	for color, hex := range colores {
		fmt.Printf("%s: %s\n", color, hex)
	}
	//verificar si existe en map
	/*
		valor, ok := colores["rojo"]
		if ok {
			fmt.Println("Si existe en el map")
		} else {
			fmt.Println("No existe en el map")
		}
		// Lo imprime si existe
		fmt.Println(valor)
	*/
	if valor, ok := colores["verde"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe en el map")
	}
	// Borrar en map
	delete(colores, "azul")
	// Imprimir el map
	for clave, valor := range colores {
		fmt.Printf("%s: %s\n", clave, valor)
	}

}
