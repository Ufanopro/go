package main

import (
	"fmt"
)

/*
type Persona struct {
	Nombre string
	Edad   int
	Correo string
}

func main() {
	// Instancia de Persona
	p1 := Persona{"Juan", 30, "juan@example.com"}
	p2 := Persona{Nombre: "Ana", Edad: 25, Correo: "ana@example.com"}

	// Modificación de valores
	p2.Edad = 26

	fmt.Println(p1)
	fmt.Println(p2)
}


type Circulo struct {
	Radio float64
}

// Método para calcular el área
func (c Circulo) Area() float64 {
	return 3.14 * c.Radio * c.Radio
}

func main() {
	c := Circulo{Radio: 5}
	fmt.Println("Área del círculo:", c.Area())
}
//PUNTERO CON STRUCT
type Contador struct {
	Valor int
}

func (c *Contador) Incrementar() {
	c.Valor++
}

func main() {
	c := Contador{Valor: 10}
	c.Incrementar()
	fmt.Println("Nuevo valor:", c.Valor)
}
// STRUCT ANIDADOS
type Direccion struct {
	Calle  string
	Ciudad string
}

type Persona struct {
	Nombre    string
	Edad      int
	Domicilio Direccion
}

func main() {
	p := Persona{
		Nombre: "Carlos",
		Edad:   40,
		Domicilio: Direccion{
			Calle:  "Av. Principal 123",
			Ciudad: "Madrid",
		},
	}

	fmt.Println(p.Nombre, "vive en", p.Domicilio.Calle, ",", p.Domicilio.Ciudad)
}
// STRUCT con JSON (Marshal y Unmarshal)
// Definimos un struct con etiquetas JSON
type Usuario struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	u := Usuario{"Mario", 28}

	// Convertir a JSON
	jsonData, _ := json.Marshal(u)
	fmt.Println(string(jsonData)) // {"nombre":"Mario","edad":28}

	// Convertir JSON a struct
	jsonStr := `{"nombre":"Laura","edad":32}`
	var u2 Usuario
	json.Unmarshal([]byte(jsonStr), &u2)
	fmt.Println(u2)
}
*/
// STRUCT con interfaces
// Definir una interfaz
type Forma interface {
	Area() float64
}

// Implementar la interfaz en diferentes structs
type Cuadrado struct {
	Lado float64
}

func (c Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func main() {
	var f Forma = Cuadrado{Lado: 4}
	fmt.Println("Área del cuadrado:", f.Area())

	f = Rectangulo{Ancho: 5, Alto: 3}
	fmt.Println("Área del rectángulo:", f.Area())
}
