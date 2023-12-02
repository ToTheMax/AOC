package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	sumId := 0
	sumPower := 0

	for _, line := range lines {
		gameID := 0
		colorStrings := ""
		fmt.Sscanf(line, "Game %d: %s", &gameID, &colorStrings)
		sets := strings.Split(line, ": ")[1]
		possible := true
		colorMax := make(map[string]int)
		for _, set := range strings.Split(sets, "; ") {
			colorCounts := make(map[string]int)
			for _, color := range strings.Split(set, ", ") {
				splitted := strings.Split(color, " ")
				amount, _ := strconv.Atoi(splitted[0])
				color := splitted[1]
				colorCounts[color] += amount
				if colorMax[color] < amount {
					colorMax[color] = amount
				}
			}
			if colorCounts["red"] > 12 || colorCounts["green"] > 13 || colorCounts["blue"] > 14 {
				possible = false
			}
		}
		if possible {
			sumId += gameID
		}
		power := 1
		for max := range colorMax {
			power *= colorMax[max]
		}
		sumPower += power
	}
	fmt.Println("Sol 1:", sumId)
	fmt.Println("Sol 2:", sumPower)
}
