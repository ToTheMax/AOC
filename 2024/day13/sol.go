package main

import (
	"fmt"
	"os"
	"strings"
)

type Button struct {
	dx, dy int
}

type Prize struct {
	x, y int
}

func solve(buttonA Button, buttonB Button, prize Prize, part int) int {
	if part == 2 {
		prize.x += 10000000000000
		prize.y += 10000000000000
	}

	det := buttonA.dx*buttonB.dy - buttonB.dx*buttonA.dy
	if det == 0 {
		return 0
	}

	// Solve for a and b
	a := (prize.x*buttonB.dy - buttonB.dx*prize.y) / det
	b := (buttonA.dx*prize.y - prize.x*buttonA.dy) / det

	if a*det != (prize.x*buttonB.dy-buttonB.dx*prize.y) ||
		b*det != (buttonA.dx*prize.y-prize.x*buttonA.dy) ||
		a < 0 || b < 0 {
		return 0
	}

	// For part 1, check if solution is within 100 presses
	if part == 1 && (a > 100 || b > 100) {
		return 0
	}

	return 3*a + b
}

func main() {
	input, _ := os.ReadFile("in.txt")
	splitted_input := strings.Split(string(input), "\n\n")

	solP1 := 0
	solP2 := 0

	for _, machine_input := range splitted_input {
		lines := strings.Split(machine_input, "\n")
		var buttonA, buttonB Button
		var prize Prize
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &buttonA.dx, &buttonA.dy)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &buttonB.dx, &buttonB.dy)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &prize.x, &prize.y)

		solP1 += solve(buttonA, buttonB, prize, 1)
		solP2 += solve(buttonA, buttonB, prize, 2)
	}

	fmt.Println("Sol 1:", solP1)
	fmt.Println("Sol 2:", solP2)
}
