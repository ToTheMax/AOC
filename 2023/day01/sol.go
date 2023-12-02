package main

import (
	"fmt"
	"os"
	"strings"
)

func replace(text string, level int) string {
	digitstrings := ""
	if level == 2 {
		text = strings.ReplaceAll(text, "one", "1")
		text = strings.ReplaceAll(text, "two", "2")
		text = strings.ReplaceAll(text, "three", "3")
		text = strings.ReplaceAll(text, "four", "4")
		text = strings.ReplaceAll(text, "five", "5")
		text = strings.ReplaceAll(text, "six", "6")
		text = strings.ReplaceAll(text, "seven", "7")
		text = strings.ReplaceAll(text, "eight", "8")
		text = strings.ReplaceAll(text, "nine", "9")
	}

	for _, char := range text {
		ascii_num := int(char - '0')
		if ascii_num > 0 && ascii_num <= 9 || char == '\n' {
			digitstrings += string(char)
		}
	}
	return digitstrings
}

func score(text string) int {
	score := 0
	for _, line := range strings.Split(text, "\n") {
		score += int(line[0]-'0')*10 + int(line[len(line)-1]-'0')
	}
	return score
}

func main() {
	input, _ := os.ReadFile("in.txt")
	text := string(input)

	fmt.Println("Sol 1:", score(replace(text, 1)))
	fmt.Println("Sol 2:", score(replace(text, 2)))
}
