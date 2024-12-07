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

	
	// Find guard and make obstructions
	var obstructions [][]string
	var guardPos Pos
	for y := range lines {
		for x, c := range lines[y] {
			if c == '^'{
				guardPos = Pos {x, y}
			}
			if c == '.'{
				obstruction := append([]string(nil), lines...)
				obstruction[y] = obstruction[y][:x] + "#" + obstruction[y][x+1:] 
				obstructions = append(obstructions, obstruction)
			}
		}
	}

	// Count out of the map steps
	seen := make(map[Pos]bool)
	guard := Guard{guardPos, 0, -1}
	seen[guard.pos] = true
	for true  {
		newPos := Pos{guard.pos.x + guard.dx, guard.pos.y + guard.dy}
		if newPos.x < 0 || newPos.y < 0 || newPos.x >= len(lines[0]) || newPos.y >= len(lines){
			break
		} 
		if lines[newPos.y][newPos.x] == '#'{
			guard.dx, guard.dy = -guard.dy, guard.dx
		} else {
			guard.pos = newPos
			seen[guard.pos] = true
		}
	}		 
	fmt.Println("Sol1", len(seen))

	// Count obstructions with a loop
	count := 0
	fmt.Println("Trying obstructions :", len(obstructions))
	for _, obstruction := range obstructions {
		guard := Guard{guardPos, 0, -1}
		seen := make(map[Guard]bool)
		seen[guard] = true
		for true  {
			newPos := Pos{guard.pos.x + guard.dx, guard.pos.y + guard.dy}
			if newPos.x < 0 || newPos.y < 0 || newPos.x >= len(lines[0]) || newPos.y >= len(lines){
				break
			} 
			if obstruction[newPos.y][newPos.x] == '#'{
				guard.dx, guard.dy = -guard.dy, guard.dx
			} else {
				guard.pos = newPos
				if _, ok := seen[guard]; ok{
					count += 1
					break
				} 
				seen[guard] = true
			}
		}	
	}
	fmt.Println("Sol2", count)
}