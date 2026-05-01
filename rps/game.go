package rps

import (
	"math/rand"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 =0
	PAPER    = 1
	SCISSORS = 2

	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

// type
type Round struct {
	Winner         int
	ComputerChoice string
	RoundResult    string
}

// returns game result, computer choice, roundresult
func PlayRound(playerValue int) Round {
	// get random values for computer choice
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0
	var result Round
	// check what computer choses

	switch computerValue {
	case ROCK:
		computerChoice = " Computer chose ROCK"
	case PAPER:
		computerChoice = " Computer chose PAPER"
	case SCISSORS:
		computerChoice = " Computer chose SCISSORS"
	default:
		return result
	}
	// CHECK WINNER
	if playerValue == computerValue {
		roundResult = "it's a draw"
		winner = DRAW // 3
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS // 1
	} else {
		roundResult = "Computer wins!"
		winner = COMPUTERWINS // 2
	}
	result = Round{Winner: winner, ComputerChoice: computerChoice, RoundResult: roundResult}
	return result
}
