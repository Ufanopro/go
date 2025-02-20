package main

import "fmt"

func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func main() {
	PrintList("Cesar", "Isabel", "Alex")
	PrintList(100, 456, 789, 658)
	PrintList("Hola", 452, 4.30, true)
}
