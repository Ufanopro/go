package main

import (
	"library/animal"
)

func main() {
	/*
			myBook := book.NewBook("Moby Dick", "Herman Melville", 300)
				var myBook = book.Book{
				Title:  "Moby Dick",
				Author: "Herman Melville",
				Pages:  300,
			}

			myBook.PrintInfo()
			myBook.SetTitle("La ballena blanca")
			fmt.Println(myBook.GetTitle())

			myTextBook := book.NewTextBook("El Quijote", "Cervanes", 1000,
				"Edube", "Secundaria")

			myTextBook.PrintInfo()

				myBook.PrintInfo()
				myBook.PrintInfo()

			book.Print(myBook)
			book.Print(myTextBook)


		miPerro := animal.Perro{Nombre: "Max"}
		miGato := animal.Gato{Nombre: "Tom"}

		animal.HacerSonido(&miPerro)
		animal.HacerSonido(&miGato)
	*/
	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
