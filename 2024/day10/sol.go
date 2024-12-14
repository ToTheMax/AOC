package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func total_hiking_trails(lines []string, pos Pos, visited map[Pos]bool, part int) int {
	if lines[pos.y][pos.x] == '9' {
		if part == 2 || !visited[pos] {
			visited[pos] = true
			return 1
		}
	}

	// Check all 4 directions, only go if next step is 1 higher
	total := 0
	directions := []Pos{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, direction := range directions {
		newPos := Pos{pos.x + direction.x, pos.y + direction.y}
		if newPos.x < 0 || newPos.x >= len(lines[0]) || newPos.y < 0 || newPos.y >= len(lines) {
			continue
		} else if lines[newPos.y][newPos.x] == lines[pos.y][pos.x]+1 {
			total += total_hiking_trails(lines, newPos, visited, part)
		}
	}
	return total
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	// Find all trailheads
	trailheads := []Pos{}
	for y, line := range lines {
		for x, char := range line {
			if char == '0' {
				trailheads = append(trailheads, Pos{x, y})
			}
		}
	}
	// Find all hiking trails
	sumP1 := 0
	sumP2 := 0
	for _, trailhead := range trailheads {
		sumP1 += total_hiking_trails(lines, trailhead, map[Pos]bool{}, 1)
		sumP2 += total_hiking_trails(lines, trailhead, map[Pos]bool{}, 2)
	}
	fmt.Println("Sol 1:", sumP1)
	fmt.Println("Sol 2:", sumP2)
}
