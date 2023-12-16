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
	energized map[int]bool
	seen map[State]bool
}

type State struct {
	x int
	y int
	direction rune
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func nextState(state State) State {
	if state.direction=='N' {
		return State{state.x, state.y-1, state.direction}
	} else if state.direction=='E' {
		return State{state.x+1, state.y, state.direction}
	} else if state.direction=='S' {
		return State{state.x, state.y+1, state.direction}
	} else{
		return State{state.x-1, state.y, state.direction}
	}
}

func move(grid Grid, state State) {
	if state.x < 0 || state.x >= grid.m || state.y < 0 || state.y >= grid.n {
		return
	}
	if _, ok := grid.seen[state]; ok {
		return
	}
	grid.seen[state] = true	
	grid.energized[state.x + state.y*grid.m] = true
	if grid.g[state.y][state.x] == "." {
		move(grid, nextState(state))
	} else if grid.g[state.y][state.x] == "|" {
		if state.direction=='N' || state.direction=='S' {
			move(grid, nextState(state))
		} else {
			move(grid, State{state.x, state.y-1, 'N'})
			move(grid, State{state.x, state.y+1, 'S'})
		}
	} else if grid.g[state.y][state.x] == "-" {
		if state.direction=='E' || state.direction=='W' {
			move(grid, nextState(state))
		} else {
			move(grid, State{state.x-1, state.y, 'W'})
			move(grid, State{state.x+1, state.y, 'E'})
		}
	} else if grid.g[state.y][state.x] == "/" {
		if state.direction=='N' {
			move(grid, State{state.x+1, state.y, 'E'})
		} else if state.direction=='E' {
			move(grid, State{state.x, state.y-1, 'N'})
		} else if state.direction=='S' {
			move(grid, State{state.x-1, state.y, 'W'})
		} else {
			move(grid, State{state.x, state.y+1, 'S'})
		}
	} else if grid.g[state.y][state.x] == "\\" {
		if state.direction=='N' {
			move(grid, State{state.x-1, state.y, 'W'})
		} else if state.direction=='E' {
			move(grid, State{state.x, state.y+1, 'S'})
		} else if state.direction=='S' {
			move(grid, State{state.x+1, state.y, 'E'})
		} else {
			move(grid, State{state.x, state.y-1, 'N'})
		}
	}
}

func getScore(grid Grid, startState State) int {
	grid = Grid{grid.g, grid.n, grid.m, make(map[int]bool), make(map[State]bool)}
	move(grid, startState)
	return len(grid.energized)

}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")
	n := len(lines)
	m := len(lines[0])

	// Read grid
	grid := Grid{make([][]string, n), n, m,  make(map[int]bool), make(map[State]bool)}
	for i := range grid.g {
		grid.g[i] = make([]string, m)
		for j, char := range lines[i] {
			grid.g[i][j] = string(char)
		}
	}

	// Part 1
	move(grid, State{0, 0, 'E'})
	solP1 := len(grid.energized)
	fmt.Println("Sol 1:", getScore(grid, State{0, 0, 'E'}))
	
	
	// Part 2
	solP2 := solP1
	for i := 0; i < grid.m; i++ {
		solP2 = max(solP2, getScore(grid, State{i, 0, 'S'}))
		solP2 = max(solP2, getScore(grid, State{i, n-1, 'N'}))
	}
	for j := 0; j < grid.n; j++ {
		solP2 = max(solP2, getScore(grid, State{0, j, 'E'}))
		solP2 = max(solP2, getScore(grid, State{m-1, j, 'W'}))
	}
	fmt.Println("Sol 2:", solP2)
}
