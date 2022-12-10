package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	contents := string(file)

	convert := strings.ReplaceAll(contents, "\n", ",")

	split := strings.Split(convert, ",,")

	var sorted []int
	for _, calories := range split {
		items := strings.Split(calories, ",")

		var total int
		for _, item := range items {
			value, _ := strconv.Atoi(item)
			total += value
		}
		sorted = append(sorted, total)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	for _, cals := range sorted {
		fmt.Println(cals)
	}

	var topThree int
	for _, three := range sorted[len(sorted)-3:] {
		topThree += three
	}

	fmt.Println(topThree)
}
