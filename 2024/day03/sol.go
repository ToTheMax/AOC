package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	p1_score, p2_score := 0, 0
	enabled := true
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if strings.HasPrefix(line[i:], "do()") {
				enabled = true
			} else if strings.HasPrefix(line[i:], "don't()") {
				enabled = false
			} else {
				var l, r int
				_, err := fmt.Sscanf(line[i:], "mul(%d,%d)", &l, &r)
				if err == nil {
					p1_score += l * r
					if enabled {
						p2_score += l * r
					}
				}
			}
		}
	}

	fmt.Println("Sol 1:", p1_score)
	fmt.Println("Sol 2:", p2_score)
}
