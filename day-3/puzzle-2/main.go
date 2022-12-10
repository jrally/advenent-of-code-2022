package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var total int
	var value int
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	groupSize := 3
	var groups [][]string
	for i := 0; i < len(lines); i += groupSize {
		end := i + groupSize

		if end > len(lines) {
			end = len(lines)
		}

		groups = append(groups, lines[i:end])

	}
	var pack1 string
	var pack2 string
	var pack3 string

	for _, group := range groups {
		for num, pack := range group {

			switch num {
			case 0:
				pack1 = pack
			case 1:
				pack2 = pack
			case 2:
				pack3 = pack
			}
		}

	out:
		for _, firstItem := range pack1 {
			for _, secondItem := range pack2 {
				for _, thirdItem := range pack3 {
					if firstItem == secondItem && firstItem == thirdItem {
						if firstItem >= 97 && firstItem <= 122 {
							value = int(firstItem) - 96
						} else {
							value = int(firstItem) - 38
						}
						total += value
						fmt.Printf("Char: %q\tValue: %d\n", firstItem, value)
						break out
					}
				}
			}
		}

		fmt.Printf("Pack 1: %s\tPack 2: %s\tPack 3: %s\n", pack1, pack2, pack3)

	}
	fmt.Printf("Total: %d\n", total)

}
