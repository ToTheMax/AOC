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
	var prev = 0

	for scanner.Scan() {
		line := scanner.Text()
		cur, _ := strconv.Atoi(line)
		if cur > prev {
			c++
		}
		prev = cur
	}
	fmt.Println(c - 1)
}
