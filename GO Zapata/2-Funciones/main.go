package main

import (
	"fmt"
)

/*
//SIMPLE
func miFuncion() {
	fmt.Println("Esta es mi funcion")
}

//PARAMETRIZADA
func funcionParametros(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("La suma es ", resultado)
}

// RETORNO SIMPLE
func funcionRetornoSimple(nombre string) string {
	return "Hola " + nombre
}

// RETORNO MULTIPLE
func funcionRetornoMultiple() (string, string, int) {
	return "Cesar", "Juan", 5
}

//FUNCION ANONIMA
var anonima = func() {
	fmt.Println("Hola desde una función anónima!")
}

//CLOSURE
func contador() func() int {
	contador := 0 // Variable externa que será capturada por el closure
	return func() int {
		contador++ // Modifica la variable externa
		return contador
	}
}

// Closure que genera una función para multiplicar por un número específico
func tablaMultiplicar(n int) func(int) int {
	return func(x int) int {
		return n * x
	}
}

func main() {
	nombre := "Cesar"
	apellido := "Diaz"
	edad := 45

	fmt.Printf("Tu nombre es %v %v, tu edad es %v\n", nombre, apellido, edad)
	miFuncion()
	funcionParametros(5, 4)
	fmt.Println(funcionRetornoSimple("Cesar"))
	nombre, apellido, edad = funcionRetornoMultiple()
	fmt.Println(funcionRetornoMultiple())
	anonima()
	incrementar := contador()
	fmt.Println(incrementar())
	fmt.Println(incrementar())
	fmt.Println(incrementar())

	numero := 5
	multiplicar := tablaMultiplicar(numero) // Closure con el número ingresado
	fmt.Println("Tabla de multiplicar del", numero)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, multiplicar(i))
	}
}
*/
//GORUTINES y CHANNELS
/*
func saludar(nombre string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Hola, %s! (%d)\n", nombre, i)
		time.Sleep(time.Millisecond * 500) // Simula una tarea lenta
	}
}

func saludar2(nombre string, ch chan string) {
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Hola, %s! (%d)", nombre, i)
		ch <- msg // Enviar mensaje al canal
		time.Sleep(time.Millisecond * 500)
	}
	close(ch) // Cerramos el canal cuando terminamos de enviar datos
}

// Función que simula el procesamiento de un pedido
func procesarPedido(id int, ch chan string) {
	fmt.Printf("📦 Procesando pedido #%d...\n", id)
	time.Sleep(2 * time.Second)                      // Simula tiempo de procesamiento
	ch <- fmt.Sprintf("✅ Pedido #%d completado", id) // Enviar resultado al canal
}

func main() {
	/*
		go saludar("Cesar") // Ejecuta la función como una goroutine
		go saludar("Juan")  // Segunda goroutine

		// Esperamos un poco para que las goroutines terminen (esto no es ideal, pero funciona para la demo)
		time.Sleep(time.Second * 3)
		fmt.Println("Fin del programa")
*/
/*
	//CHANNEL
	ch := make(chan string) // Creamos un canal de tipo string

	go saludar2("Cesar", ch) // Llamamos a la goroutine y pasamos el canal

	// Recibimos datos del canal hasta que se cierre
	for mensaje := range ch {
		fmt.Println(mensaje)
	}
	fmt.Println("Fin del programa")
*/
/*
	//PROCESO PEDIDOS
	numPedidos := 5
	ch := make(chan string, numPedidos) // Canal con buffer para 5 pedidos

	// Iniciamos 5 goroutines para procesar los pedidos en paralelo
	for i := 1; i <= numPedidos; i++ {
		go procesarPedido(i, ch)
	}

	// Recibimos los resultados desde el canal
	for i := 1; i <= numPedidos; i++ {
		fmt.Println(<-ch) // Imprimimos cada pedido completado
	}

	fmt.Println("🎉 ¡Todos los pedidos han sido procesados!")

}
*/
/*
//RECURSIVIDAD
func factorial(n int) int {
	// Caso base: si n es 0 o 1, el factorial es 1
	if n == 0 || n == 1 {
		return 1
	}
	// Llamada recursiva: n * factorial(n-1)
	return n * factorial(n-1)
}

// Función recursiva para calcular Fibonacci
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	// Llamada recursiva: sumamos los dos términos anteriores
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	numero := 5
	resultado := factorial(numero)
	fmt.Printf("El factorial de %d es %d\n", numero, resultado)
	numero = 10
	fmt.Printf("El término %d de la secuencia de Fibonacci es %d\n", numero, fibonacci(numero))
}
*/
// Función que maneja la recuperación de un pánico
func manejarPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recuperación en proceso. Error detectado:", r)
	}
}

func dividir(a, b int) int {
	defer manejarPanic() // Se ejecutará antes de salir de la función si hay un panic

	if b == 0 {
		panic("¡Error! No se puede dividir por cero") // Lanza un pánico
	}

	return a / b
}

func main() {
	fmt.Println("Inicio del programa")

	resultado := dividir(10, 2)
	fmt.Println("Resultado de la división:", resultado)

	resultado = dividir(5, 0) // Provoca un pánico, pero se recupera
	fmt.Println("Esto se ejecuta después del panic")

	fmt.Println("Fin del programa")
}
