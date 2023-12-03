package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scheme struct {
	schematic []string
	height    int
	width     int
	gears     map[int]Gear
}

type Gear struct {
	count int
	ratio int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

func isSymbol(char rune) bool {
	return !isDigit(char) && char != '.'
}

func isDigit(char rune) bool {
	ascii_num := int(char - '0')
	return ascii_num >= 0 && ascii_num <= 9
}

func checkAdjacent(scheme Scheme, row int, from int, to int, number int) bool {
	foundSymbol := false
	for y := max(row-1, 0); y <= min(row+1, scheme.height-1); y++ {
		for x := max(from-1, 0); x <= min(to+1, scheme.width-1); x++ {
			if isSymbol(rune(scheme.schematic[y][x])) {
				foundSymbol = true
				if rune(scheme.schematic[y][x]) == '*' {
					coordinate := y*scheme.width + x
					// check if coordinate exists in gears
					if gear, ok := scheme.gears[coordinate]; ok {
						gear.count++
						gear.ratio *= number
						scheme.gears[coordinate] = gear
					} else {
						scheme.gears[coordinate] = Gear{ratio: number, count: 1}
					}
				}
			}
		}
	}
	return foundSymbol
}

func main() {

	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")
	scheme := Scheme{lines, len(lines), len(lines[0]), map[int]Gear{}}

	sumParts := 0

	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := line[j]
			if isDigit(rune(char)) {
				from := j
				to := j
				for x := j; x < scheme.width; x++ {
					if isDigit(rune(scheme.schematic[i][x])) {
						to = x
					} else {
						break
					}
				}
				number, _ := strconv.Atoi(scheme.schematic[i][from : to+1])
				// fmt.Println("\tRow", i, "From", from, "To", to)
				// fmt.Println("\tFound number", string(scheme.schematic[i][from:to]))
				if checkAdjacent(scheme, i, from, to, number) {
					// fmt.Println("\tFound number:", number)
					sumParts += number
				}
				j += to - from
			}
		}
	}
	fmt.Println("Sol 1:", sumParts)

	sumGears := 0
	for _, gear := range scheme.gears {
		if gear.count == 2 {
			// fmt.Println(gear.ratio)
			sumGears += gear.ratio
		}
	}
	fmt.Println("Sol 2:", sumGears)
}
