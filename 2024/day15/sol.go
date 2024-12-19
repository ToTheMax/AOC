package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func tryMove(movement rune, grid []string, robotPos *Pos, boxes map[Pos]bool) bool {
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

	newPos := Pos{robotPos.x + dx, robotPos.y + dy}
	// If the new position is a wall skip movement
	if grid[newPos.y][newPos.x] == '#' {
		return false
	}

	// If new position is a box, check if it/they can be moved
	if boxes[newPos] {
		stackedBoxes := []Pos{newPos}
		currentPos := newPos
		for {
			nextPos := Pos{currentPos.x + dx, currentPos.y + dy}
			if boxes[nextPos] {
				stackedBoxes = append(stackedBoxes, nextPos)
				currentPos = nextPos
			} else if grid[nextPos.y][nextPos.x] == '#' {
				return false
			} else {
				break
			}
		}
		// Move boxes
		delete(boxes, newPos)
		for _, pos := range stackedBoxes {
			newBoxPos := Pos{pos.x + dx, pos.y + dy}
			boxes[newBoxPos] = true
		}
	}

	*robotPos = newPos
	return true
}

func main() {
	input, _ := os.ReadFile("in.txt")
	inputs := strings.Split(string(input), "\n\n")

	grid := strings.Split(inputs[0], "\n")
	movements := strings.ReplaceAll(inputs[1], "\n", "")

	var robotPos Pos
	boxes := make(map[Pos]bool)

	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				robotPos = Pos{x, y}
				grid[y] = grid[y][:x] + "." + grid[y][x+1:]
			} else if cell == 'O' {
				boxes[Pos{x, y}] = true
				grid[y] = grid[y][:x] + "." + grid[y][x+1:]
			}
		}
	}

	for _, movement := range movements {
		tryMove(movement, grid, &robotPos, boxes)
		// print map
		// fmt.Println("Move:", string(movement))
		// for y, row := range grid {
		// 	for x, cell := range row {
		// 		if boxes[Pos{x, y}] {
		// 			fmt.Print("O")
		// 		} else if robotPos == (Pos{x, y}) {
		// 			fmt.Print("@")
		// 		} else {
		// 			fmt.Print(string(cell))
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()
	}

	solP1 := 0
	for pos := range boxes {
		solP1 += pos.x + pos.y*100
	}
	fmt.Println("Sol 1:", solP1)
}
