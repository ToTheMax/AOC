package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	pos := 0
	depth := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_splitted := strings.Split(line, " ")

		command := line_splitted[0]
		val, _ := strconv.Atoi(line_splitted[1])

		switch command {
		case "forward":
			pos = pos + val
		case "down":
			depth = depth + val
		case "up":
			depth = depth - val
		}
	}

	fmt.Println(pos * depth)
}
