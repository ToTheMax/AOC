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

func bitStringToIntArray(line string) []int {
	bitArray := make([]int, len(line))
	for i, char := range line {
		if char == '1' {
			bitArray[i] = 1
		} else {
			bitArray[i] = 0
		}
	}
	return bitArray
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func find(lines []string, findMostCommon bool) string {
	lineWidth := 12
	isMostCommon := make([]bool, len(lines))
	for i := 0; i < len(lines); i++ {
		isMostCommon[i] = true
	}

	mostCommonResult := ""
	mostCommonCount := len(lines)

	for bitNumber := 0; bitNumber < lineWidth; bitNumber++ {

		// Find one bits
		oneBitCounts := 0
		for i, line := range lines {
			if isMostCommon[i] {
				if rune(line[bitNumber]) == '1' {
					oneBitCounts++
				}
			}
		}

		mostCommon := '1'
		if findMostCommon {
			if oneBitCounts*2 < mostCommonCount {
				mostCommon = '0'
			}
		} else {
			if oneBitCounts*2 >= mostCommonCount {
				mostCommon = '0'
			}
		}

		// Set boolean for most common
		mostCommonCount = 0
		for i, line := range lines {
			if isMostCommon[i] && rune(line[bitNumber]) == mostCommon {
				isMostCommon[i] = true
				mostCommonResult = line
				mostCommonCount++
			} else {
				isMostCommon[i] = false
			}
		}

		// Found result
		if mostCommonCount == 1 {
			break
		}
	}
	return mostCommonResult
}

func main() {

	lines, err := readLines("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	mostCommonResult := find(lines, true)
	leastCommonResult := find(lines, false)

	finalResult := bitArrayToInt(bitStringToIntArray(mostCommonResult)) * bitArrayToInt(bitStringToIntArray(leastCommonResult))

	fmt.Println(finalResult)
}
