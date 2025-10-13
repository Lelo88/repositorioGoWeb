package rockpaperscissors

import (
	"math"
	"math/rand"
	"strconv"
)


const (
	ROCK    = 0 // Representa piedra. La logica se manejará con módulos.
	PAPER   = 1 // Representa papel.
	SCISSORS = 2 // Representa tijeras.
)

type Round struct {
	Message string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult string `json:"round_result"`
	ComputerChoiceInt int `json:"computer_choice_int"`
	ComputerScore string `json:"computer_score"`
	PlayerScore string `json:"player_score"`
}

var winMessages = []string{
	"Bien hecho!",
	"¡Ganaste esta ronda!",
	"¡Eres el ganador!",
}

var loseMessages = []string{
	"¡Oh no! La computadora gana esta ronda.",
	"¡La computadora te ha vencido!",
	"¡La computadora es la ganadora!",
}

var drawMessages = []string{
	"¡Es un empate! Inténtalo de nuevo.",
	"¡Ninguno gana esta ronda!",
	"¡Empate! Sigue jugando.",
}

var computerScore, playerScore int

func PlayRound(playerValue int) Round {
	computerValue := int(math.Floor(3 * rand.Float64()))

	var result Round
	
	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió piedra."
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió papel."
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió tijeras."
	}

	messageInt := rand.Intn(3) // 0, 1, 2

	var message string
	if playerValue == computerValue {
		roundResult = "Empate"
		message = drawMessages[messageInt]
	} else if (playerValue-computerValue+3)%3 == 1 {
		roundResult = "Ganaste"
		playerScore++
		message = winMessages[messageInt]
	} else {
		roundResult = "Perdiste"
		computerScore++
		message = loseMessages[messageInt]
	}

	result = Round{
		Message: message,
		ComputerChoice: computerChoice,
		RoundResult: roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:  strconv.Itoa(computerScore),
		PlayerScore: strconv.Itoa(playerScore),
	}

	return result
}