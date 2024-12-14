package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	stone  int
	blinks int
}

func count_stones(stone int, blinks int, cache map[State]int) int {
	if blinks == 0 {
		return 1
	}
	if cached_blinks, ok := cache[State{stone, blinks}]; ok {
		return cached_blinks
	} else {
		stoneStr := strconv.Itoa(stone)
		result := 0
		if stone == 0 {
			result = count_stones(1, blinks-1, cache)
		} else if len(stoneStr)%2 == 0 {
			leftStone, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
			rightStone, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
			result = count_stones(leftStone, blinks-1, cache) + count_stones(rightStone, blinks-1, cache)
		} else {
			result = count_stones(stone*2024, blinks-1, cache)
		}
		cache[State{stone, blinks}] = result
		return result
	}
}

func solve(stones []int, blinks int) int {
	cache := make(map[State]int)
	total := 0
	for _, stone := range stones {
		total += count_stones(stone, blinks, cache)
	}
	return total
}

func main() {
	input, _ := os.ReadFile("in.txt")
	arrangement := string(input)

	stones := make([]int, len(strings.Split(arrangement, " ")))
	for i, s := range strings.Split(arrangement, " ") {
		stones[i], _ = strconv.Atoi(s)
	}

	fmt.Println("Sol 1:", solve(stones, 25))
	fmt.Println("Sol 2:", solve(stones, 75))
}
