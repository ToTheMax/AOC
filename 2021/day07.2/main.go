package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

func calculate_fuel(positions int) int {
	result := 0
	for i := 1; i <= positions; i++ {
		result += i
	}
	return result
}

func main() {

	max_pos := 2000

	lines := readLines("input.txt")
	splitted_line := strings.Split(lines[0], ",")
	positions := make([]int, len(splitted_line))
	position_counts := make([]int, max_pos)
	for i, val := range splitted_line {
		position, _ := strconv.Atoi(val)
		positions[i] = position
		position_counts[position]++
	}

	// Check each horizontal position
	best_total_fuel := math.MaxInt32
	for horz_pos := 0; horz_pos < max_pos; horz_pos++ {
		total_fuel := 0
		for pos, pos_count := range position_counts {
			total_fuel += calculate_fuel(int(math.Abs(float64(horz_pos-pos)))) * pos_count
		}
		if total_fuel < best_total_fuel {
			best_total_fuel = total_fuel
		}
	}
	fmt.Println(best_total_fuel)
}
