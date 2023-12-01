package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func unionCount(s1 string, s2 string) int {
	count := 0
	for _, c1 := range s1 {
		for _, c2 := range s2 {
			if c1 == c2 {
				count++
			}
		}
	}
	return count
}

func main() {

	lines := readLines("input.txt")

	total_sum := 0
	for _, line := range lines {

		splitted_line := strings.Split(line, "|")
		signal_patterns := strings.Fields(splitted_line[0])
		output_values := strings.Fields(splitted_line[1])

		patterns := make([]string, 10)

		/*
				0:      1:      2:      3:      4:
				aaaa    ....    aaaa    aaaa    ....
			   b    c  .    c  .    c  .    c  b    c
			   b    c  .    c  .    c  .    c  b    c
				....    ....    dddd    dddd    dddd
			   e    f  .    f  e    .  .    f  .    f
			   e    f  .    f  e    .  .    f  .    f
				gggg    ....    gggg    gggg    ....

				 5:      6:      7:      8:      9:
				aaaa    aaaa    aaaa    aaaa    aaaa
			   b    .  b    .  .    c  b    c  b    c
			   b    .  b    .  .    c  b    c  b    c
				dddd    dddd    ....    dddd    dddd
			   .    f  e    f  .    f  e    f  .    f
			   .    f  e    f  .    f  e    f  .    f
				gggg    gggg    ....    gggg    gggg
		*/
		// Find 1,4,7,8 patterns
		for _, signal_pattern := range signal_patterns {
			switch len(signal_pattern) {
			case 2:
				patterns[1] = signal_pattern
			case 3:
				patterns[7] = signal_pattern
			case 4:
				patterns[4] = signal_pattern
			case 7:
				patterns[8] = signal_pattern
			}
		}

		// Find rest of the patterns
		for _, signal_pattern := range signal_patterns {
			if len(signal_pattern) == 5 {
				// Find 3 by union with 1
				if unionCount(signal_pattern, patterns[1]) == 2 {
					patterns[3] = signal_pattern
					// Find 2, 5  by union with 4
				} else if unionCount(signal_pattern, patterns[4]) == 2 {
					patterns[2] = signal_pattern
				} else {
					patterns[5] = signal_pattern
				}
			} else if len(signal_pattern) == 6 {
				// Find 6 by union with 1
				if unionCount(signal_pattern, patterns[1]) == 1 {
					patterns[6] = signal_pattern
					// Find 0, 9  by union with 4
				} else if unionCount(signal_pattern, patterns[4]) == 3 {
					patterns[0] = signal_pattern
				} else {
					patterns[9] = signal_pattern
				}
			}
		}

		// Calculate output value
		digits := make(map[string]int)
		for i, pattern := range patterns {
			digits[sortStringByCharacter(pattern)] = i
		}
		sum := 0
		for i, output_value := range output_values {
			output_pattern := sortStringByCharacter(output_value)
			decoded := digits[output_pattern]
			sum += int(math.Pow(float64(10), float64(len(output_values)-i-1))) * decoded
		}
		total_sum += sum
	}
	fmt.Println(total_sum)
}
