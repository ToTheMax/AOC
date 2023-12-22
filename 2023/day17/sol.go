package main

import (
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Pos struct {
	x int
	y int
}

type State struct {
	pos Pos
	dir Pos
	streak     int
}

type Grid struct {
	vals [][]int
	width int
	height int
}


func calcHeatLoss(grid Grid, start, end Pos, minConseq, maxConseq int) int {
	queue := []State{{start, Pos{1, 0}, 0}, {start, Pos{0, 1}, 0}}
	visited := map[State]int{{start, Pos{0, 0}, 0}: 0}
	minHeatLoss := 1000000000

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == end && current.streak >= minConseq {
			minHeatLoss = min(minHeatLoss, visited[current])
		}

		newDirs := []Pos{current.dir, {current.dir.y, -current.dir.x}, {-current.dir.y, current.dir.x}}
		for _, dir := range newDirs {
			nextPos := Pos{current.pos.x + dir.x, current.pos.y + dir.y}

			// Check if in grid
			if nextPos.x < 0 || nextPos.x >= grid.width  {
				continue
			} else if nextPos.y < 0 || nextPos.y >= grid.height {
				continue
			}

			totalHeatLoss := visited[current] + grid.vals[nextPos.y][nextPos.x]
			nextStreak := 1
			if dir == current.dir {
				nextStreak = current.streak + 1
			}
			if (dir == current.dir && current.streak < maxConseq) ||
				(dir != current.dir && current.streak >= minConseq) {
				nextState := State{nextPos, dir, nextStreak}
				if val, found := visited[nextState]; !found || val > totalHeatLoss {
					visited[nextState] = totalHeatLoss
					queue = append(queue, nextState)
				}
			}
		}
	}
	return minHeatLoss
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	grid := Grid{make([][]int, len(lines)), len(lines[0]), len(lines)}
	for i, line := range lines {
		grid.vals[i] = make([]int, grid.width)
		for j := 0; j < len(line); j++ {
			grid.vals[i][j] = int(line[j] - '0')
		}
	}

	start := Pos{0, 0}
	end := Pos{grid.width- 1, grid.height - 1}

	fmt.Println("Sol P2", calcHeatLoss(grid, start, end, 1, 3))
	fmt.Println("Sol P2", calcHeatLoss(grid, start, end, 4, 10))
}