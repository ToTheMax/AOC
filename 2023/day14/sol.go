package main

import (
	"fmt"
	"os"
	"strings"
)

func tilt(platform []string, fullCycle bool) ([]string, int) {

	n := len(platform)
	m := len(platform[0])

	tilts := 1
	if fullCycle {
		tilts = 4
	}

	score := 0
	newPlatform := make([]string, n)
	for t := 0; t < tilts; t++ {
		// New cycle
		score = 0
		newPlatform = make([]string, n)
		for i := 0; i < n; i++ {
			newPlatform[i] = strings.Repeat(".", m)
		}
		for x := 0; x < m; x++ {
			top := m
			for y := 0; y < n; y++ {
				if platform[y][x] == '#' {
					top = m - (y + 1)
					newPlatform[x] = newPlatform[x][:m-y-1] + "#" + newPlatform[x][m-y:]
					continue
				} else if platform[y][x] == 'O' {
					score += top
					top = top - 1
					newPlatform[x] = newPlatform[x][:top] + "O" + newPlatform[x][top+1:]
				}
			}
		}
		platform = newPlatform
	}
	return newPlatform, score
}

func getScore(platform []string) int {
	score := 0
	for i, row := range platform {
		for _, c := range row {
			if c == 'O' {
				score += len(row) - i
			}
		}
	}
	return score
}

func main() {
	input, _ := os.ReadFile("in.txt")
	platform := strings.Split(string(input), "\n")

	// Part 1
	_, solP1 := tilt(platform, false)
	fmt.Println("Sol 1:", solP1)

	// Part 2
	cycles := 100000

	// Find pattern
	patternStart, patternSize := 0, 0
	seenPlatforms := make(map[string]int)
	for c := 0; c < cycles; c++ {
		platform, _ = tilt(platform, true)
		platform_str := strings.Join(platform, "")
		if c2, ok := seenPlatforms[platform_str]; ok {
			patternStart = c2 + 1
			patternSize = c - c2
			break
		} else {
			seenPlatforms[platform_str] = c
		}
	}

	// Find diffs
	diffs := make([]int, patternSize)
	scoreP2 := getScore(platform)
	score, prevScore := 0, scoreP2
	for i := 0; i < patternSize; i++ {
		platform, score = tilt(platform, true)
		score = getScore(platform)
		diffs[i] = score - prevScore
		prevScore = score
	}
	for d := 0; d < (cycles-patternStart)%len(diffs); d++ {
		scoreP2 += diffs[d]
	}

	fmt.Println("Sol 2:", scoreP2)
}
