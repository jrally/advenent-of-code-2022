package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Move(from, to []string) ([]string, []string) {
	copy := from[len(from)-1]
	to = append(to, copy)

	from = from[:len(from)-1]

	return from, to
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Panic(err)
	}

	// stack := [][]string{{"Z", "N"}, {"M", "C", "D"}, {"P"}}

	stack := [][]string{{"T", "P", "Z", "C", "S", "L", "Q", "N"}, {"L", "P", "T", "V", "H", "C", "G"}, {"D", "C", "Z", "F"}, {"G", "W", "T", "D", "L", "M", "V", "C"}, {"P", "W", "C"}, {"P", "F", "J", "D", "C", "T", "S", "Z"}, {"V", "W", "G", "B", "D"}, {"N", "J", "S", "Q", "H", "W"}, {"R", "C", "Q", "F", "S", "L", "V"}}

	inputs := strings.Split(string(file), "\n")

	for _, input := range inputs {
		action := strings.Split(input, " ")
		move, _ := strconv.Atoi(action[1])
		from, _ := strconv.Atoi(action[3])
		to, _ := strconv.Atoi(action[5])

		for i := 0; i < move; i++ {
			stack[from-1], stack[to-1] = Move(stack[from-1], stack[to-1])
		}
	}

	fmt.Printf("%s%s%s%s%s%s%s%s%s\n", stack[0][len(stack[0])-1], stack[1][len(stack[1])-1], stack[2][len(stack[2])-1], stack[3][len(stack[3])-1], stack[4][len(stack[4])-1], stack[5][len(stack[5])-1], stack[6][len(stack[6])-1], stack[7][len(stack[7])-1], stack[8][len(stack[8])-1])

}
