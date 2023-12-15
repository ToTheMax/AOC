package main

import (
	"fmt"
	"os"
	"strings"
)

func hash(step string) int {
	curValue := 0
	for _, char := range step {
		curValue += int(char)
		curValue *= 17
		curValue %= 256
	}
	return curValue
}

func main() {
	input, _ := os.ReadFile("in.txt")
	steps := strings.Split(string(input), ",")

	scoreP1 := 0
	for _, step := range steps {
		scoreP1 += hash(step)
	}

	// Part 1
	fmt.Println("Sol 1:", scoreP1)
}
