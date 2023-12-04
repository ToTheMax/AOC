package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(strings.ReplaceAll(string(input), "  ", " "), "\n")
	totalPoints := 0

	scratchCards := make([]int, len(lines))
	for i := range scratchCards {
		scratchCards[i] = 1
	}

	for i, line := range lines {
		splittedNumbers := strings.Split(strings.Split(line, ":")[1], " | ")

		// Fill winning numbers
		winningNumbers := map[int]int{}
		for _, number := range strings.Split(splittedNumbers[0], " ") {
			number, _ := strconv.Atoi(number)
			winningNumbers[number] = number
		}

		// Numbers you have
		matchingNumbers := 0
		for _, number := range strings.Split(splittedNumbers[1], " ") {
			number, _ := strconv.Atoi(number)
			if _, ok := winningNumbers[number]; ok {
				matchingNumbers++
			}
		}

		// Add points
		if matchingNumbers > 0 {
			totalPoints += int(math.Pow(2, float64(matchingNumbers-1)))
		}

		// Add scratchcards
		for j := 0; j < matchingNumbers; j++ {
			scratchCards[i+1+j] += scratchCards[i]
		}
	}
	fmt.Println("Sol 1:", totalPoints)

	scratchCardCount := 0
	for i := range scratchCards {
		scratchCardCount += scratchCards[i]
	}
	fmt.Println("Sol 2:", scratchCardCount)
}
