package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	focalLength int
	label       string
	index       int
}

func hash(step string) int {
	curValue := 0
	for _, char := range step {
		curValue = (curValue + int(char)) * 17 % 256
	}
	return curValue
}

func main() {
	input, _ := os.ReadFile("in.txt")
	steps := strings.Split(string(input), ",")

	// Part 1
	sumP1 := 0
	for _, step := range steps {
		sumP1 += hash(step)
	}
	fmt.Println("Sol 1:", sumP1)

	// Part 2
	sumP2 := 0
	buckets := make([]map[string]Lens, 256)
	for i := range buckets {
		buckets[i] = make(map[string]Lens)
	}
	for _, step := range steps {
		// Lens removal
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			bucket := buckets[hash(label)]
			if removedLens, ok := bucket[label]; ok {
				delete(bucket, label)
				for _, l := range bucket {
					if l.index > removedLens.index {
						bucket[l.label] = Lens{l.focalLength, l.label, l.index - 1}
					}
				}
			}
		}

		// Lens addition
		if step[len(step)-2] == '=' {
			split := strings.Split(step, "=")
			label := split[0]
			value, _ := strconv.Atoi(split[1])

			bucket := buckets[hash(label)]
			if lens, ok := bucket[label]; ok {
				bucket[label] = Lens{value, label, lens.index}
			} else {
				bucket[label] = Lens{value, label, len(bucket)}
			}
		}
	}
	for b, bucket := range buckets {
		for _, lens := range bucket {
			sumP2 += (b + 1) * (lens.index + 1) * lens.focalLength
		}
	}
	fmt.Println("Sol 2:", sumP2)
}
