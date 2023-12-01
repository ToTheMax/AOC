package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	file, _ := os.Open(path)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func absDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {

	lines := readLines("input.txt")
	mapSize := 1000

	vents := make([][]int, mapSize)
	for i, _ := range vents {
		vents[i] = make([]int, mapSize)
	}

	for _, line := range lines {

		// Parse line
		splitted := strings.Split(line, " -> ")
		c1 := strings.Split(splitted[0], ",")
		c2 := strings.Split(splitted[1], ",")
		x1, _ := strconv.Atoi(c1[0])
		y1, _ := strconv.Atoi(c1[1])
		x2, _ := strconv.Atoi(c2[0])
		y2, _ := strconv.Atoi(c2[1])

		// Vertical lines
		if x1 == x2 {
			from := y1
			to := y2
			if y1 > y2 {
				from = y2
				to = y1
			}
			for i := from; i <= to; i++ {
				vents[i][x1]++
			}
		}

		// Horizontal lines
		if y1 == y2 {
			from := x1
			to := x2
			if x1 > x2 {
				from = x2
				to = x1
			}
			for i := from; i <= to; i++ {
				vents[y1][i]++
			}
		}

		// Verical lines
		if absDiffInt(x1, x2) == absDiffInt(y1, y2) {
			xdiff := 1
			if x2 < x1 {
				xdiff = -1
			}
			ydiff := 1
			if y2 < y1 {
				ydiff = -1
			}
			for i := 0; i <= absDiffInt(y1, y2); i++ {
				vents[y1+i*ydiff][x1+i*xdiff]++
			}
		}
	}

	dangerousAreas := 0
	for _, row := range vents {
		for _, val := range row {
			if val >= 2 {
				dangerousAreas++
			}
		}
	}
	fmt.Println(dangerousAreas)
}
