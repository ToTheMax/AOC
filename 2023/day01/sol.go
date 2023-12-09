package main

import (
	"fmt"
	"os"
	"strings"
)

func replace(text string, level int) string {
	digitstrings := ""
	if level == 2 {
		for i := range text {
			mappings := map[string]string{
				"one": "1",
				"two": "2",
				"three": "3",
				"four": "4",
				"five": "5",
				"six": "6",
				"seven": "7",
				"eight": "8",
				"nine": "9",
			}
			for from, to := range mappings {
				if i+len(from) <= len(text) && text[i:i+len(from)] == from {
					text = text[:i] + to + text[i+3:]
				}
			}
		}
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
