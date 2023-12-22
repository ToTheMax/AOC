package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	lava   []int
	height int
	width  int
}

type Pos struct {
	x int
	y int
}

func fill(grid map[Pos]string, pos Pos) {
	grid[pos] = ""
	surrounedPos := []Pos{
		{pos.x, pos.y + 1},
		{pos.x, pos.y - 1},
		{pos.x + 1, pos.y},
		{pos.x - 1, pos.y},
	}
	for _, newPos := range surrounedPos {
		if _, ok := grid[newPos]; !ok {
			fill(grid, newPos)
		}
	}
}

func main() {
	input, _ := os.ReadFile("in.txt")

	// Part 1
	lines := strings.Split(string(input), "\n")
	grid := map[Pos]string{}
	curPos := Pos{0, 0}
	for _, line := range lines {
		split := strings.Split(line, " ")
		dir := split[0]
		amount, _ := strconv.Atoi(split[1])
		color := split[2]
		for i := 0; i < amount; i++ {
			switch dir {
			case "D":
				curPos.y++
			case "U":
				curPos.y--
			case "R":
				curPos.x++
			case "L":
				curPos.x--
			}
			grid[curPos] = color
		}
	}

	for pos := range grid {
		fmt.Println(pos)
	}
	fill(grid, Pos{1, 1})
	fmt.Println("Sol 1", len(grid))

	// // Part 2
	// fmt.Println("Sol 2:", 0)
}
