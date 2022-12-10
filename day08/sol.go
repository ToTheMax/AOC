package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	height      int
	seen        bool
	scenicScore int
}

func reverseTree(trees []*Tree) []*Tree {
	reversed := make([]*Tree, len(trees))
	for i := 0; i < len(trees); i++ {
		j := len(trees) - i - 1
		reversed[i] = trees[j]
	}
	return reversed
}

func scoreTree(tree *Tree, neighbours []*Tree) int {
	viewingDistance := 0
	for _, neighbourTree := range neighbours {
		viewingDistance += 1
		if neighbourTree.height >= tree.height {
			break
		}
	}
	return viewingDistance
}

func scoreTreeLine(treeLine []*Tree) int {
	score := 0
	maxHeight := -1
	for i, tree := range treeLine {
		if tree.height > maxHeight {
			maxHeight = tree.height
			if !tree.seen {
				tree.seen = true
				score += 1
			}
		}
		tree.scenicScore *= scoreTree(tree, reverseTree(treeLine[:i]))
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
				height:      n,
				seen:        false,
				scenicScore: 1,
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

	// Merge treelines
	var treeLines [][]*Tree
	treeLines = append(treeLines, horzTreeLines...)
	treeLines = append(treeLines, vertTreeLines...)

	// Part 1
	score := 0
	for _, treeLine := range treeLines {
		score += scoreTreeLine(treeLine) + scoreTreeLine(reverseTree(treeLine))
	}
	fmt.Println("Sol1:", score)

	// Part 2
	maxScenicScore := 0
	for _, treeLine := range horzTreeLines {
		for _, tree := range treeLine {
			if tree.scenicScore > maxScenicScore {
				maxScenicScore = tree.scenicScore
			}
		}
	}
	fmt.Println("Sol2:", maxScenicScore)
}
