package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func getAccessibleRolls(grid map[Pos]rune, adjacentOffsets []Pos) []Pos {
	var accessibleRolls []Pos
	for pos, ch := range grid {
		if ch != '@' {
			continue
		}
		adjacentCount := 0
		for _, offset := range adjacentOffsets {
			adj := Pos{pos.x + offset.x, pos.y + offset.y}
			if grid[adj] == '@' {
				adjacentCount++
			}
		}
		if adjacentCount < 4 {
			accessibleRolls = append(accessibleRolls, pos)
		}
	}
	return accessibleRolls
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	grid := make(map[Pos]rune)

	for y, line := range lines {
		for x, ch := range line {
			grid[Pos{x, y}] = ch
		}
	}

	adjacentOffsets := []Pos{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	removableRolls := getAccessibleRolls(grid, adjacentOffsets)
	fmt.Println("Sol 1:", len(removableRolls))

	score2 := 0
	for {
		if len(removableRolls) == 0 {
			break
		}
		for _, pos := range removableRolls {
			score2 += 1
			grid[pos] = '.'
		}
		removableRolls = getAccessibleRolls(grid, adjacentOffsets)
	}

	fmt.Println("Sol 2:", score2)
}
