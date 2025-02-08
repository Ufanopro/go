package main

import (
	"fmt"
	"time"
)

// Función que envía datos a un channel
func enviarMensaje(ch chan<- string, mensaje string) {
	fmt.Println("Enviando mensaje:", mensaje)
	time.Sleep(2 * time.Second) // Simulamos un retraso
	ch <- mensaje               // Enviamos el mensaje al channel
	fmt.Println("Mensaje enviado:", mensaje)
}

// Función que recibe datos de un channel
func recibirMensaje(ch <-chan string) {
	fmt.Println("Esperando recibir un mensaje...")
	mensaje := <-ch // Recibimos el mensaje del channel
	fmt.Println("Mensaje recibido:", mensaje)
}

func main() {
	// Creamos un channel de tipo string
	ch := make(chan string)

	// Lanzamos una goroutine para enviar un mensaje
	go enviarMensaje(ch, "Hola, mundo!")

	// Lanzamos una goroutine para recibir el mensaje
	go recibirMensaje(ch)

	// Esperamos un momento para que las goroutines terminen
	time.Sleep(3 * time.Second)

	fmt.Println("Programa terminado.")
}
