package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Estructura de contactos
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guardar contactos en un archivo JSON
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Formatear el JSON para que sea más legible
	return encoder.Encode(contacts)
}

// Cargar contactos desde un archivo JSON
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Si el archivo no existe, no hay contactos para cargar
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(contacts)
}

func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes
	if err := loadContactsFromFile(&contacts); err != nil {
		fmt.Println("Error al cargar contactos:", err)
		return
	}

	// Instancia para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar menú de opciones
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opción: ")

		// Leer la opción del usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Opción inválida:", err)
			continue
		}

		// Manejar la opción seleccionada
		switch option {
		case 1:
			// Ingresar y crear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			c.Name = strings.TrimSpace(c.Name)

			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			c.Email = strings.TrimSpace(c.Email)

			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')
			c.Phone = strings.TrimSpace(c.Phone)

			// Agregar contacto al slice
			contacts = append(contacts, c)

			// Guardar en archivo JSON
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			} else {
				fmt.Println("Contacto guardado exitosamente.")
			}

		case 2:
			// Mostrar todos los contactos
			fmt.Println("====================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s\n   Email: %s\n   Teléfono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("====================")

		case 3:
			// Salir del programa
			fmt.Println("Saliendo...")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
