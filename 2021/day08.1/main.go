package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) []string {
	file, _ := os.Open(path)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {

	lines := readLines("input.txt")
	count := 0

	for _, line := range lines {

		splitted_line := strings.Split(line, "|")
		output_values := strings.Fields(splitted_line[1])
		for _, output_value := range output_values {
			l := len(output_value)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}
