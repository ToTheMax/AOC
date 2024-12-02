package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	left_list := make([]int, len(lines))
	right_list := make([]int, len(lines))

	for i, line := range lines {
		var left, right int
		fmt.Sscanf(line, "%d %d", &left, &right)
		left_list[i] = left
		right_list[i] = right
	}

	// Sort lists for Part 1
	sort.Ints(left_list)
	sort.Ints(right_list)

	// Make frequency map for Part 2
	right_counts := make(map[int]int)
	for _, num := range right_list {
		right_counts[num]++
	}

	total_distance := 0
	similarity_score := 0

	for i := 0; i < len(lines); i++ {
		total_distance += int(math.Abs(float64(left_list[i]) - float64(right_list[i])))
		similarity_score += left_list[i] * right_counts[left_list[i]]
	}

	fmt.Println("Sol 1:", total_distance)
	fmt.Println("Sol 2:", similarity_score)
}
