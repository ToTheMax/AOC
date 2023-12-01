package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	c := 0
	var prev2 = -1
	var prev1 = -1
	var prev0 = -1

	var prevSum = -1

	for scanner.Scan() {
		line := scanner.Text()
		cur, _ := strconv.Atoi(line)

		prev0 = cur
		curSum := prev0 + prev1 + prev2

		if prev0 != -1 && prev1 != -1 && prev2 != -1 {
			if curSum > prevSum {
				c++
			}
		}
		prev2 = prev1
		prev1 = prev0
		prevSum = curSum
	}
	fmt.Println(c - 1)
}
