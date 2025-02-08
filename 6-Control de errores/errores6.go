package main

import (
	"log"
	"os"
)

func main() {

	log.Print("Primer mensaje de registro")

	//Para detener el programa Fatal y Panic
	//log.Fatal("Este registro para el flujo")
	//log.Panic("Este registro para el flujo")
	log.Print("Segundo mensaje de registro")

	//Añadir prefijo o tag
	log.SetPrefix(("MAIN "))
	log.Print("Primer mensaje de registro")
	log.Print("Segundo mensaje de registro")

	//Registro de logs en un archivo
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Cerramos archivo
	defer file.Close()

	//fijamos la escritura en el archivo
	log.SetOutput(file)
	log.Print("¡Log en archivo!")

}
