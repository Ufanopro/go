package main

import (
	"errors"
	"fmt"
)

// Definimos la estructura de la cola
type Queue struct {
	items []int
}

// Método para agregar un elemento a la cola (Enqueue)
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Método para eliminar y devolver el primer elemento de la cola (Dequeue)
func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("la cola está vacía")
	}
	item := q.items[0]    // Obtenemos el primer elemento
	q.items = q.items[1:] // Eliminamos el primer elemento
	return item, nil
}

// Método para ver el primer elemento de la cola sin eliminarlo (Peek)
func (q *Queue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("la cola está vacía")
	}
	return q.items[0], nil
}

// Método para verificar si la cola está vacía
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Método para obtener el tamaño de la cola
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := Queue{}

	// Agregamos elementos a la cola
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Mostramos el tamaño de la cola
	fmt.Println("Tamaño de la cola:", queue.Size())

	// Vemos el primer elemento sin eliminarlo
	front, err := queue.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Primer elemento de la cola:", front)
	}

	// Eliminamos elementos de la cola
	for !queue.IsEmpty() {
		item, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Elemento eliminado de la cola:", item)
		}
	}

	// Verificamos si la cola está vacía
	fmt.Println("La cola está vacía:", queue.IsEmpty())
}
