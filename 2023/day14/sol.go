package main

import (
	"fmt"
	"os"
	"strings"
)

func tilt(rows []string, cycles int) int {

	n := len(rows)
	m := len(rows[0])

	// New cucle
	newCycle := make([]string, n)
	for i := 0; i < n; i++ {
		newCycle[i] = strings.Repeat(".", m)
	}

	score := 0
	for x := 0; x < m; x++ {
		top := m
		for y := 0; y < n; y++ {
			if rows[y][x] == '#' {
				top = m - (y + 1)
				// newLines[y] = newLines[y][:x] + "#" + newLines[y][x+1:]
				// newCycle[x] = newCycle[x][:y] + "#" + newCycle[x][y+1:]
				newCycle[x] = newCycle[x][:m-y-1] + "#" + newCycle[x][m-y:]
				continue
			} else if rows[y][x] == 'O' {
				score += top
				top = top - 1
				// newLines[m-top-1] = newLines[m-top-1][:x] + "O" + newLines[m-top-1][x+1:]
				// newCycle[x] = newCycle[x][:m-top-1] + "O" + newCycle[x][m-top:]
				newCycle[x] = newCycle[x][:top] + "O" + newCycle[x][top+1:]
			}
		}
	}
	if cycles > 0 {
		if newCycle[0] == rows[0] {
			return score
		}
		// for _, line := range newCycle {
		// 	fmt.Println(line)
		// }
		// fmt.Println()
		return tilt(newCycle, cycles-1)
	}
	return score
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	fmt.Println("Sol 1:", tilt(lines, 0))

	// 129
	// 110
	fmt.Println("Sol 2:", tilt(lines, 4*1000000000))
}
