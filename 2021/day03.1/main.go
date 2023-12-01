package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func bitArrayToInt(bitArray []int) int {
	sum := 0
	for i := range bitArray {
		index := len(bitArray) - 1 - i
		if bitArray[index] == 1 {
			pow := int(math.Pow(2, float64(i)))
			sum = sum + pow
		}
	}
	return sum
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lineWidth := 12

	lineCount := 0
	oneBitCounts := make([]int, lineWidth)
	for i := range oneBitCounts {
		oneBitCounts[i] = 0
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		for i, char := range line {
			if char == '1' {
				oneBitCounts[i]++
			}
		}
	}

	gamma := make([]int, lineWidth)
	epsilon := make([]int, lineWidth)
	for i := range gamma {
		if oneBitCounts[i] > lineCount/2 {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	gamma_decimal := bitArrayToInt(gamma)
	epsilon_decimal := bitArrayToInt(epsilon)

	fmt.Println((gamma_decimal * epsilon_decimal))
}
