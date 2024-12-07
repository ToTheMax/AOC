package main

import (
	"fmt"
	"os"
	"strings"
)

type Guard struct {
	pos Pos
	dx, dy int
}

type Pos struct {
	x, y int
}


func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	// Find guard
	var guard Guard
	for y := range lines {
		for x, c := range lines[y] {
			if c == '^'{
				guard = Guard{Pos{x, y}, 0, -1}
			}
		}
	}

	// Do steps
	seen := make(map[Pos]bool)
	seen[guard.pos] = true
	for true  {
		newPos := Pos{guard.pos.x + guard.dx, guard.pos.y + guard.dy}
		if newPos.x < 0 || newPos.y < 0 || newPos.x >= len(lines[0]) || newPos.y >= len(lines){
			break
		} 
		if lines[newPos.y][newPos.x] == '#'{
			guard.dx, guard.dy = -guard.dy, guard.dx
		} else{
			guard.pos = newPos
			seen[guard.pos] = true
		}
	}		 
	fmt.Println("Sol1", len(seen))
}
