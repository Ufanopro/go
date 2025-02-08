package main

import (
	"errors"
	"fmt"
)

// Definimos la estructura de la pila
type Stack struct {
	items []int
}

// Método para agregar un elemento a la pila (Push)
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Método para eliminar y devolver el elemento superior de la pila (Pop)
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Método para ver el elemento superior de la pila sin eliminarlo (Peek)
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Método para verificar si la pila está vacía
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Método para obtener el tamaño de la pila
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := Stack{}

	// Agregamos elementos a la pila
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Mostramos el tamaño de la pila
	fmt.Println("Tamaño de la pila:", stack.Size())

	// Vemos el elemento superior sin eliminarlo
	top, err := stack.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Elemento superior:", top)
	}

	// Eliminamos elementos de la pila
	for !stack.IsEmpty() {
		item, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Elemento eliminado:", item)
		}
	}

	// Verificamos si la pila está vacía
	fmt.Println("La pila está vacía:", stack.IsEmpty())
}
