package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

type Movement struct {
	direction rune
	amount    int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcMove(diff int) int {
	if diff > 1 {
		return 1
	} else if diff < -1 {
		return -1
	}
	return 0
}

func positionsVisited(ropeLength int, movements []*Movement) int {
	rope := make([]Position, ropeLength)
	visited := make(map[Position]bool)
	for _, movement := range movements {
		for i := 0; i < movement.amount; i++ {
			// Move head
			switch movement.direction {
			case 'U':
				rope[0].Y++
			case 'R':
				rope[0].X++
			case 'D':
				rope[0].Y--
			case 'L':
				rope[0].X--
			}

			// Move tail
			for j := 0; j < len(rope)-1; j++ {
				// Check horz / vert
				horzDiff := rope[j].X - rope[j+1].X
				vertDiff := rope[j].Y - rope[j+1].Y
				moveTailX := calcMove(horzDiff)
				moveTailY := calcMove(vertDiff)

				// Check diagonal
				if abs(vertDiff)+abs(horzDiff) > 2 {
					if moveTailX == 0 {
						moveTailX += horzDiff
					}
					if moveTailY == 0 {
						moveTailY += vertDiff
					}
				}
				rope[j+1].X += moveTailX
				rope[j+1].Y += moveTailY
			}
			visited[rope[len(rope)-1]] = true
		}
	}
	return len(visited)
}

func main() {

	// Parse input
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")
	movements := make([]*Movement, len(lines))
	for i, line := range lines {
		splitted := strings.Split(line, " ")
		amount, _ := strconv.Atoi(string(splitted[1]))
		movements[i] = &Movement{
			direction: rune(splitted[0][0]),
			amount:    amount,
		}
	}

	fmt.Println("Sol1:", positionsVisited(2, movements))
	fmt.Println("Sol2:", positionsVisited(10, movements))
}
