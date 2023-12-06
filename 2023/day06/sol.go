package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringToInts(s string) []int {
	splitted := strings.Split(s, " ")
	ints := make([]int, len(splitted))
	for i, s := range splitted {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func getScore(time int, record int) int {

	a := -1.0
	b := float64(time)
	c := -float64(record)

	det := b*b - 4*a*c

	if det > 0 {
		root1 := int(math.Ceil((-b+math.Sqrt(det))/(2*a) + 0.01))
		root2 := int(math.Floor((-b-math.Sqrt(det))/(2*a) - 0.01))
		return root2 - root1 + 1
	}
	return 0
}

func main() {

	input, _ := os.ReadFile("in.txt")
	input_str := strings.ReplaceAll(string(input), "  ", " ")
	input_str = strings.ReplaceAll(input_str, "  ", " ")
	input_str = strings.ReplaceAll(input_str, "  ", " ")
	lines := strings.Split(input_str, "\n")
	times := stringToInts(lines[0][6:])
	distances := stringToInts(lines[1][10:])

	totalWays := 1

	for i := 0; i < len(times); i++ {
		totalWays *= getScore(times[i], distances[i])
	}

	fmt.Println("Sol 1:", totalWays)

	input_str = strings.ReplaceAll(input_str, " ", "")
	lines = strings.Split(input_str, "\n")
	times = stringToInts(lines[0][5:])
	distances = stringToInts(lines[1][9:])

	fmt.Println("Sol 2:", getScore(times[0], distances[0]))
}
