package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkExpansion(line string) bool {
	for _, char := range line{
		if char == '#'{
			return false
		}
	}
	return true
}

func getExpansions(lines []string) []int {
	expansions := make([]int, 0)
	for i, line := range lines {
		if checkExpansion(line){
			expansions = append(expansions, i)
		}
	}
	return expansions
}

func isBetween(a, b, c int) bool {
	return (a < b && b < c) || (c < b && b < a)
}

func getDistance(
	p1, p2 Pos, 
	rowExpansions, colExpansions []int,
	expansionSize int,
	) int {
	distance := abs(p1.x - p2.x) + abs(p1.y - p2.y)
	for _, r := range rowExpansions {
		if isBetween(p1.y, r, p2.y){
			distance += expansionSize
		}
	}
	for _, c := range colExpansions {
		if isBetween(p1.x, c, p2.x){
			distance += expansionSize
		}
	}
	return distance
}

func main() {

	// Read input
	input, _ := os.ReadFile("in.txt")
	rows := strings.Split(string(input), "\n")
	cols := make([]string, len(rows[0]))
	galaxies := make([]Pos, 0)
	for y, row := range rows {
		for x, char := range row{
			cols[x] += string(char)
			if char == '#'{
				galaxies = append(galaxies, Pos{x, y})
			}
		}
	}

	rowExpansions := getExpansions(rows)
	colExpansions := getExpansions(cols)

	// Loop all pairs of galaxies
	solP1 := 0
	solP2 := 0
	for i, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies[:i]{
			if galaxy1 != galaxy2{
				solP1 += getDistance(galaxy1, galaxy2, rowExpansions, colExpansions, 1)
				solP2 += getDistance(galaxy1, galaxy2, rowExpansions, colExpansions, 1000000-1)
			}
		}
	}
	fmt.Println("Sol 1:", solP1)
	fmt.Println("Sol 2:", solP2)
}
