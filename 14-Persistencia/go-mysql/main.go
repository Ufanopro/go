package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	database "mysql-go/db"
	handlers "mysql-go/handler"
	"mysql-go/models"

	//IMPORTANTE añadir el link
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Establecer la conexion a la DB
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	/*
					// Listamos all Contacts
					handlers.ListContacts(db)

					// Listamos por ID
					contactId := 5
					handlers.GetContactByID(db, contactId)

				// Creamos un Contacto
				newContact := models.Contact{
					Name:  "Cesar Diaz",
					Email: "cesar@example.com",
					Phone: "123456789",
				}

				// Registrar el contacto
				handlers.CreateContact(db, newContact)

			// Crear una instancia de Contact con los detalles actualizados
			updatedContact := models.Contact{
				Id:    5, // ID del contacto que deseas actualizar
				Name:  "Isabel Blazquez",
				Email: "isabel@example.com",
				Phone: "987654321",
			}

			// Actualizar el contacto en la base de datos
			handlers.UpdateContact(db, updatedContact)
		// ID del contacto que deseas eliminar
		contactID := 6

		// Eliminar el contacto de la base de datos
		handlers.DeleteContact(db, contactID)

		// Listar contactos (opcional, solo para verificar que el contacto fue eliminado)
		handlers.ListContacts(db)
	*/

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opción seleccionada por el usuario
		var option int
		fmt.Scanln(&option)

		// Ejecutar la opción seleccionada
		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese el ID del contacto: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactByID(db, idContact)
		case 3:
			newContact := inputContactDetails()
			handlers.CreateContact(db, newContact)
		case 4:
			updatedContact := inputContactDetails()
			handlers.UpdateContact(db, updatedContact)
		case 5:
			fmt.Print("Ingrese el ID del contacto que quiere eliminar: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}

// Función para ingresar los detalles del contacto desde la entrada estándar
func inputContactDetails() models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electrónico del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el número de teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact

}
