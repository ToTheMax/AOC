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

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	width := 101
	height := 103

	robots := make([]Robot, len(lines))
	for i, line := range lines {
		var p Pos
		var v Velocity
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p.x, &p.y, &v.x, &v.y)

		robots[i] = Robot{pos: p, velocity: v}
	}

	for second := 0; second < 100; second++ {
		for i, robot := range robots {
			robots[i].pos.x = mod(robot.pos.x+robot.velocity.x, width)
			robots[i].pos.y = mod(robot.pos.y+robot.velocity.y, height)
		}
	}

	// Count how many robots are in each quadrant
	horCenter := width / 2
	verCenter := height / 2
	quadrant := make([]int, 4)
	for _, robot := range robots {
		if robot.pos.x < horCenter && robot.pos.y < verCenter {
			quadrant[0]++
		} else if robot.pos.x < horCenter && robot.pos.y > verCenter {
			quadrant[1]++
		} else if robot.pos.x > horCenter && robot.pos.y < verCenter {
			quadrant[2]++
		} else if robot.pos.x > horCenter && robot.pos.y > verCenter {
			quadrant[3]++
		}
	}
	fmt.Println("Sol 1:", quadrant[0]*quadrant[1]*quadrant[2]*quadrant[3])
}
