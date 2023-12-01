package main

import (
	"fmt"
	"os"
	"strings"
)

// Rock 0
// Paper 1
// Scissor 2

// Rock - Rock = 0			-> Draw
// Rock - Paper = -1		-> Lose
// Rock - Scissor = -2		-> Win
// Paper - Rock = 1			-> Win
// Paper - Paper = 0		-> Draw
// Paper - Scissor = -1		-> Lose
// Scissor - Rock = 2		-> Lose
// Scissor - Paper = 1		-> Win
// Scissor - Scissor = 0 	-> Draw

// -2 = Win
// -1 = Lose
// 0 = Draw
// 1 = Win
// 2 = Lose

func main() {
	input, _ := os.ReadFile("in.txt")
	text := string(input)
	games := strings.Split(text, "\n")

	score := 0
	for _, game := range games {
		items := strings.Split(game, " ")
		hand1 := int(items[0][0]) - int('A')
		hand2 := int(items[1][0]) - int('X')
		result := (hand1 - hand2) % 3
		if result < 0 {
			result += 3
		}
		if result == 0 {
			score += 3
		} else if result == 1 {
			score += 0
		} else if result == 2 {
			score += 6
		}
		score += hand2 + 1
	}
	fmt.Println("Sol 1:", score)

	score = 0
	for _, game := range games {
		items := strings.Split(game, " ")
		hand1 := int(items[0][0]) - int('A')
		ending := int(items[1][0]) - int('X')
		score += ending * 3
		hand2 := 0
		if ending == 0 {
			hand2 = (hand1 + 2) % 3
		} else if ending == 1 {
			hand2 = hand1
		} else if ending == 2 {
			hand2 = (hand1 + 1) % 3
		}
		score += hand2 + 1
	}
	fmt.Println("Sol 2:", score)
}
