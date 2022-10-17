package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
	//PLAYERWINS   = 1
	//COMPUTERWINS = 2
	//DRAW         = 3
)

type Round struct {
	//Winner int `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket!",
}

var looseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day..",
}

var drawMessages = []string{
	"Great minds think alike",
	"Uh oh! Try again",
	"Never give up!",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	//winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	messageInt := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw!"
		//winner = DRAW
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins"
		//winner = PLAYERWINS
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins"
		//winner = COMPUTERWINS
		message = looseMessages[messageInt]
	}

	var result Round
	result.Message = message
	//result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
