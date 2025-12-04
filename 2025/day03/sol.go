package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMaxJoltage(bank string, numDigits int) int {
	result := ""
	start := 0
	for i := 0; i < numDigits; i++ {
		maxIdx, maxDigit := start, bank[start]
		for j := start; j <= len(bank)-(numDigits-i); j++ {
			if bank[j] > maxDigit {
				maxDigit, maxIdx = bank[j], j
			}
		}
		result += string(maxDigit)
		start = maxIdx + 1
	}
	max, _ := strconv.Atoi(result)
	return max
}

func main() {
	input, _ := os.ReadFile("in.txt")
	banks := strings.Split(string(input), "\n")

	score1 := 0
	score2 := 0
	for _, bank := range banks {
		max := findMaxJoltage(bank, 2)
		score1 += max

		max2 := findMaxJoltage(bank, 12)
		score2 += max2
	}

	fmt.Println("Sol 1:", score1)
	fmt.Println("Sol 2:", score2)
}
