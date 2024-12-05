package main

import (
	"fmt"
	"os"
	"strings"
)

func inBounds(lines []string, x, y int) bool {
	return x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines)
}

func searchXmas(lines []string, x, y int) int {

	// Check all directions
	xmas_checks := [][][]int{
		{{x, y}, {x, y + 1}, {x, y + 2}, {x, y + 3}},             // Horizontal
		{{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y}},             // Vertical
		{{x, y}, {x + 1, y + 1}, {x + 2, y + 2}, {x + 3, y + 3}}, // Descending diagonal
		{{x, y}, {x - 1, y + 1}, {x - 2, y + 2}, {x - 3, y + 3}}, // Ascending diagonal
	}
	count := 0
	for _, check := range xmas_checks {
		word := ""
		for _, coord := range check {
			if !inBounds(lines, coord[0], coord[1]) {
				break
			}
			word += string(lines[coord[1]][coord[0]])
		}
		if word == "XMAS" || word == "SAMX" {
			count++
		}
	}
	return count
}

func searchMas(lines []string, x, y int) int {

	// Check if there is A in the center
	if y < len(lines)-1 && x < len(lines[0])-1 && lines[y+1][x+1] != 'A' {
		return 0
	}

	// Check if there are M and S in the corners
	M_count := 0
	S_count := 0
	mas_check := [][]int{{x, y}, {x + 2, y}, {x, y + 2}, {x + 2, y + 2}}
	for _, coord := range mas_check {
		if !inBounds(lines, coord[0], coord[1]) {
			return 0
		}
		if lines[coord[1]][coord[0]] == 'M' {
			M_count++
		} else if lines[coord[1]][coord[0]] == 'S' {
			S_count++
		}
	}
	if M_count == 2 && S_count == 2 {
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
	fmt.Println(countP2) // 1900 was too high
}
