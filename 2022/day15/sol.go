package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}
type Beacon struct {
	location Point
	radius   int
	sensor   Point
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func excludeRow(x int, amount int, y int, row int, excluded map[Point]bool) {
	if y == row {
		for i := 0; i < amount; i++ {
			if _, ok := excluded[Point{x + i, y}]; !ok {
				excluded[Point{x + i, y}] = true
			}
		}
	}
}

func exclude(centre Point, size int, row int, excluded map[Point]bool) {
	top := centre.Y - size
	bot := centre.Y + size

	width := 1
	for i := 0; i < size; i++ {
		excludeRow(centre.X-i, i*2+1, top+i, row, excluded)
		excludeRow(centre.X-i, i*2+1, bot-i, row, excluded)
		width += 2
	}
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	var beacons []Beacon
	for _, line := range lines {
		var beaconLoc Point
		var sensorLoc Point
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensorLoc.X, &sensorLoc.Y, &beaconLoc.X, &beaconLoc.Y)
		beacons = append(beacons, Beacon{
			location: beaconLoc,
			sensor:   sensorLoc,
			radius:   distance(beaconLoc, sensorLoc),
		})
	}

	row := 2000000
	excluded := make(map[Point]bool)
	for _, beacon := range beacons {
		exclude(beacon.sensor, beacon.radius, row, excluded)
	}
	for _, beacon := range beacons {
		delete(excluded, beacon.location)
	}
	fmt.Println("Sol1:", len(excluded))

	limit := 4000000
	possiblePoints := make(map[Point]bool)
	for _, beacon := range beacons {
		radiusBorder := beacon.radius + 1
		for i := 0; i < radiusBorder; i++ {
			if beacon.sensor.X+i > 0 && beacon.sensor.X+i < limit {
				if beacon.sensor.Y-radiusBorder+1+i > 0 && beacon.sensor.Y-radiusBorder+1+i < limit {
					possiblePoints[Point{beacon.sensor.X + i, beacon.sensor.Y - radiusBorder + 1 + i}] = true
				}
				if beacon.sensor.Y+radiusBorder-1-i > 0 && beacon.sensor.Y+radiusBorder-1-i < limit {
					possiblePoints[Point{beacon.sensor.X + i, beacon.sensor.Y + radiusBorder - i}] = true
				}
			}
			if beacon.sensor.X-i > 0 && beacon.sensor.X-i < limit {
				if beacon.sensor.Y-radiusBorder+1+i > 0 && beacon.sensor.Y-radiusBorder+1+i < limit {
					possiblePoints[Point{beacon.sensor.X - i, beacon.sensor.Y - radiusBorder + 1 + i}] = true
				}
				if beacon.sensor.Y+radiusBorder-1-i > 0 && beacon.sensor.Y+radiusBorder-1-i < limit {
					possiblePoints[Point{beacon.sensor.X - i, beacon.sensor.Y + radiusBorder - 1 - i}] = true
				}
			}
		}
	}

	score := 0
	for possiblePoint := range possiblePoints {
		possible := true
		for _, beacon := range beacons {
			if distance(beacon.sensor, possiblePoint) <= beacon.radius {
				possible = false
				break
			}
		}
		if possible {
			score = possiblePoint.X*4000000 + possiblePoint.Y
			break
		}
	}
	fmt.Println("Sol2:", score)
}
