package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Definimos la estructura de una persona
type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

// Función para mostrar los elementos del array ordenados por edad ascendente
func mostrarPorEdadAscendente(personas []Persona) {
	// Ordenamos el slice de personas por edad
	sort.Slice(personas, func(i, j int) bool {
		return personas[i].Edad < personas[j].Edad
	})

	// Mostramos los elementos ordenados
	fmt.Println("Personas ordenadas por edad (ascendente):")
	for _, persona := range personas {
		fmt.Printf("Nombre: %s %s, Edad: %d\n", persona.Nombre, persona.Apellido, persona.Edad)
	}
}

func main() {
	// Creamos un slice de personas con 10 elementos
	personas := []Persona{
		{Nombre: "Juan", Apellido: "Pérez", Edad: 25},
		{Nombre: "María", Apellido: "Gómez", Edad: 30},
		{Nombre: "Carlos", Apellido: "López", Edad: 22},
		{Nombre: "Ana", Apellido: "Martínez", Edad: 28},
		{Nombre: "Luis", Apellido: "Rodríguez", Edad: 35},
		{Nombre: "Laura", Apellido: "Fernández", Edad: 27},
		{Nombre: "Pedro", Apellido: "Sánchez", Edad: 40},
		{Nombre: "Sofía", Apellido: "Díaz", Edad: 33},
		{Nombre: "Miguel", Apellido: "Hernández", Edad: 29},
		{Nombre: "Elena", Apellido: "Jiménez", Edad: 31},
	}

	// Convertimos el slice de personas a JSON (solo para demostración)
	jsonData, err := json.Marshal(personas)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Imprimimos el JSON resultante
	fmt.Println("Array JSON original:")
	fmt.Println(string(jsonData))

	// Llamamos a la función para mostrar los elementos ordenados por edad
	mostrarPorEdadAscendente(personas)
}
