package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		fmt.Println("==== Nueva Ronda ====")
		fmt.Printf("Jugador eligió: %d\n", i)
		fmt.Printf("Computadora eligió: %s\n", round.ComputerChoice)
		fmt.Printf("Resultado de la ronda: %s\n", round.RoundResult)
		fmt.Printf("Mensaje: %s\n", round.Message)
		fmt.Printf("Elección de la computadora (int): %d\n", round.ComputerChoiceInt)
		fmt.Printf("Puntuación Computadora: %d\n", round.ComputerScore)
		fmt.Printf("Puntuación Jugador: %d\n", round.PlayerScore)
		fmt.Println("==============================")
	}
}
