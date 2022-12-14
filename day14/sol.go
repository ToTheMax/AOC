package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Point struct {
	X int
	Y int
}

func addRocks(rocks map[Point]bool, p1 Point, p2 Point) int {
	xs := []int{p1.X, p2.X}
	ys := []int{p1.Y, p2.Y}
	sort.Ints(ys)
	sort.Ints(xs)
	for x := xs[0]; x < xs[1]+1; x++ {
		for y := ys[0]; y < ys[1]+1; y++ {
			rocks[Point{x, y}] = true
		}
	}
	return ys[1]
}

func dropSand(p Point, sand map[Point]bool, rocks map[Point]bool, border int) bool {
	if p.Y > border {
		return false
	}
	// Fall
	for _, dest := range []Point{{p.X, p.Y + 1}, {p.X - 1, p.Y + 1}, {p.X + 1, p.Y + 1}} {
		if !rocks[dest] && !sand[dest] {
			return dropSand(dest, sand, rocks, border)
		}
	}
	// Rest
	sand[p] = true
	return p != Point{500, 0}
}

func main() {
	input, _ := os.ReadFile("in.txt")

	lines := strings.Split(string(input), "\n")

	// Make rocks
	rocks := make(map[Point]bool)
	voidBorder := 0
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		for i := 0; i < len(points)-1; i++ {
			var x1, y1, x2, y2 int
			fmt.Sscanf(points[i], "%d,%d:", &x1, &y1)
			fmt.Sscanf(points[i+1], "%d,%d:", &x2, &y2)
			if lowestRock := addRocks(rocks, Point{x1, y1}, Point{x2, y2}); lowestRock > voidBorder {
				voidBorder = lowestRock
			}
		}
	}

	// Pour sand
	sand := make(map[Point]bool)
	source := Point{500, 0}
	for dropSand(source, sand, rocks, voidBorder) {
	}

	fmt.Println("Sol1:", len(sand))

	addRocks(rocks, Point{-1000, voidBorder + 2}, Point{1000, voidBorder + 2})
	for dropSand(source, sand, rocks, voidBorder+2) {
	}
	fmt.Println("Sol2:", len(sand))
}
