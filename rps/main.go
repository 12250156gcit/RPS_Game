package rps

import (
	"math/rand"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 =0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

// returns game result, computer choice, roundresult
func PlayRound(playerValue int) (int, string, string) {
	// get random values for computer choice
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundresult := ""
	winner := 0
	// check what computer choses

	switch computerValue {
	case ROCK:
		computerChoice = " Computer chose ROCK"
	case PAPER:
		computerChoice = " Computer chose PAPER"
	case SCISSORS:
		computerChoice = " Computer chose SCISSORS"
	default:
		return winner, " Invalid computer choice", "Invalid round"
	}
	// CHECK WINNER
	if playerValue == computerValue {
		roundresult = "it's a draw"
		winner = DRAW // 3
	} else if playerValue == (computerValue+1)%3 {
		roundresult = "Player wins!"
		winner = PLAYERWINS // 1
	} else {
		roundresult = "Computer wins!"
		winner = COMPUTERWINS // 2
	}
	return winner, computerChoice, roundresult
}
