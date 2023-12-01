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
				vents[x1][i]++
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
				vents[i][y1]++
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
