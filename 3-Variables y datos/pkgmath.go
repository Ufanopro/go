package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("PI: ", math.Pi) //PI: 3.141592653589793
	fmt.Println("E: ", math.E)   //E: 2.718281828459045

	//funciones de potencia
	fmt.Println(math.Pow(2, 3)) // Imprime: 8
	fmt.Println(math.Sqrt(64))  // Imprime: 8
	fmt.Println(math.Cbrt(27))  // Imprime: 3
	fmt.Println(math.Pow10(2))  // Imprime: 100

	//funcion triangulo
	const precision = 2 // Número de decimales para la salida

	// Solicitar al usuario que ingrese los lados del triángulo rectángulo
	var base, altura float64
	fmt.Print("Ingrese la longitud de la base del triángulo: ")
	fmt.Scanln(&base)
	fmt.Print("Ingrese la longitud de la altura del triángulo: ")
	fmt.Scanln(&altura)

	// Calcular la hipotenusa usando el teorema de Pitágoras
	hipotenusa := math.Sqrt(math.Pow(base, 2) + math.Pow(altura, 2))

	// Calcular el área del triángulo
	area := (base * altura) / 2

	// Calcular el perímetro del triángulo
	perimetro := base + altura + hipotenusa

	// Imprimir el área y el perímetro con la precisión especificada
	fmt.Printf("El área del triángulo es: %.*f\n", precision, area)
	fmt.Printf("El perímetro del triángulo es: %.*f\n", precision, perimetro)

}
