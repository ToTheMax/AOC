package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkRule(rules map[int][]int, page_numbers []int, page_indices map[int]int) bool {
	for _, page_number := range page_numbers {
		for _, after := range rules[page_number] {
			if i, ok := page_indices[after]; ok && page_indices[page_number] > i {
				return false
			}
		}
	}
	return true
}

func main() {
	input, _ := os.ReadFile("in.txt")
	input_split := strings.Split(string(input), "\n\n")
	rule_lines := strings.Split(string(input_split[0]), "\n")
	page_lines := strings.Split(string(input_split[1]), "\n")

	rules := make(map[int][]int)
	for _, line := range rule_lines {
		nums := strings.Split(line, "|")
		left, _ := strconv.Atoi(nums[0])
		right, _ := strconv.Atoi(nums[1])
		rules[left] = append(rules[left], right)
	}

	scoreP1 := 0
	scoreP2 := 0

	for _, line := range page_lines {
		nums := strings.Split(line, ",")
		page_numbers := make([]int, len(nums))
		page_indices := make(map[int]int)
		for i, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			page_numbers[i] = num
			page_indices[num] = i
		}

		if checkRule(rules, page_numbers, page_indices) {
			scoreP1 += page_numbers[len(page_numbers)/2]
		} else {
			scoreP2 += 0 // TODO order the wrong ones
		}
	}
	fmt.Println("Sol 1:", scoreP1)
	fmt.Println("Sol 2:", scoreP2)
}
