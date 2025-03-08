package main

import (
	"bufio"
	"clientes/conexion"
	"clientes/handlers"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Cargar las variables de entorno
	err := conexion.LoadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	// Crear el string de conexión
	dsn := conexion.CreateDSN()

	// Conectar a la base de datos
	db, err := conexion.ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión
	err = conexion.PingDB(db)
	if err != nil {
		log.Fatal(err)
	}

	// Mensaje de éxito
	fmt.Println("✅ Conexión exitosa a la base de datos.")
	// Obtener todos los clientes
	clientes, err := handlers.AllClientes()
	if err != nil {
		log.Fatalf("Error obteniendo clientes: %v", err)
	}
	// Mostrar todos los clientes
	for _, c := range clientes {
		handlers.MostrarCliente(c)
	}

	//MENU CONSOLA
	// Crear un lector de entrada
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar el menú ASCII
		fmt.Println(`
		--- Menú de Clientes ---
		+--------------------------+
		| 1. Listar todos los clientes |
		| 2. Buscar cliente por ID    |
		| 3. Crear un nuevo cliente   |
		| 4. Actualizar un cliente    |
		| 5. Eliminar un cliente     |
		| 6. Salir                   |
		+--------------------------+
		`)

		// Leer la opción del usuario
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			// Listar todos los clientes
			clientes, err := handlers.AllClientes()
			if err != nil {
				log.Printf("Error obteniendo clientes: %v", err)
			} else {
				for _, c := range clientes {
					handlers.MostrarCliente(c)
				}
			}

		case "2":
			// Buscar cliente por ID
			fmt.Print("Ingrese el ID del cliente: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("ID inválido: %v", err)
			} else {
				cliente, err := handlers.ByIdCliente(id)
				if err != nil {
					log.Printf("Error obteniendo cliente: %v", err)
				} else {
					handlers.MostrarCliente(*cliente)
				}
			}

		case "3":
			// Crear un nuevo cliente
			fmt.Print("Ingrese el nombre: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Ingrese el correo: ")
			correo, _ := reader.ReadString('\n')
			correo = strings.TrimSpace(correo)

			fmt.Print("Ingrese el teléfono: ")
			telefono, _ := reader.ReadString('\n')
			telefono = strings.TrimSpace(telefono)

			fmt.Print("Ingrese la fecha (YYYY-MM-DD HH:MM:SS): ")
			fecha, _ := reader.ReadString('\n')
			fecha = strings.TrimSpace(fecha)

			clienteCreado, err := handlers.CrearCliente(nombre, correo, telefono, fecha)
			if err != nil {
				log.Printf("Error creando cliente: %v", err)
			} else {
				fmt.Println("✅ Cliente creado exitosamente.")
				handlers.MostrarCliente(*clienteCreado)
			}

		case "4":
			// Actualizar un cliente
			fmt.Print("Ingrese el ID del cliente a actualizar: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("ID inválido: %v", err)
			} else {
				fmt.Print("Ingrese el nuevo nombre: ")
				nombre, _ := reader.ReadString('\n')
				nombre = strings.TrimSpace(nombre)

				fmt.Print("Ingrese el nuevo correo: ")
				correo, _ := reader.ReadString('\n')
				correo = strings.TrimSpace(correo)

				fmt.Print("Ingrese el nuevo teléfono: ")
				telefono, _ := reader.ReadString('\n')
				telefono = strings.TrimSpace(telefono)

				fmt.Print("Ingrese la nueva fecha (YYYY-MM-DD HH:MM:SS): ")
				fecha, _ := reader.ReadString('\n')
				fecha = strings.TrimSpace(fecha)

				clienteActualizado, err := handlers.UpdateCliente(id, nombre, correo, telefono, fecha)
				if err != nil {
					log.Printf("Error actualizando cliente: %v", err)
				} else {
					fmt.Println("✅ Cliente actualizado exitosamente.")
					handlers.MostrarCliente(*clienteActualizado)
				}
			}

		case "5":
			// Eliminar un cliente
			fmt.Print("Ingrese el ID del cliente a eliminar: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("ID inválido: %v", err)
			} else {
				err := handlers.DeleteCliente(id)
				if err != nil {
					log.Printf("Error eliminando cliente: %v", err)
				} else {
					fmt.Println("✅ Cliente eliminado exitosamente.")
				}
			}

		case "6":
			// Salir del programa
			fmt.Println("Saliendo...")
			return

		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}
