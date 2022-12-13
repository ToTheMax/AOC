package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Node struct {
	value    int
	parent   *Node
	children []*Node
}

func parsePacket(input string) []interface{} {
	var array []interface{}
	json.Unmarshal([]byte(input), &array)
	return array
}

func compare(left []interface{}, right []interface{}) int {
	for i := 0; i < int(math.Max(float64(len(left)), float64(len(right)))); i++ {
		// Check if we hit an end
		if i == len(left) {
			return 1
		} else if i == len(right) {
			return 0
		}
		// Compare numbers or expand list
		leftNumber, leftIsNumber := left[i].(float64)
		rightNumber, rightIsNumber := right[i].(float64)

		if leftIsNumber && rightIsNumber {
			if leftNumber > rightNumber {
				return 0
			} else if leftNumber < rightNumber {
				return 1
			}
			continue
		} else if leftIsNumber {
			if x := compare([]interface{}{leftNumber}, right[i].([]interface{})); x > -1 {
				return x
			}
		} else if rightIsNumber {
			if x := compare(left[i].([]interface{}), []interface{}{rightNumber}); x > -1 {
				return x
			}
		} else if x := compare(left[i].([]interface{}), right[i].([]interface{})); x > -1 {
			return x
		}
	}
	return -1
}

func main() {
	input, _ := os.ReadFile("in.txt")

	// Part 1
	sum := 0
	pairs := strings.Split(string(input), "\n\n")
	for i, pair := range pairs {
		splitted := strings.Split(pair, "\n")
		left, right := parsePacket(splitted[0]), parsePacket(splitted[1])
		sum += compare(left, right) * (i + 1)
	}
	fmt.Println("Sol1:", sum)

	// Part 2
	lines := strings.Split(strings.ReplaceAll(string(input), "\n\n", "\n"), "\n")
	divider1 := "[[2]]"
	divider2 := "[[6]]"
	lines = append(lines, []string{divider1, divider2}...)
	sort.Slice(lines, func(x int, y int) bool {
		return compare(parsePacket(lines[x]), parsePacket(lines[y])) == 1
	})
	product := 1
	for i, line := range lines {
		if line == divider1 || line == divider2 {
			product *= i + 1
		}
	}
	fmt.Println("Sol2:", product)
}
