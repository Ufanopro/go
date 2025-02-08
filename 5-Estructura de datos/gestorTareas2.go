package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Método para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Método para marcar como completado
func (l *ListaTareas) marcarCompletado(index int) {
	if index >= 0 && index < len(l.tareas) {
		l.tareas[index].completado = true
	} else {
		fmt.Println("Índice fuera de rango")
	}
}

// Método para editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	if index >= 0 && index < len(l.tareas) {
		l.tareas[index] = t
	} else {
		fmt.Println("Índice fuera de rango")
	}
}

// Método para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	if index >= 0 && index < len(l.tareas) {
		l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
	} else {
		fmt.Println("Índice fuera de rango")
	}
}

func main() {
	// Instanciamos la Lista de tareas
	lista := ListaTareas{}

	// Instancia bufio para la entrada de datos
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar Tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar Tarea\n",
			"5. Salir\n")
		fmt.Print("Ingrese opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			t.nombre = strings.TrimSpace(t.nombre) // Eliminar el salto de línea
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			t.desc = strings.TrimSpace(t.desc) // Eliminar el salto de línea
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el índice de la tarea que quiere completar: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
		case 3:
			var index int
			fmt.Print("Ingrese el índice de la tarea que quiere editar: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(lista.tareas) {
				var t Tarea
				fmt.Print("Ingrese el nuevo nombre de la tarea: ")
				t.nombre, _ = leer.ReadString('\n')
				t.nombre = strings.TrimSpace(t.nombre)
				fmt.Print("Ingrese la nueva descripción de la tarea: ")
				t.desc, _ = leer.ReadString('\n')
				t.desc = strings.TrimSpace(t.desc)
				lista.editarTarea(index, t)
				fmt.Println("Tarea editada correctamente")
			} else {
				fmt.Println("Índice fuera de rango")
			}
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea que quiere eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del gestor de tareas...")
			return
		default:
			fmt.Println("Opción no válida")
		}

		// Mostrar lista de tareas
		fmt.Println("\nLista de tareas:")
		fmt.Println("====================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("====================\n")
	}
}
