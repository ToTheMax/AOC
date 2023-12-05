package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	source      int
	destination int
	length      int
}

type Map struct {
	name   string
	ranges []Range
}

func stringToInts(s string) []int {
	splitted := strings.Split(s, " ")
	ints := make([]int, len(splitted))
	for i, s := range splitted {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func stringToRanges(s string) [][]int {
	splitted := strings.Split(s, " ")
	ranges := make([][]int, len(splitted)/2)
	for i := 0; i < len(splitted); i += 2 {
		p1, _ := strconv.Atoi(splitted[i])
		p2, _ := strconv.Atoi(splitted[i+1])
		ranges[i/2] = []int{p1, p1 + p2}
	}
	return ranges
}

func seedToLocation(seed int, maps []Map) int {
	location := seed
	for _, m := range maps {
		for _, r := range m.ranges {
			if location >= r.source && location <= r.source+r.length {
				location = r.destination + (location - r.source)
				break
			}
		}
	}
	return location
}

func locationToSeed(location int, maps []Map) int {
	seed := location
	for i := range maps {
		for _, r := range maps[len(maps)-i-1].ranges {
			if seed >= r.destination && seed <= r.destination+r.length {
				seed = r.source + (seed - r.destination)
				break
			}
		}
	}
	return seed
}

func main() {

	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	seeds := stringToInts(lines[0][7:])
	seedRanges := stringToRanges(lines[0][7:])

	maps := []Map{}
	for i := 2; i < len(lines); i++ {
		name := strings.Split(lines[i], "map")[0]
		newMap := Map{name, []Range{}}
		for j := i + 1; j < len(lines); j++ {
			i++
			if lines[j] == "" || i == len(lines)-1 {
				maps = append(maps, newMap)
				break
			} else {
				r := stringToInts(lines[j])
				newMap.ranges = append(newMap.ranges, Range{r[1], r[0], r[2]})
			}
		}
	}

	min := seedToLocation(seeds[0], maps)
	for _, seed := range seeds {
		location := seedToLocation(seed, maps)
		if location < min {
			min = location
		}
	}
	fmt.Println("Sol 1:", min)

	foundSolution := false
	location := 0
	for !foundSolution {
		location++
		initialSeed := locationToSeed(location, maps)
		for _, seedRange := range seedRanges {
			if initialSeed >= seedRange[0] && initialSeed <= seedRange[1] {
				foundSolution = true
				break
			}
		}
	}
	fmt.Println("Sol 2:", location)
}
