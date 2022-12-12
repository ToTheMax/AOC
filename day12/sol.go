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

func BFS(grid Grid, start int, goal int) int {

	queue := []int{start}
	visited := make(map[int]int)

	curPos := 0
	for len(queue) > 0 {
		// Check if goal is reached
		curPos = queue[0]
		if curPos == goal {
			// Traverse path length
			pathLength := 0
			for curPos != start {
				curPos = visited[curPos]
				pathLength++
			}
			return pathLength
		}
		// Add neighbors to queue
		for _, newPos := range getNeighborPositions(grid, curPos) {
			if _, exists := visited[newPos]; !exists {
				if grid.heights[newPos]-1 <= grid.heights[curPos] {
					visited[newPos] = curPos
					queue = append(queue, newPos)
				}
			}
		}
		queue = queue[1:]
	}
	return -1
}

func main() {
	input, _ := os.ReadFile("in.txt")

	lines := strings.Split(string(input), "\n")
	heights := make([]int, len(lines)*len(lines[0]))

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
	grid := Grid{heights, len(lines), len(lines[0])}

	distance := BFS(grid, startPos, goalPos)
	fmt.Println("Sol1:", distance)

	shortest := distance
	for pos, height := range grid.heights {
		if height == 0 {
			distance = BFS(grid, pos, goalPos)
			if distance > 0 && distance < shortest {
				shortest = distance
			}
		}
	}
	fmt.Println("Sol2:", shortest)
}
