package main

import (
	"fmt"
	"os"
	"strings"
)

type Robot struct {
	pos      Pos
	velocity Velocity
}

type Pos struct {
	x, y int
}

type Velocity struct {
	x, y int
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func parseRobots(lines []string) []Robot {
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		var p Pos
		var v Velocity
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p.x, &p.y, &v.x, &v.y)
		robots[i] = Robot{pos: p, velocity: v}
	}
	return robots
}

func solve(robots []Robot, width, height int, part int) int {
	maxLocations := len(robots)

	for second := 0; second < 100000; second++ {
		for i, robot := range robots {
			robots[i].pos.x = mod(robot.pos.x+robot.velocity.x, width)
			robots[i].pos.y = mod(robot.pos.y+robot.velocity.y, height)
		}

		if part == 1 && second == 100 {
			quads := make([]int, 4)
			for _, robot := range robots {
				if robot.pos.x < width/2 && robot.pos.y < height/2 {
					quads[0]++
				} else if robot.pos.x < width/2 && robot.pos.y > height/2 {
					quads[1]++
				} else if robot.pos.x > width/2 && robot.pos.y < height/2 {
					quads[2]++
				} else if robot.pos.x > width/2 && robot.pos.y > height/2 {
					quads[3]++
				}
			}
			return quads[0] * quads[1] * quads[2] * quads[3]
		}
		if part == 2 {
			locations := make(map[Pos]bool)
			for _, robot := range robots {
				locations[robot.pos] = true
			}
			if len(locations) == maxLocations {
				printMap(locations, width, height)
				return second + 1
			}
		}
	}
	return -1
}

// Print the robots in the map, robot an R, the rest a dot
func printMap(locations map[Pos]bool, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if locations[Pos{x, y}] {
				fmt.Print("R")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	fmt.Println("Sol 1:", solve(parseRobots(lines), 101, 103, 1))
	fmt.Println("Sol 2:", solve(parseRobots(lines), 101, 103, 2))
}
