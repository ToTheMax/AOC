package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

func stringToInts(s string) []int {
	splitted := strings.Split(s, " ")
	ints := make([]int, len(splitted))
	for i, s := range splitted {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func getDiff(values []int, level int) int {
	i := (2 - level) * (len(values) - 2)
	m := (3 - level*2)

	// Create difference array
	diffs := make([]int, len(values)-1)
	allZero := true
	for i := 0; i < len(values)-1; i++ {
		diffs[i] = values[i+1] - values[i]
		if diffs[i] != 0 {
			allZero = false
		}
	}
	// If not all zero, recurse
	if allZero {
		return diffs[i]
	} else {
		return diffs[i] + m*getDiff(diffs, level)
	}
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	sumPredictions := 0
	sumPrevPredictions := 0
	for _, line := range lines {
		values := stringToInts(line)
		sumPredictions += values[len(values)-1] + getDiff(values, 1)
		sumPrevPredictions += values[0] - getDiff(values, 2)
	}

	fmt.Println("Sol 1:", sumPredictions)
	fmt.Println("Sol 2:", sumPrevPredictions)
}
