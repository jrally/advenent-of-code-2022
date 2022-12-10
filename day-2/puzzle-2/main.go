package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var scoreTotal int
	for scanner.Scan() {
		round := scanner.Text()

		player := strings.Split(round, " ")
		op := player[0]
		me := player[1]

		var scoreShape int
		var scoreOutcome int
		switch me {
		case "X":
			// lost
			scoreOutcome = 0

			switch op {
			case "A":
				// rock
				// scissor score
				scoreShape = 3
			case "B":
				// paper
				// rock score
				scoreShape = 1
			case "C":
				// scissor
				// paper score
				scoreShape = 2
			}

		case "Y":
			// draw
			scoreOutcome = 3

			switch op {
			case "A":
				// rock
				// rock score
				scoreShape = 1
			case "B":
				// paper
				// paper score
				scoreShape = 2
			case "C":
				// scissor
				// scissor score
				scoreShape = 3
			}

		case "Z":
			// win
			scoreOutcome = 6

			switch op {
			case "A":
				// rock
				// paper score
				scoreShape = 2
			case "B":
				// paper
				// scissor score
				scoreShape = 3
			case "C":
				// scissor
				// rock score
				scoreShape = 1
			}

		}

		scoreRound := scoreShape + scoreOutcome
		scoreTotal += scoreRound

		fmt.Printf("Opponent: %s\tMe: %s\tShape Score: %d\tHand Score: %d\tRound Score: %d\n", op, me, scoreShape, scoreOutcome, scoreRound)
	}
	fmt.Printf("Score Total: %d\n", scoreTotal)
}
