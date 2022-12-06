package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createStack(lines []string) [][]rune {
	indexLine := lines[len(lines)-1]
	stacks := make([][]rune, len(strings.Split(indexLine, "   ")))
	for index, char := range indexLine {
		if char != ' ' {
			stackNumber := int(char - '0')
			var stack []rune
			for j := len(lines) - 2; j >= 0; j-- {
				if lines[j][index] != ' ' {
					stack = append(stack, rune(lines[j][index]))
				}
			}
			stacks[stackNumber-1] = stack
		}
	}
	return stacks
}

func moveCrate(stacks [][]rune, amount int, from int, to int) {
	index := len(stacks[from-1]) - amount
	crates := stacks[from-1][index:]
	stacks[to-1] = append(stacks[to-1], crates...)
	stacks[from-1] = stacks[from-1][:index]
}

func readStackTops(stacks [][]rune) string {
	result := ""
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}
	return result
}

func main() {

	input, _ := os.ReadFile("in.txt")
	text := string(input)
	lines := strings.Split(text, "\n")

	// Read stacks
	lineNumber := 0
	for lines[lineNumber] != "" {
		lineNumber++
	}

	// Initialize stacks
	stacksP1 := createStack(lines[:lineNumber])
	stacksP2 := createStack(lines[:lineNumber])

	// Read actions
	for _, line := range lines[lineNumber+1:] {
		splitted := strings.Split(line, " ")
		amount, _ := strconv.Atoi(splitted[1])
		from, _ := strconv.Atoi(splitted[3])
		to, _ := strconv.Atoi(splitted[5])

		// Part 1
		for i := 0; i < amount; i++ {
			moveCrate(stacksP1, 1, from, to)
		}
		// Part 2
		moveCrate(stacksP2, amount, from, to)
	}

	// Print top of stacks
	fmt.Println("Sol1: ", readStackTops(stacksP1))
	fmt.Println("Sol1: ", readStackTops(stacksP2))
}
