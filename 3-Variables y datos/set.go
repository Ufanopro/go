package main

import (
	"fmt"
)

// Definimos un tipo Set como un mapa de claves de tipo int y valores de tipo bool
type Set map[int]bool

// Método para agregar un elemento al conjunto
func (s Set) Add(element int) {
	s[element] = true
}

// Método para eliminar un elemento del conjunto
func (s Set) Remove(element int) {
	delete(s, element)
}

// Método para verificar si un elemento está en el conjunto
func (s Set) Contains(element int) bool {
	_, exists := s[element]
	return exists
}

// Método para obtener el tamaño del conjunto
func (s Set) Size() int {
	return len(s)
}

// Método para imprimir los elementos del conjunto
func (s Set) Print() {
	fmt.Print("Conjunto: { ")
	for element := range s {
		fmt.Print(element, " ")
	}
	fmt.Println("}")
}

func main() {
	// Creamos un conjunto vacío
	mySet := make(Set)

	// Agregamos elementos al conjunto
	mySet.Add(10)
	mySet.Add(20)
	mySet.Add(30)
	mySet.Add(10) // Este elemento no se agregará porque ya existe

	// Imprimimos el conjunto
	mySet.Print()

	// Verificamos si un elemento está en el conjunto
	fmt.Println("¿El conjunto contiene 20?", mySet.Contains(20))
	fmt.Println("¿El conjunto contiene 40?", mySet.Contains(40))

	// Eliminamos un elemento del conjunto
	mySet.Remove(20)
	fmt.Println("Después de eliminar 20:")
	mySet.Print()

	// Obtenemos el tamaño del conjunto
	fmt.Println("Tamaño del conjunto:", mySet.Size())
}
