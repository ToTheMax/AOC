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
		if steps > 100000 {
			return -1
		}
	}

	return steps
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
	// startNodes := []Node{}
	// endNodes := []Node{}
	// for _, node := range network {
	// 	if node.name[2] == 'A' {
	// 		startNodes = append(startNodes, node)
	// 	} else if node.name[2] == 'Z' {
	// 		endNodes = append(endNodes, node)
	// 	}
	// }
	// fmt.Println(len(startNodes), len(endNodes))
	// freqs := make([][]int, len(startNodes))
	// for s, startNodes := range startNodes {
	// 	freqs[s] = make([]int, len(endNodes))
	// 	for e, endNodes := range endNodes {
	// 		freqs[s][e] = calcSteps(network, instructions, startNodes.name, endNodes.name)
	// 	}
	// }
	fmt.Println("Sol 2:")

}
