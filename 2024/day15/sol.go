package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func isMovableBox(dx, dy int, grid []string, boxPos Pos, boxes map[Pos]rune, movableBoxes *[]Pos, seen map[Pos]bool) bool {

	// If it is a wall, not movable
	if grid[boxPos.y][boxPos.x] == '#' {
		return false
	}

	// If it is already seen, don't check it again
	if seen[boxPos] {
		return true
	}

	// If it is a box, see what is behind it
	if _, exists := boxes[boxPos]; exists {

		// Add the box itself
		if !seen[boxPos] {
			*movableBoxes = append(*movableBoxes, boxPos)
			seen[boxPos] = true
		}

		// O box
		if boxes[boxPos] == 'O' {
			return isMovableBox(dx, dy, grid, Pos{boxPos.x + dx, boxPos.y + dy}, boxes, movableBoxes, seen)
		}

		// [] box
		multiplier := 1
		if boxes[boxPos] == ']' {
			multiplier = -1
		}

		boxPos2 := Pos{boxPos.x + multiplier, boxPos.y}
		if !seen[boxPos2] {
			*movableBoxes = append(*movableBoxes, boxPos2)
			seen[boxPos2] = true
		}

		// Horizontal movement
		if dx != 0 {
			return isMovableBox(dx, dy, grid, Pos{boxPos.x + 2*multiplier, boxPos.y}, boxes, movableBoxes, seen)
		}
		// Vertical movement
		if dy != 0 {
			return isMovableBox(dx, dy, grid, Pos{boxPos.x, boxPos.y + dy}, boxes, movableBoxes, seen) &&
				isMovableBox(dx, dy, grid, Pos{boxPos.x + multiplier, boxPos.y + dy}, boxes, movableBoxes, seen)
		}
	}
	return true
}

func tryMove(movement rune, grid []string, currRobotPos Pos, boxes map[Pos]rune) Pos {
	var dx, dy int
	switch movement {
	case '^':
		dx, dy = 0, -1
	case '>':
		dx, dy = 1, 0
	case 'v':
		dx, dy = 0, 1
	case '<':
		dx, dy = -1, 0
	}

	newRobotPos := Pos{currRobotPos.x + dx, currRobotPos.y + dy}

	// If the new position is a wall, don't move
	if grid[newRobotPos.y][newRobotPos.x] == '#' {
		return currRobotPos
	}

	// If new position is a box, check whether it/they can be moved
	if _, exists := boxes[newRobotPos]; exists {

		movableBoxes := []Pos{}
		visited := make(map[Pos]bool)

		if !isMovableBox(dx, dy, grid, newRobotPos, boxes, &movableBoxes, visited) {
			return currRobotPos
		}

		// Move boxes
		boxesValues := make(map[Pos]rune)
		for _, boxPos := range movableBoxes {
			boxesValues[boxPos] = boxes[boxPos]
		}
		for _, boxPos := range movableBoxes {
			delete(boxes, boxPos)
		}
		for _, boxPos := range movableBoxes {
			boxes[Pos{boxPos.x + dx, boxPos.y + dy}] = boxesValues[boxPos]
		}

	}
	return newRobotPos
}

func solve(inputs []string, part int) int {

	grid := strings.Split(inputs[0], "\n")
	movements := strings.ReplaceAll(inputs[1], "\n", "")

	// For part 2, make map twice as wide
	if part == 2 {
		grid_input := inputs[0]
		grid_input = strings.ReplaceAll(grid_input, "#", "##")
		grid_input = strings.ReplaceAll(grid_input, "O", "[]")
		grid_input = strings.ReplaceAll(grid_input, ".", "..")
		grid_input = strings.ReplaceAll(grid_input, "@", "@.")
		grid = strings.Split(grid_input, "\n")
	}

	var robotPos Pos
	boxes := make(map[Pos]rune)

	// Make grid
	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				robotPos = Pos{x, y}
				grid[y] = grid[y][:x] + "." + grid[y][x+1:]
			} else if cell == 'O' || cell == '[' || cell == ']' {
				boxes[Pos{x, y}] = cell
				grid[y] = grid[y][:x] + "." + grid[y][x+1:]
			}
		}
	}

	// Do moves
	for _, movement := range movements {
		robotPos = tryMove(movement, grid, robotPos, boxes)
	}

	// Calculate score
	score := 0
	for pos, char := range boxes {
		if char == '[' || char == 'O' {
			score += pos.x + pos.y*100
		}
	}
	return score
}

func main() {
	input, _ := os.ReadFile("in.txt")
	inputs := strings.Split(string(input), "\n\n")

	fmt.Println("Sol 1:", solve(inputs, 1))
	fmt.Println("Sol 2:", solve(inputs, 2))
}
