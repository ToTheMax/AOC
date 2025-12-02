package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Rotation struct {
	Number    int
	Direction Direction
}

func main() {
	input, _ := os.ReadFile("in.txt")
	single_line := strings.ReplaceAll(string(input), "\n", "")

	// Make the ranges
	var ranges [][2]int
	ranges_str := strings.Split(single_line, ",")
	for _, range_str := range ranges_str {
		bounds := strings.Split(range_str, "-")
		from, _ := strconv.Atoi(bounds[0])
		to, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, [2]int{from, to})
	}

	// Loop over the ranges
	score1, score2 := 0, 0
	for _, r := range ranges {
		from := r[0]
		to := r[1]
		for i := from; i <= to+1; i++ {
			// For part 1, check if the number has two equal halves
			str := strconv.Itoa(i)
			half := len(str) / 2
			if str[:half] == str[half:] {
				invalid_id, _ := strconv.Atoi(str)
				score1 += invalid_id
			}

			// For part 2, check for any recurring pattern
			n := len(str)
			for l := 1; l <= n/2; l++ {
				pattern := str[:l]
				repeats := n / l
				built := strings.Repeat(pattern, repeats)
				if built == str {
					invalid_id, _ := strconv.Atoi(str)
					score2 += invalid_id
					break
				}
			}
		}
	}

	fmt.Println("Sol 1:", score1)
	fmt.Println("Sol 2:", score2)
}
