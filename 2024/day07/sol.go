package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	Add Operator = iota
	Multiply
	Concat
)

func solveEquation(equation []int, operators []Operator, value int, target int) bool {
	if len(equation) == 0 {
		return value == target
	}

	solved := false
	for _, operator := range operators {
		switch operator {
		case Add:
			solved = solveEquation(equation[1:], operators, value+equation[0], target)
		case Multiply:
			solved = solved || solveEquation(equation[1:], operators, value*equation[0], target)
		case Concat:
			value, _ = strconv.Atoi(fmt.Sprintf("%d%d", value, equation[0]))
			solved = solved || solveEquation(equation[1:], operators, value, target)
		}
	}
	return solved
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	solP1 := 0
	solP2 := 0
	for _, line := range lines {
		splitted := strings.Split(line, ": ")
		target, _ := strconv.Atoi(splitted[0])
		splitted = strings.Split(splitted[1], " ")

		ints := make([]int, len(splitted))
		for i, s := range splitted {
			ints[i], _ = strconv.Atoi(s)
		}
		if solveEquation(ints[1:], []Operator{Add, Multiply}, ints[0], target) {
			solP1 += target
		}
		if solveEquation(ints[1:], []Operator{Add, Multiply, Concat}, ints[0], target) {
			solP2 += target
		}
	}
	fmt.Println("Sol1", solP1)
	fmt.Println("Sol2", solP2)
}
