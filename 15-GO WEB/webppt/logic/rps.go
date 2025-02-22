package rps

import (
	"math/rand"
)

// Constantes de las opciones
const (
	PIEDRA = 0
	PAPEL  = 1
	TIJERA = 2
)

// Estructura del resultado de la ronda
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     int    `json:"computer_score"`
	PlayerScore       int    `json:"player_score"`
}

// Mensajes de ganador, perdedor y empate
var winMessages = []string{"¡Bien hecho!", "¡Genial!", "¡Estás de suerte!"}
var loseMessages = []string{"¡Lástima!", "¡La próxima vez!", "¡Hoy no es tu día!"}
var drawMessages = []string{"¡Iguales!", "¡Casi, estáis conectados!", "¡Empate!"}

// Variables de puntuación
var ComputerScore, PlayerScore int

// Función para jugar una ronda
func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	computerChoice := ""

	// Determinar la elección de la computadora
	switch computerValue {
	case PIEDRA:
		computerChoice = "La computadora eligió PIEDRA"
	case PAPEL:
		computerChoice = "La computadora eligió PAPEL"
	case TIJERA:
		computerChoice = "La computadora eligió TIJERA"
	}

	// Generar mensaje aleatorio
	messageInt := rand.Intn(3)
	message := ""
	roundResult := ""

	// Lógica del juego
	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	// Retornar resultado
	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerValue,
		ComputerScore:     ComputerScore,
		PlayerScore:       PlayerScore,
	}
}
