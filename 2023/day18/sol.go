package main

import (
	"fmt"
	"os"
	"strconv"
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcArea(turningPos []Pos) int {
	area := 0
	for i := 0; i < len(turningPos); i++ {
		pointA := turningPos[i]
		pointB := turningPos[(i+1)%(len(turningPos))]
		area += (pointA.x * pointB.y) - (pointB.x * pointA.y) 
		area += max(abs(pointA.x-pointB.x), abs(pointA.y-pointB.y))
	}
	return area / 2
}

func main() {
	input, _ := os.ReadFile("in.txt")

	// Preprocess
	lines := strings.Split(string(input), "\n")
	curPos1 := Pos{0, 0}
	curPos2 := Pos{0, 0}
	turningPos1 := []Pos{{0, 0}}
	turningPos2 := []Pos{{0, 0}}
	for _, line := range lines {
		split := strings.Split(line, " ")
		// P1
		dir1 := split[0]
		dist1, _ := strconv.Atoi(split[1])
		for i := 0; i < dist1; i++ {
			switch dir1 {
			case "D":
				curPos1.y++
			case "U":
				curPos1.y--
			case "R":
				curPos1.x++
			case "L":
				curPos1.x--
			}
			turningPos1 = append(turningPos1, curPos1)
		}
		// P2
		color := split[2][2 : len(split[2])-1]
		dist2, _ := (strconv.ParseInt(color[:5], 16, 0))
		dir2 := color[5]
		switch dir2 {
		case '0':
			curPos2.x += int(dist2)
		case '1':
			curPos2.y += int(dist2)
		case '2':
			curPos2.x -= int(dist2)
		case '3':
			curPos2.y -= int(dist2)
		}
		turningPos2 = append(turningPos2, curPos2)
	}

	fmt.Println("Sol 1", calcArea(turningPos1)+1)
	fmt.Println("Sol 2:", calcArea(turningPos2)+1)
}
