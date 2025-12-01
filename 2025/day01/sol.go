package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Rotation struct {
	Number    int
	Direction Direction
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	rotations := make([]Rotation, 0, len(lines))

	for _, line := range lines {
		var dirChar string
		var number int
		fmt.Sscanf(line, "%1s%d", &dirChar, &number)
		var direction Direction
		if dirChar == "L" {
			direction = Left
		} else {
			direction = Right
		}
		rotations = append(rotations, Rotation{
			Number:    number,
			Direction: direction,
		})
	}

	current_digit := 50
	score1 := 0
	score2 := 0

	for _, rotation := range rotations {
		if rotation.Direction == Left {
			score2 += int(math.Abs(float64(current_digit-rotation.Number))) / 100
			if current_digit != 0 && current_digit <= rotation.Number {
				score2 += 1
			}
			current_digit = ((current_digit-rotation.Number)%100 + 100) % 100
		} else {
			score2 += (current_digit + rotation.Number) / 100
			current_digit = (current_digit + rotation.Number) % 100
		}

		if current_digit%100 == 0 {
			score1 += 1
		}
	}

	fmt.Println("Sol 1:", score1)
	fmt.Println("Sol 2:", score2)
}
