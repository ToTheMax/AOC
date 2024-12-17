package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func get_region(lines []string, visited map[Pos]bool, pos Pos) []Pos {
	letter := lines[pos.y][pos.x]
	visited[pos] = true

	region := []Pos{pos}
	for _, direction := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		nx, ny := pos.x+direction[0], pos.y+direction[1]

		if nx < 0 || nx >= len(lines[0]) || ny < 0 || ny >= len(lines) || // Bounds
			visited[Pos{nx, ny}] || lines[ny][nx] != letter {
			continue
		}
		region = append(region, get_region(lines, visited, Pos{nx, ny})...)
	}
	return region
}

func get_perimeter(lines []string, region []Pos) int {
	region_map := make(map[Pos]bool)
	for _, pos := range region {
		region_map[pos] = true
	}

	perimeter := 0
	for _, pos := range region {
		for _, direction := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nx, ny := pos.x+direction[0], pos.y+direction[1]
			// Ff not in region_map, it is a perimeter
			if !region_map[Pos{nx, ny}] {
				perimeter++
			}
		}
	}
	return perimeter
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	visited := make(map[Pos]bool)

	solP1 := 0
	for y, line := range lines {
		for x := range line {
			if !visited[Pos{x, y}] {
				region := get_region(lines, visited, Pos{x, y})
				solP1 += len(region) * get_perimeter(lines, region)
			}
		}
	}

	fmt.Println("Sol 1:", solP1)
}
