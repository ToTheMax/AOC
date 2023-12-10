package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	g [][]string
	n int
	m int
}

type Move struct {
	dx int
	dy int
}

type Pos struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findMove(prevMove Move, tile string) Move {
	// Mapping of prevMove to nextMove
	mapping := map[string][]Move{
		"F": {{1, 0}, {0, 1}},
		"7": {{-1, 0}, {0, 1}},
		"L": {{1, 0}, {0, -1}},
		"J": {{-1, 0}, {0, -1}},
	}
	if _, ok := mapping[tile]; ok {
		return mapping[tile][abs(prevMove.dx)]
	} else if strings.Index("|-", tile) != -1{
		return Move{prevMove.dx, prevMove.dy}
	} else{
		return Move{0, 0}
	}
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")
	n := len(lines)
	m := len(lines[0])

	// Read grid
	grid := Grid{make([][]string, n), n, m}
	startPos := Pos{0, 0}
	for i := range grid.g {
		grid.g[i] = make([]string, m)
		for j, char := range lines[i] {
			grid.g[i][j] = string(char)
			if char == 'S' {
				startPos = Pos{j, i}
			}
		}
	}

	// Starting Move
	curPos := startPos
	var startMove Move
	for _, move := range []Move{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		checkPos := Pos{curPos.x + move.dx, curPos.y + move.dy}
		checkMove := findMove(move, grid.g[checkPos.y][checkPos.x])
		if checkMove.dx != 0 || checkMove.dy != 0 {
			startMove = move
			curPos = checkPos
			break
		}
	}

	// Follow path
	prevMove := startMove
	steps := 1
	visited := map[Pos]bool{startPos:true, curPos:true}
	for curPos != startPos {
		prevMove = findMove(prevMove, grid.g[curPos.y][curPos.x])
		curPos = Pos{curPos.x + prevMove.dx, curPos.y + prevMove.dy}
		visited[curPos] = true
		steps++
	}

	fmt.Println("Sol 1:", steps/2)

	// Find enclosed tiles
	totalCount := 0
	for y := range grid.g {
		count := 0
		for x := range grid.g[y] {
			t := grid.g[y][x]
			if strings.Index("|JL", t) != -1 && visited[Pos{x, y}] {
				count++
			}
			if count%2 == 1 && !visited[Pos{x, y}] {
				totalCount++
			}
		
		}
	}
	fmt.Println("Sol 2:", totalCount)
}
