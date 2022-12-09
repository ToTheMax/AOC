package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	// id     int
	height int
	seen   bool
}

func reverse(trees []*Tree) []*Tree {
	for i := 0; i < len(trees)/2; i++ {
		j := len(trees) - i - 1
		trees[i], trees[j] = trees[j], trees[i]
	}
	return trees
}

func checkLine(treeLine []*Tree) int {
	score := 0
	maxHeight := -1
	for _, tree := range treeLine {
		if tree.height > maxHeight {
			maxHeight = tree.height
			if !tree.seen {
				tree.seen = true
				score += 1
			}
		}
	}
	return score
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	horzSize := len(lines)
	vertSize := len(lines[0])

	// Read horizontal treelines
	var horzTreeLines [][]*Tree
	for _, line := range lines {
		horzTreeline := make([]*Tree, horzSize)
		for x, char := range line {
			n, _ := strconv.Atoi(string(char))
			tree := &Tree{
				height: n,
				seen:   false,
			}
			horzTreeline[x] = tree
		}
		horzTreeLines = append(horzTreeLines, horzTreeline)
	}

	// Read vertical treelines
	var vertTreeLines [][]*Tree
	for x := 0; x < horzSize; x++ {
		vertTreeLine := make([]*Tree, vertSize)
		for y, horzTreeLine := range horzTreeLines {
			vertTreeLine[y] = horzTreeLine[x]
		}
		vertTreeLines = append(vertTreeLines, vertTreeLine)
	}

	var treeLines [][]*Tree
	treeLines = append(treeLines, horzTreeLines...)
	treeLines = append(treeLines, vertTreeLines...)

	score := 0
	for _, treeLine := range treeLines {
		score += checkLine(treeLine) + checkLine(reverse(treeLine))
	}
	fmt.Println("Sol1:", score)

}
