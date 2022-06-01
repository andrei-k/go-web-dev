package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // Beats scissors (scissors + 1) % 3 = 0
	PAPER    = 1 // Beats rock (rock + 1) % 3 = 1
	SCISSORS = 2 // Beats paper (paper + 1) % 3 = 2
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

var winMessages = []string{
	"Good job, you won!",
	"You won, you're a rock star!",
	"Great work!",
}

var loseMessages = []string{
	"You lost, better luck next time!",
	"Too bad!",
	"This is just not your day!",
}

var drawMessages = []string{
	"Great minds think alike.",
	"Uh oh, try again.",
	"Nobody wins, but you can try again.",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	}

	messageInt := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw."
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "You won!"
		message = winMessages[messageInt]
	} else {
		roundResult = "You lost!"
		message = loseMessages[messageInt]
	}

	// var result Round
	// result.Message = message
	// result.ComputerChoice = computerChoice
	// result.RoundResult = roundResult

	// Same thing as above but in one line
	result := Round{message, computerChoice, roundResult}

	return result
}
