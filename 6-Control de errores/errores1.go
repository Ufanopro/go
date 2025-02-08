package main

import (
	"errors"
	"fmt"
	"os"
)

// Función que devuelve un error si el número es negativo
func esPositivo(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("el número no puede ser negativo")
	}
	return num, nil
}

// Función que divide dos números y maneja el error de división por cero
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}

// Función que simula la apertura de un archivo y maneja el error si no existe
func abrirArchivo(nombre string) (*os.File, error) {
	file, err := os.Open(nombre)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %w", err)
	}
	return file, nil
}

func main() {
	// Manejo de error con esPositivo
	num, err := esPositivo(-10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Número positivo:", num)
	}

	// Manejo de error con dividir
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado de la división:", resultado)
	}

	// Manejo de error con abrirArchivo
	file, err := abrirArchivo("archivo_inexistente.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer file.Close()
		fmt.Println("Archivo abierto correctamente")
	}

	// Uso de errors.Is para verificar un tipo específico de error
	_, err = abrirArchivo("archivo_inexistente.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("El archivo no existe")
		} else {
			fmt.Println("Otro error:", err)
		}
	}

	// Uso de errors.As para convertir y manejar un error específico
	_, err = abrirArchivo("archivo_inexistente.txt")
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Printf("Error de ruta: %s, Operación: %s, Ruta: %s\n", pathError.Err, pathError.Op, pathError.Path)
	} else if err != nil {
		fmt.Println("Otro error:", err)
	}
}
