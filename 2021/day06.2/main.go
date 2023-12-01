package main

import (
	"bufio"
	"fmt"
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

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {

	lines := readLines("input.txt")
	new_lf_timers := make([]int, 9)
	old_lf_timers := make([]int, 7)
	for _, val := range strings.Split(lines[0], ",") {
		initial_state, _ := strconv.Atoi(val)
		new_lf_timers[initial_state]++
	}

	for day := 1; day <= 256; day++ {

		// Save zero timers
		new_lf_timer_zero := new_lf_timers[0]
		old_lf_timer_zero := old_lf_timers[0]

		// Simulate new LF
		for i := 1; i < len(new_lf_timers); i++ {
			new_lf_timers[i-1] = new_lf_timers[i]
		}

		// Simulate old LF
		for i := 1; i < len(old_lf_timers); i++ {
			old_lf_timers[i-1] = old_lf_timers[i]
		}

		// Spawn children
		new_lf_timers[8] = new_lf_timer_zero + old_lf_timer_zero
		old_lf_timers[6] = new_lf_timer_zero + old_lf_timer_zero
	}

	fmt.Println(sum(new_lf_timers) + sum(old_lf_timers))
}
