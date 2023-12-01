package main

import (
	"fmt"
	"os"
	"strings"
)

func findIntersectChar(strings []string) rune {
	for _, c1 := range strings[0] {
		count := 1
		for _, str := range strings[1:] {
			for _, c2 := range str {
				if c1 == c2 {
					count += 1
					if count == len(strings) {
						return c1
					} else {
						break
					}
				}
			}
		}
	}
	return 0
}

func charScore(char rune) int {
	if int(char) > int('a') {
		return int(char) - int('a') + 1
	} else {
		return int(char) - int('A') + 27
	}
}

func main() {
	input, _ := os.ReadFile("in.txt")
	rucksacks := strings.Split(string(input), "\n")

	score := 0
	for _, rucksack := range rucksacks {

		first := rucksack[:len(rucksack)/2]
		second := rucksack[len(rucksack)/2:]
		intersectChar := findIntersectChar([]string{first, second})
		score += charScore(intersectChar)
	}
	fmt.Println("Sol1:", score)

	score = 0
	for i := 0; i < len(rucksacks); i = i + 3 {
		intersectChar := findIntersectChar(rucksacks[i : i+3])
		score += charScore(intersectChar)
	}
	fmt.Println("Sol2:", score)
}
