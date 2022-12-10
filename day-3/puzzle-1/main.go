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

	var total int
	for scanner.Scan() {
		line := scanner.Text()

		length := len(line)
		first := line[0 : length/2]
		second := line[length/2 : length]
		var value int

	out:
		for _, firstChar := range first {
			for _, secondChar := range second {
				if firstChar == secondChar {
					if firstChar >= 97 && firstChar <= 122 {
						value = int(firstChar) - 96
					} else {
						value = int(firstChar) - 38
					}
					total += value
					fmt.Printf("Line: %-48s\tFirst: %-24s\tSecond: %-24s\tChar: %q\tValue: %d\n", line, first, second, firstChar, value)
					break out
				}
			}
		}
	}
	fmt.Printf("Total: %d\n", total)

}
