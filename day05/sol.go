package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.ReadFile("in.txt")
	text := string(input)
	lines := strings.Split(text, "\n")

	// Read stacks
	lineNumber := 0
	for lines[lineNumber] != "" {
		lineNumber++
	}
	indexLineNumber := lineNumber - 1
	indexLine := lines[indexLineNumber]

	// Initialize stacks
	stacksP1 := make([][]rune, len(strings.Split(indexLine, "   ")))
	stacksP2 := make([][]rune, len(stacksP1))
	for index, char := range indexLine {
		if char != ' ' {
			stackNumber := int(char - '0')
			var stack []rune
			j := indexLineNumber - 1
			for j >= 0 && lines[j][index] != ' ' {
				stack = append(stack, rune(lines[j][index]))
				j--
			}
			stacksP1[stackNumber-1] = stack
			stacksP2[stackNumber-1] = stack
		}
	}

	// Read actions
	for _, line := range lines[indexLineNumber+2:] {
		splitted := strings.Split(line, " ")
		amount, _ := strconv.Atoi(splitted[1])
		from, _ := strconv.Atoi(splitted[3])
		to, _ := strconv.Atoi(splitted[5])

		// Move containers
		// Part 1
		for i := 0; i < amount; i++ {
			index := len(stacksP1[from-1]) - 1
			crates := stacksP1[from-1][index:]
			stacksP1[to-1] = append(stacksP1[to-1], crates...)
			stacksP1[from-1] = stacksP1[from-1][:index]
		}
		// Part 2
		index := len(stacksP2[from-1]) - amount
		crates := stacksP2[from-1][index:]
		stacksP2[to-1] = append(stacksP2[to-1], crates...)
		stacksP2[from-1] = stacksP2[from-1][:index]
	}

	// Print top of stacks
	fmt.Print("Sol1: ")
	for _, stack := range stacksP1 {
		fmt.Print(string(stack[len(stack)-1]))
	}
	fmt.Print("\n")
	fmt.Print("Sol2: ")
	for _, stack := range stacksP2 {
		fmt.Print(string(stack[len(stack)-1]))
	}
	fmt.Print("\n")
}
