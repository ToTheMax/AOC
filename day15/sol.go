package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	SENSOR  = 1
	BEACON  = 2
	EXCLUDE = 3
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func excludeRow(x int, amount int, y int, grid map[Point]int) {
	if y == 2000000 {
		for i := 0; i < amount; i++ {
			if _, ok := grid[Point{x + i, y}]; !ok {
				grid[Point{x + i, y}] = EXCLUDE
			}
		}
	}
}

func exclude(centre Point, size int, grid map[Point]int) {
	top := centre.Y - size
	bot := centre.Y + size

	width := 1
	for i := 0; i < size; i++ {
		excludeRow(centre.X-i, i*2+1, top+i, grid)
		excludeRow(centre.X-i, i*2+1, bot-i, grid)
		width += 2
	}
}

type Point struct {
	X int
	Y int
}

func main() {
	input, _ := os.ReadFile("in.txt")

	lines := strings.Split(string(input), "\n")

	sensors := make([]Point, len(lines))
	beacons := make([]Point, len(lines))
	distances := make([]int, len(lines))
	grid := make(map[Point]int)
	for i, line := range lines {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensors[i].X, &sensors[i].Y, &beacons[i].X, &beacons[i].Y)
		distances[i] = abs(distance(sensors[i], beacons[i]))
		grid[sensors[i]] = SENSOR
		grid[beacons[i]] = BEACON
	}

	for i, sensor := range sensors {
		exclude(sensor, distances[i], grid)
	}

	count := 0
	for pos, t := range grid {
		if pos.Y == 2000000 && t == EXCLUDE {
			count++
		}
	}
	fmt.Println("Sol1", count)
}
