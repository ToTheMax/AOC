package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("in.txt")
	pairs := strings.Split(string(input), "\n")
	score1 := 0
	score2 := 0
	for _, pair := range pairs {
		elves := strings.Split(string(pair), ",")
		elve1 := strings.Split(string(elves[0]), "-")
		elve2 := strings.Split(string(elves[1]), "-")

		x1, _ := strconv.Atoi(elve1[0])
		y1, _ := strconv.Atoi(elve1[1])

		x2, _ := strconv.Atoi(elve2[0])
		y2, _ := strconv.Atoi(elve2[1])

		// Part 1
		if x1 >= x2 && y1 <= y2 {
			score1 += 1
		} else if x2 >= x1 && y2 <= y1 {
			score1 += 1
		}

		// Part 2
		if x1 <= x2 && y1 >= x2 {
			score2 += 1
		} else if x2 <= x1 && y2 >= x1 {
			score2 += 1

		}
	}
	fmt.Println("Sol1:", score1)
	fmt.Println("Sol2:", score2)
}
