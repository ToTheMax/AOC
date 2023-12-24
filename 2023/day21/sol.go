package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

type State struct {
	pos Pos
	steps int
}


type Grid struct {
	g [][]string
	h int
	w int
}


func takeStep(grid Grid, prevVisited map[Pos]bool) map[Pos]bool {
	visited := map[Pos]bool{}

	for pos := range prevVisited {
		for _, dir := range []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			newPos := Pos{pos.x + dir.x,pos.y + dir.y}
			if newPos.x < 0 || newPos.x >= grid.w || newPos.y < 0 || newPos.y >= grid.h {
				continue
			}
			if grid.g[newPos.y][newPos.x] == "#" {
				continue
			}
			visited[newPos] = true
			}
		}
	return visited
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	grid := Grid{make([][]string, len(lines)), len(lines), len(lines[0])}
	var startPos Pos
	for y, line := range lines {
		grid.g[y] = make([]string, len(line))
		for x, char := range line {
			grid.g[y][x] = string(char)
			if char == 'S'{
				startPos = Pos{x, y}
			}
		}
	}

	visited := map[Pos]bool{startPos: true}
	for i:=0; i<64; i++ {
		visited = takeStep(grid, visited)
	}
	fmt.Println("Sol 1:", len(visited))
}
