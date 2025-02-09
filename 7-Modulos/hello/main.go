package main

import (
	"fmt"
	"log"

	"github.com/Ufanopro/go/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Cesar", "Alex", "Isabel"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	/*message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	*/
	fmt.Println(messages)
}
