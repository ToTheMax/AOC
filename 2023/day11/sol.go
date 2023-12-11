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
}

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

func getDistance(p1 Pos, p2 Pos, rowExpansions []int, colExpansions []int) int {
	distance := abs(p1.x - p2.x) + abs(p1.y - p2.y)
	for _, r := range rowExpansions {
		if (p1.y < r && r < p2.y) || (p2.y < r && r < p1.y){
			distance += 1
		}
	}
	for _, c := range colExpansions {
		if (p1.x < c && c < p2.x) || (p2.x < c && c < p1.x){
			distance += 1
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
	for y, row := range rows{
		for x, char := range row{
			cols[x] += string(char)
			if char == '#'{
				galaxies = append(galaxies, Pos{x, y})
			}
		}
	}

	rowExpansions := getExpansions(rows)
	colExpansions := getExpansions(cols)


	totalDist := 0

	// Loop all pairs of galaxies
	for i, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies[:i]{
			if galaxy1 != galaxy2{
				dist := getDistance(galaxy1, galaxy2, rowExpansions, colExpansions)
				// fmt.Println("From", i+1, "to", j+1, dist)
				totalDist += dist
			}
		}
	}

	fmt.Println("Sol 1:", totalDist)
	fmt.Println("Sol 2:", 0)
	// too high 742306702870
}
