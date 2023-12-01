package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("in.txt")
	text := string(input)
	inventories := strings.Split(text, "\n\n")
	sums := make([]int, len(inventories))

	for i, inventory := range inventories {
		items := strings.Split(inventory, "\n")
		count := 0
		for _, item := range items {
			j, _ := strconv.Atoi(item)
			count += j
		}
		sums[i] = count
	}
	sort.Ints(sums)
	fmt.Println("Sol 1:", sums[len(sums)-1])
	fmt.Println("Sol 2:", sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3])
}
