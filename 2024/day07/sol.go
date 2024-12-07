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

func solveEquation(equation []int, operators []Operator, value int, target int) int {
	solves := 0
	if len(equation) == 0 {
		if value == target {
			return 1
		}
		return 0
	}

	for _, operator := range operators {
		switch operator {
		case Add:
			solves += solveEquation(equation[1:], operators, value+equation[0], target)
		case Multiply:
			solves += solveEquation(equation[1:], operators, value*equation[0], target)
		case Concat:
			value, _ = strconv.Atoi(fmt.Sprintf("%d%d", value, equation[0]))
			solves += solveEquation(equation[1:], operators, value, target)
		}
	}
	return solves
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
		if solveEquation(ints[1:], []Operator{Add, Multiply}, ints[0], target) > 0 {
			solP1 += target
		}
		if solveEquation(ints[1:], []Operator{Add, Multiply, Concat}, ints[0], target) > 0 {
			solP2 += target
		}
	}
	fmt.Println("Sol1", solP1)
	fmt.Println("Sol2", solP2)
}
