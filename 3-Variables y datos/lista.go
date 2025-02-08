package main

import (
	"fmt"
)

// Definimos la estructura de un nodo
type Node struct {
	Value int
	Next  *Node
}

// Definimos la estructura de la lista enlazada
type LinkedList struct {
	Head *Node
}

// Método para agregar un elemento al final de la lista
func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}

	if l.Head == nil {
		// Si la lista está vacía, el nuevo nodo es la cabeza
		l.Head = newNode
	} else {
		// Recorremos la lista hasta llegar al último nodo
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		// Agregamos el nuevo nodo al final
		current.Next = newNode
	}
}

// Método para imprimir la lista enlazada
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// Método para buscar un valor en la lista
func (l *LinkedList) Search(value int) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Método para eliminar un valor de la lista
func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		return // La lista está vacía
	}

	// Si el valor a eliminar es la cabeza
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	// Buscamos el nodo a eliminar
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			// Eliminamos el nodo saltándolo
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func main() {
	// Creamos una lista enlazada vacía
	list := LinkedList{}

	// Agregamos elementos a la lista
	list.Append(10)
	list.Append(20)
	list.Append(30)

	// Imprimimos la lista
	fmt.Println("Lista enlazada:")
	list.Print()

	// Buscamos un valor en la lista
	fmt.Println("¿El valor 20 está en la lista?", list.Search(20))
	fmt.Println("¿El valor 40 está en la lista?", list.Search(40))

	// Eliminamos un valor de la lista
	fmt.Println("Eliminando el valor 20...")
	list.Delete(20)
	list.Print()

	// Eliminamos la cabeza de la lista
	fmt.Println("Eliminando el valor 10...")
	list.Delete(10)
	list.Print()

	// Intentamos eliminar un valor que no existe
	fmt.Println("Eliminando el valor 50...")
	list.Delete(50)
	list.Print()
}
