package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const WIDTH = 40

func drawPixel(X int, cycle *int) {
	if *cycle%WIDTH >= X && *cycle%WIDTH <= X+2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if *cycle%WIDTH == 0 {
		fmt.Print("\n")
	}
}

func incCycle(cycle *int, X int) int {
	*cycle += 1
	if (*cycle-WIDTH/2)%WIDTH == 0 {
		return X * *cycle
	}
	return 0
}

func tick(cycle *int, X *int, increment int) int {
	drawPixel(*X, cycle)
	*X += increment
	return incCycle(cycle, *X)
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	cycle := 1
	sum := 0
	X := 1

	fmt.Println("Sol2:")
	for _, line := range lines {
		if line == "noop" {
			sum += tick(&cycle, &X, 0)
		} else {
			sum += tick(&cycle, &X, 0)
			splitted := strings.Split(line, " ")
			value, _ := strconv.Atoi(splitted[1])
			sum += tick(&cycle, &X, value)
		}
	}
	fmt.Println("Sol1:", sum)
}
