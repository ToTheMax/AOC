package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	heights []int
	height  int
	width   int
}

func getNeighborPositions(grid Grid, curPos int) []int {
	var neighborPositions []int

	curX := curPos % grid.width
	curY := (curPos - curX) / grid.width

	if curX < grid.width-1 { // Check right
		neighborPositions = append(neighborPositions, (curX+1)+curY*grid.width)
	}
	if curX > 0 { // Check left
		neighborPositions = append(neighborPositions, (curX-1)+curY*grid.width)
	}
	if curY < grid.height-1 { // Check up
		neighborPositions = append(neighborPositions, curX+(curY+1)*grid.width)
	}
	if curY > 0 { // Check down
		neighborPositions = append(neighborPositions, curX+(curY-1)*grid.width)
	}
	return neighborPositions
}

func BFS(grid Grid, source int, sink int) map[int]int {
	distances := make(map[int]int)
	queue := []int{source}
	curPos := 0
	for len(queue) > 0 {
		// Check if goal is reached
		curPos = queue[0]
		if curPos == sink {
			break
		}
		// Add neighbors to queue
		for _, newPos := range getNeighborPositions(grid, curPos) {
			if _, exists := distances[newPos]; !exists {
				if grid.heights[newPos]+1 >= grid.heights[curPos] {
					distances[newPos] = distances[curPos] + 1
					queue = append(queue, newPos)
				}
			}
		}
		queue = queue[1:]
	}
	return distances
}

func main() {
	input, _ := os.ReadFile("in.txt")

	lines := strings.Split(string(input), "\n")
	heights := make([]int, len(lines)*len(lines[0]))
	grid := Grid{heights, len(lines), len(lines[0])}

	// Fill grid
	startPos := 0
	goalPos := 0
	for y, line := range lines {
		for x, char := range line {
			pos := x + len(lines[0])*y
			switch char {
			case 'S':
				startPos = pos
				char = 'a'
			case 'E':
				goalPos = pos
				char = 'z'
			}
			heights[x+len(lines[0])*y] = int(char) - int('a')
		}
	}

	// Calculate shortest paths
	distances := BFS(grid, goalPos, startPos)

	fmt.Println("Sol1:", distances[startPos])

	shortest := distances[startPos]
	for pos, height := range grid.heights {
		if height == 0 {
			if distances[pos] > 0 && distances[pos] < shortest {
				shortest = distances[pos]
			}
		}
	}
	fmt.Println("Sol2:", shortest)
}
