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
		switch me {
		case "X":
			// rock
			scoreShape = 1
		case "Y":
			// paper
			scoreShape = 2
		case "Z":
			// scissor
			scoreShape = 3
		}

		var scoreOutcome int
		switch {
		case op == "A" && me == "Z", op == "B" && me == "X", op == "C" && me == "Y":
			// lost
			scoreOutcome = 0
		case op == "A" && me == "X", op == "B" && me == "Y", op == "C" && me == "Z":
			// draw
			scoreOutcome = 3
		case op == "A" && me == "Y", op == "B" && me == "Z", op == "C" && me == "X":
			// win
			scoreOutcome = 6
		}

		scoreRound := scoreShape + scoreOutcome
		scoreTotal += scoreRound

		fmt.Printf("Opponent: %s\tMe: %s\tShape Score: %d\tHand Score: %d\tRound Score: %d\n", op, me, scoreShape, scoreOutcome, scoreRound)
	}
	fmt.Printf("Score Total: %d\n", scoreTotal)
}
