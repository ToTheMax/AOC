package main

import (
	"fmt"
	"os"
	"strings"
)

func searchXmas(lines []string, x, y int) int {
	// Check all directions
	words := []string{}
	if x+3 < len(lines[y]) {
		words = append(words, fmt.Sprintf("%c%c%c%c", lines[y][x], lines[y][x+1], lines[y][x+2], lines[y][x+3]))
	}
	if y+3 < len(lines) {
		words = append(words, fmt.Sprintf("%c%c%c%c", lines[y][x], lines[y+1][x], lines[y+2][x], lines[y+3][x]))
	}
	if y+3 < len(lines) && x+3 < len(lines[y]) {
		words = append(words, fmt.Sprintf("%c%c%c%c", lines[y][x], lines[y+1][x+1], lines[y+2][x+2], lines[y+3][x+3]))
	}
	if y+3 < len(lines) && x-3 >= 0 {
		words = append(words, fmt.Sprintf("%c%c%c%c", lines[y][x], lines[y+1][x-1], lines[y+2][x-2], lines[y+3][x-3]))
	}
	count := 0
	for _, word := range words {
		if word == "XMAS" || word == "SAMX" {
			count++
		}
	}
	return count
}

func searchMas(lines []string, x, y int) int {
	// Check bounds and center A
	if x > len(lines[0])-3 || y > len(lines)-3 {
		return 0
	}

	// Check MAS in diagonals
	word1 := fmt.Sprintf("%c%c%c", lines[y][x], lines[y+1][x+1], lines[y+2][x+2])
	word2 := fmt.Sprintf("%c%c%c", lines[y][x+2], lines[y+1][x+1], lines[y+2][x])

	if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
		return 1
	}
	return 0
}

func main() {
	// Read input as single string
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	countP1 := 0
	countP2 := 0
	for y := range lines {
		for x := range lines[y] {
			countP1 += searchXmas(lines, x, y)
			countP2 += searchMas(lines, x, y)
		}
	}

	fmt.Println(countP1)
	fmt.Println(countP2)
}
