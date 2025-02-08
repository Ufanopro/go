package main

import "fmt"

func main() {
	// Crear un map donde la clave es un string y el valor es un int
	empleados := make(map[string]int)

	// Asignar valores al map
	empleados["Juan"] = 1001
	empleados["Ana"] = 1002
	empleados["Pedro"] = 1003

	// Imprimir el map
	fmt.Println("Empleados: ", empleados)

	// Acceder a un valor mediante su clave
	idJuan := empleados["Juan"]
	fmt.Println("ID de Juan:", idJuan)
}
