package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func stringToInts(s string) []int {
	splitted := strings.Split(s, ",")
	ints := make([]int, len(splitted))
	for i, s := range splitted {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}


func findCombinations(conditions string, index int, groups []int) int{

	combs := 0

	if len(groups) == 0 {
		fitted := strings.ReplaceAll(string(conditions), "?", ".")
		if index < len(conditions) && strings.Contains(conditions[index:], "#"){
			return 0
		} else{
			fmt.Println("Found fit", fitted)
			return 1
		}
	}

	group := groups[0]

	// Check how many conditions left
	// This can be optimized (sum of groups + cnt)
	if index + group > len(conditions){
		return 0
	}

	// Check if group fits in next sequence
	foundFit := true
	if index-1 >= 0  && conditions[index-1] == '#' {
		foundFit = false
	}
	for i:=0; i < group; i++{
		if conditions[index+i] == '.' {
			foundFit = false
		}
	}
	if index+group < len(conditions) && conditions[index+group] == '#' {
		foundFit = false
	}

	if foundFit {
		newConditions := make([]byte, len(conditions))
		copy(newConditions, conditions)
		for i:=0; i<group; i++{
			newConditions[index+i] = '#'
		}

		if index+group < len(conditions){
			newConditions[index+group] = '.'
		}
		combs += findCombinations(string(newConditions), index + group+1, groups[1:])
	}

	if conditions[index] == '?' {
		conditionsBytes := []byte(conditions)
		conditionsBytes[index] = '.'
		conditions = string(conditionsBytes)
	}
	if conditions[index] != '#' {
		return combs + findCombinations(conditions, index+1, groups)
	}
	return combs
}

func main() {

	// Read input
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	sumCombs := 0
	for i, line := range lines {
		splitted_line := strings.Split(line, " ")
		conditions := splitted_line[0]
		groups := stringToInts(splitted_line[1])
		combs := findCombinations(conditions, 0, groups)
		fmt.Println("Line", i, ":", combs)
		sumCombs += combs
	}


	// 6871
	// 2043098029844

	fmt.Println("Sol 1:", sumCombs) // 8376 too high
	fmt.Println("Sol 2:", 0)
}
