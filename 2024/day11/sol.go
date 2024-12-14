package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("in.txt")
	arrangement := string(input)

	// convert to ints
	stonesStr := strings.Split(arrangement, " ")
	stones := make([]int, len(stonesStr))
	for i, stone := range stonesStr {
		stones[i], _ = strconv.Atoi(stone)
	}

	// Blink
	for b := 0; b < 25; b++ {
		newStones := make([]int, len(stones)*2)
		newStonesIndex := 0
		// fmt.Println("After", b, "blinks", len(stones))
		for _, stone := range stones {
			stoneStr := strconv.Itoa(stone)
			if stone == 0 {
				newStones[newStonesIndex] = 1
			} else if len(stoneStr)%2 == 0 {
				newStones[newStonesIndex], _ = strconv.Atoi(stoneStr[:len(stoneStr)/2])
				newStones[newStonesIndex+1], _ = strconv.Atoi(stoneStr[len(stoneStr)/2:])
				newStonesIndex++
			} else {
				newStones[newStonesIndex] = stone * 2024
			}
			newStonesIndex++
		}
		stones = newStones[:newStonesIndex]
	}

	fmt.Println("Sol 1:", len(stones))
}
