package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func (p Pos) isInMap(lines []string) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(lines[0]) && p.y < len(lines)
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	antenas := make(map[rune][]Pos)

	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				antenas[c] = append(antenas[c], Pos{x, y})
			}
		}
	}

	antinodesP1 := make(map[Pos]int)
	antinodesP2 := make(map[Pos]int)

	for _, antena := range antenas {
		for i := 0; i < len(antena); i++ {
			for j := i + 1; j < len(antena); j++ {
				diffx := antena[j].x - antena[i].x
				diffy := antena[j].y - antena[i].y

				// For part 1 just try to add two nodes
				antinode1 := Pos{antena[i].x - diffx, antena[i].y - diffy}
				antinode2 := Pos{antena[i].x + 2*diffx, antena[i].y + 2*diffy}
				if antinode1.isInMap(lines){
					antinodesP1[antinode1] += 1
				}
				if antinode2.isInMap(lines){
					antinodesP1[antinode2] += 1
				}
		
				// For part 2 resonate from the antenna as long as they stay in the map
				antinode1 = antena[i]
				antinode2 = antena[j]
				for antinode1.isInMap(lines){
					antinodesP2[antinode1] += 1
					antinode1 = Pos{antinode1.x - diffx, antinode1.y - diffy}
				}
				for antinode2.isInMap(lines){
					antinodesP2[antinode2] += 1
					antinode2 = Pos{antinode2.x + diffx, antinode2.y + diffy}
				}

			}
		}
	}

	fmt.Println("Sol1:", len(antinodesP1))
	fmt.Println("Sol2:", len(antinodesP2))
}

