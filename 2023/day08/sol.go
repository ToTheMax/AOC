package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

func calcSteps(network map[string]Node, instructs string, from string, to string) int {
	steps := 0
	curNode := network[from]
	for curNode.name != to {
		if instructs[steps%len(instructs)] == 'R' {
			curNode = network[curNode.right]
		} else {
			curNode = network[curNode.left]
		}
		steps++
		// Cycle
		if steps > 100000 {
			return -1
		}
	}

	return steps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcmOfList(nums []int) int {
	lcmResult := nums[0]
	for _, num := range nums[1:] {
		lcmResult = (lcmResult / gcd(lcmResult, num)) * num 
	}
	return lcmResult
}

func main() {

	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	instructions := lines[0]
	network := make(map[string]Node)

	for _, line := range lines[2:] {
		l := strings.Split(line, " = ")
		lr := strings.Split(l[1][1:9], ", ")
		node := Node{l[0], lr[0], lr[1]}
		network[l[0]] = node
	}

	// Find steps AAA -> ZZZ
	fmt.Println("Sol 1:", calcSteps(network, instructions, "AAA", "ZZZ"))

	// Find steps **A -> **Z
	rates := []int{}
	for _, n1 := range network {
		for _, n2 := range network {
			if n1.name[2] == 'A' && n2.name[2] == 'Z'{
				rate := calcSteps(network, instructions, n1.name, n2.name)
				if rate > 0 {
					rates = append(rates, rate)
				}
			}
		}
	}
	fmt.Println("Sol 2:", lcmOfList(rates))
}
