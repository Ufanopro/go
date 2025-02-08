package main

import (
	"fmt"
)

func main() {
	// Tipos de datos básicos en Go

	// Tipos enteros
	var entero int = 42
	var enteroSinSigno uint = 100
	var entero8 int8 = -128
	var entero16 int16 = 32767
	var entero32 int32 = -2147483648
	var entero64 int64 = 9223372036854775807

	// Tipos de punto flotante
	var flotante32 float32 = 3.1415
	var flotante64 float64 = 2.7182818284

	// Booleanos
	var booleano bool = true

	// Cadenas de texto
	var cadena string = "Hola, Go!"

	// Caracter (rune en Go, que representa un Unicode code point)
	var caracter rune = 'G'

	// Tipo byte (alias de uint8)
	var byteValue byte = 255

	// Declaración corta (inferencia de tipos)
	autoEntero := 10
	autoFlotante := 1.23
	autoBooleano := false
	autoCadena := "Inferido"

	// Imprimir los valores y tipos
	fmt.Println("--- Tipos de datos básicos ---")
	fmt.Printf("entero: %d (tipo: %T)\n", entero, entero)
	fmt.Printf("enteroSinSigno: %d (tipo: %T)\n", enteroSinSigno, enteroSinSigno)
	fmt.Printf("entero8: %d (tipo: %T)\n", entero8, entero8)
	fmt.Printf("entero16: %d (tipo: %T)\n", entero16, entero16)
	fmt.Printf("entero32: %d (tipo: %T)\n", entero32, entero32)
	fmt.Printf("entero64: %d (tipo: %T)\n", entero64, entero64)
	fmt.Printf("flotante32: %f (tipo: %T)\n", flotante32, flotante32)
	fmt.Printf("flotante64: %f (tipo: %T)\n", flotante64, flotante64)
	fmt.Printf("booleano: %t (tipo: %T)\n", booleano, booleano)
	fmt.Printf("cadena: %s (tipo: %T)\n", cadena, cadena)
	fmt.Printf("caracter: %c (tipo: %T)\n", caracter, caracter)
	fmt.Printf("byteValue: %d (tipo: %T)\n", byteValue, byteValue)
	fmt.Println("--- Declaración corta e inferencia ---")
	fmt.Printf("autoEntero: %d (tipo: %T)\n", autoEntero, autoEntero)
	fmt.Printf("autoFlotante: %f (tipo: %T)\n", autoFlotante, autoFlotante)
	fmt.Printf("autoBooleano: %t (tipo: %T)\n", autoBooleano, autoBooleano)
	fmt.Printf("autoCadena: %s (tipo: %T)\n", autoCadena, autoCadena)

	// Tipos compuestos
	fmt.Println("--- Tipos de datos compuestos ---")
	// Arrays
	var arreglo [3]int = [3]int{1, 2, 3}
	fmt.Printf("arreglo: %v (tipo: %T)\n", arreglo, arreglo)

	// Slices
	slice := []string{"Go", "es", "genial"}
	fmt.Printf("slice: %v (tipo: %T)\n", slice, slice)

	// Mapas
	mapa := map[string]int{"uno": 1, "dos": 2}
	fmt.Printf("mapa: %v (tipo: %T)\n", mapa, mapa)

	// Estructuras
	type Persona struct {
		Nombre string
		Edad   int
	}
	persona := Persona{Nombre: "Juan", Edad: 30}
	fmt.Printf("persona: %+v (tipo: %T)\n", persona, persona)

	// Punteros
	puntero := &entero
	fmt.Printf("puntero: %v (tipo: %T)\n", puntero, puntero)
}
