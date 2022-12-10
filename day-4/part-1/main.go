package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SplitConvert(section string) []int {
	var ints []int
	split := strings.Split(section, "-")
	for _, value := range split {
		convert, _ := strconv.Atoi(value)
		ints = append(ints, convert)
	}

	return ints
}

func Expand(start, end int) []int {
	var expanded []int
	for i := start; i <= end; i++ {
		expanded = append(expanded, i)
	}

	return expanded
}

func CheckContents(section []int, start, end int) bool {
	var contains int
	for _, num := range section {
		if num == start {
			contains++
		}
		if num == end {
			contains++
		}
	}

	if contains == 2 {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")

		firstSect := SplitConvert(pair[0])
		secondSect := SplitConvert(pair[1])

		expandFirst := Expand(firstSect[0], firstSect[1])
		expandSecond := Expand(secondSect[0], secondSect[1])

		if firstSect[0] == secondSect[0] && firstSect[1] == secondSect[1] {
			count++
			continue
		}

		if ok := CheckContents(expandFirst, secondSect[0], secondSect[1]); ok {
			count++
		}

		if ok := CheckContents(expandSecond, firstSect[0], firstSect[1]); ok {
			count++
		}

	}

	fmt.Printf("Part 1 Count: %d\n", count)

}
