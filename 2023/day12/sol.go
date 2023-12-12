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


func findCombinations(conditions string, cIndex int, groups []int, gIndex int) int {

	combs := 0

	if cIndex > len(conditions) && gIndex +1 < len(groups){
		return 0
	}
	
	if gIndex >= len(groups) {
		if cIndex < len(conditions) && strings.Contains(conditions[cIndex:], "#"){
			return 0
		} else{
			return 1
		}
	}

	group := groups[gIndex]

	// Check how many conditions left
	groupSum := 0
	for i:=gIndex; i < len(groups); i++{
		groupSum += groups[i]
	}

	if cIndex + groupSum > len(conditions){
		return 0
	}

	// Check if group fits in next sequence
	foundFit := true
	if cIndex-1 >= 0  && conditions[cIndex-1] == '#' {
		foundFit = false
	}
	for i:=0; i < group; i++{
		if conditions[cIndex+i] == '.' {
			foundFit = false
		}
	}
	if cIndex+group < len(conditions) && conditions[cIndex+group] == '#' {
		foundFit = false
	}

	if foundFit {
		combs += findCombinations(conditions, cIndex + group+1, groups, gIndex+1)
	}

	if conditions[cIndex] != '#' {
		return combs + findCombinations(conditions, cIndex+1, groups, gIndex)
	}
	return combs
}

func main() {

	// Read input
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	sumP1 := 0
	sumP2 := 0
	for _, line := range lines {
		splitted_line := strings.Split(line, " ")
		conditions := splitted_line[0]
		groups := stringToInts(splitted_line[1])
		combs := findCombinations(conditions, 0, groups, 0)
		sumP1 += combs
	}
	fmt.Println("Sol 1:", sumP1)

	for _, line := range lines{
		splitted_line := strings.Split(line, " ")
		conditions := splitted_line[0]
		groups := stringToInts(splitted_line[1])
		conditions = strings.Repeat(conditions + "?", 5)
		conditions = conditions[:len(conditions)-1]
		mgroups := append(groups, groups...)
		mgroups = append(mgroups, mgroups...)
		mgroups = append(mgroups, groups...)
		combs := findCombinations(conditions, 0, mgroups, 0)
		sumP2 += combs
	}

	fmt.Println("Sol 2:", sumP2)
}
