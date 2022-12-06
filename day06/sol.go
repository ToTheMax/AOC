package main

import (
	"fmt"
	"os"
)

func findMarker(input string, distinctChars int) int {
	for i := range input {
		seen := make(map[rune]bool)
		for j := 0; j < distinctChars; j++ {
			seen[rune(input[i+j])] = true
		}
		if len(seen) == distinctChars {
			return i + distinctChars
		}
	}
	return 0
}

func main() {
	input, _ := os.ReadFile("in.txt")
	line := string(input)
	fmt.Println("Sol1:", findMarker(line, 4))
	fmt.Println("Sol2:", findMarker(line, 14))
}
