package main

import "fmt"

func main() {
	// 1. Crear un map
	colores := map[string]string{
		"rojo":     "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
		"amarillo": "#FFFF00",
		"blanco":   "#FFFFFF",
		"negro":    "#000000",
	}

	fmt.Println("Map original:", colores)

	// 2. Acceder a un valor por clave
	fmt.Println("Código hexadecimal del rojo:", colores["rojo"])

	// 3. Agregar un nuevo par clave-valor
	colores["gris"] = "#808080"
	fmt.Println("Map después de agregar 'gris':", colores)

	// 4. Modificar un valor existente
	colores["azul"] = "#0000CC"
	fmt.Println("Map después de modificar 'azul':", colores)

	// 5. Eliminar un par clave-valor
	delete(colores, "amarillo")
	fmt.Println("Map después de eliminar 'amarillo':", colores)

	// 6. Verificar si una clave existe
	valor, existe := colores["verde"]
	if existe {
		fmt.Println("El valor de 'verde' es:", valor)
	} else {
		fmt.Println("'verde' no existe en el map")
	}

	// 7. Iterar sobre un map
	fmt.Println("Iterando sobre el map:")
	for clave, valor := range colores {
		fmt.Printf("%s: %s\n", clave, valor)
	}

	// 8. Obtener la longitud del map
	fmt.Println("Longitud del map:", len(colores))

	// 9. Crear un map vacío
	mapaVacio := make(map[string]int)
	fmt.Println("Map vacío:", mapaVacio)

	// 10. Limpiar un map (eliminar todos los elementos)
	for clave := range colores {
		delete(colores, clave)
	}
	fmt.Println("Map después de limpiar:", colores)
}
