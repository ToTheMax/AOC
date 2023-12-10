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

func findMove(grid Grid, prevMove Move, curPos Pos) Move {
	x := curPos.x
	y := curPos.y

	if grid.g[y][x] == "F" {
		if prevMove.dy == -1 {
			return Move{1, 0}
		} else {
			return Move{0, 1}
		}
	} else if grid.g[y][x] == "7" {
		if prevMove.dy == -1 {
			return Move{-1, 0}
		} else {
			return Move{0, 1}
		}
	} else if grid.g[y][x] == "L" {
		if prevMove.dy == 1 {
			return Move{1, 0}
		} else {
			return Move{0, -1}
		}
	} else if grid.g[y][x] == "J" {
		if prevMove.dy == 1 {
			return Move{-1, 0}
		} else {
			return Move{0, -1}
		}
	} else if grid.g[y][x] == "|" {
		if prevMove.dy == 1 {
			return Move{0, 1}
		} else {
			return Move{0, -1}
		}
	} else if grid.g[y][x] == "-" {
		if prevMove.dx == 1 {
			return Move{1, 0}
		} else {
			return Move{-1, 0}
		}
	}
	return Move{0, 0}
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")
	n := len(lines)
	m := len(lines[0])

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

	// First Move
	curPos := startPos
	var startMove Move
	fmt.Println(grid)
	for _, move := range []Move{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		checkPos := Pos{curPos.x + move.dx, curPos.y + move.dy}
		x := findMove(grid, move, checkPos)
		if x.dx != 0 || x.dy != 0 {
			startMove = move
			curPos = checkPos
			break
		}
	}
	fmt.Println("Starting", curPos, startMove)

	// Follow path
	prevMove := startMove
	steps := 1
	for curPos != startPos {
		// fmt.Println("Pos:", curPos, "Move:", prevMove)
		prevMove = findMove(grid, prevMove, curPos)
		curPos.x += prevMove.dx
		curPos.y += prevMove.dy
		steps++
	}

	fmt.Println("Sol 1:", steps/2)
	fmt.Println("Sol 2:", 0)
}
