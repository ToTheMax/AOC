package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func checkBingo(board [][]string, drawnNumbers []string, time int) bool {
	for _, rowcol := range board {
		complete := true
		for _, value := range rowcol {
			valueIsDrawn := false
			for _, drawnNumber := range drawnNumbers[0:time] {
				if drawnNumber == value {
					valueIsDrawn = true
				}
			}
			if !valueIsDrawn {
				complete = false
				break
			}
		}
		if complete {
			return true
		}
	}
	return false
}

func main() {

	lines, _ := readLines("input.txt")

	boardSize := 5

	drawnNumbers := strings.Split(lines[0], ",")

	boardCount := len(lines) / (boardSize + 1)
	boards := make([][][]string, boardCount)

	for board := 0; board < boardCount; board++ {
		// fmt.Println("BOARD", board)
		boards[board] = make([][]string, boardSize*2)

		index := board*(boardSize+1) + 2
		lines := lines[index : index+boardSize]

		// Add rows to board
		for i, line := range lines {
			row := strings.Fields(line)
			boards[board][i] = row
			// fmt.Println(row)
		}

		// Add columns to board
		for j := 0; j < boardSize; j++ {
			column := make([]string, boardSize)
			for i, row := range boards[board][0:boardSize] {
				column[i] = row[j]
			}
			// fmt.Println(column)
			boards[board][boardSize+j] = column
		}
	}
	foundBingo := false
	for time := 0; time < len(drawnNumbers); time++ {
		for _, board := range boards {
			if checkBingo(board, drawnNumbers, time) {
				foundBingo = true

				sum := 0
				for _, row := range board[0:boardSize] {
					for _, val := range row {
						isDrawn := false
						for _, drawnNumber := range drawnNumbers[0:time] {
							if val == drawnNumber {
								isDrawn = true
							}
						}
						if !isDrawn {
							valInt, _ := strconv.Atoi(val)
							sum = sum + valInt
						}
					}
				}

				lastDrawnInt, _ := strconv.Atoi(drawnNumbers[time-1])
				fmt.Println(lastDrawnInt * sum)
				break
			}
		}
		if foundBingo {
			break
		}
	}
}
