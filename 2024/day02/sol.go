package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInts(s string) []int {
	var nums []int
	for _, numStr := range strings.Fields(s) {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	return nums
}

func reportIsSafe(report []int, m int, violations_allowed int) bool {
	violations := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff*m < 1 || diff*m > 3 {
			violations++

			if violations > violations_allowed {
				return false
			}

			// Check if we can remove first level
			if i == 0 {
				second_diff := report[i+2] - report[i+1]
				if !(second_diff*m < 1 || second_diff*m > 3) {
					continue
				}
			}
			// Always possible to remove last level
			if i+2 >= len(report) {
				return true
			}

			next_diff := report[i+2] - report[i]
			if next_diff*m < 1 || next_diff*m > 3 {
				return false
			}
			i++
		}
	}
	return true
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	p1_score := 0
	p2_score := 0

	for _, line := range lines {
		report := stringToInts(line)
		if reportIsSafe(report, -1, 0) || reportIsSafe(report, 1, 0) {
			p1_score++
		}
		if reportIsSafe(report, -1, 1) || reportIsSafe(report, 1, 1) {
			p2_score++
		}
	}

	fmt.Println("Sol 1:", p1_score)
	fmt.Println("Sol 2:", p2_score)
}
