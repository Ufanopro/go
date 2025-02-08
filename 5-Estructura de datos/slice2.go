package main

import "fmt"

func main() {
	// Crear el slice diasSemana con los días de la semana
	diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}

	// Crear el slice numerosSemana con los números del 1 al 7
	numerosSemana := []int{1, 2, 3, 4, 5, 6, 7}

	// Imprimir los slices
	fmt.Println("Días de la semana:", diasSemana)
	fmt.Println("Números de la semana:", numerosSemana)

	// Imprimir los días con sus números correspondientes
	fmt.Println("\nDías con sus números:")
	for i := 0; i < len(diasSemana); i++ {
		fmt.Printf("%d: %s\n", numerosSemana[i], diasSemana[i])
	}
	// Mostrar solo los días impares
	fmt.Println("\nDías impares:")
	for i := 0; i < len(numerosSemana); i++ {
		if numerosSemana[i]%2 != 0 { // Verificar si el número es impar
			fmt.Printf("%d: %s\n", numerosSemana[i], diasSemana[i])
		}
	}
	//Mostrar los indices deseados
	sliceCustom := diasSemana[0:5]
	fmt.Println(sliceCustom)
	//Longitud slice
	fmt.Println(len(sliceCustom))
	fmt.Println(cap(sliceCustom))

}
