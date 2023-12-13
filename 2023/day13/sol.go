package main

import (
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func intersec(a, b []int) []int {
	m := make(map[int]bool)
	var intersect []int
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if _, exists := m[item]; exists {
			delete(m, item)
			intersect = append(intersect, item)
		}
	}
	return intersect
}

func getLineMirrors(str string, ignore int) []int {
	mirrors := []int{}
	for i := 1; i < len(str); i++ {
		if i == ignore {
			continue
		}
		l1 := str[i:]
		l2 := reverseString(str[:i])
		length := min(len(l1), len(l2))
		if l1[:length] == l2[:length] {
			mirrors = append(mirrors, i)
		}

	}
	return mirrors
}

func getMirror(strs []string, ignore int) int {
	mirrors := getLineMirrors(strs[0], ignore)
	for _, str := range strs {
		mirrors = intersec(mirrors, getLineMirrors(str, ignore))
	}
	if len(mirrors) == 0 {
		return 0
	}
	return mirrors[0]
}

func main() {
	input, _ := os.ReadFile("in.txt")
	sumP1 := 0
	sumP2 := 0

	patterns := strings.Split(string(input), "\n\n")

	for _, pattern := range patterns {
		rows := strings.Split(string(pattern), "\n")
		cols := make([]string, len(rows[0]))
		for _, row := range rows {
			for x, char := range row {
				cols[x] += string(char)
			}
		}

		// Part 1
		rp1 := getMirror(rows, -1)
		cp1 := getMirror(cols, -1)

		sumP1 += rp1
		sumP1 += 100 * cp1

		// Part 2
		foundReflection := false
		for i := 0; i < len(rows); i++ {
			if !foundReflection {
				for j := 0; j < len(cols); j++ {
					// Flip smudge
					orig := string(rows[i][j])
					flip := string(".#"[strings.Index("#.", orig)])
					cols[j] = cols[j][:i] + flip + cols[j][i+1:]
					rows[i] = rows[i][:j] + flip + rows[i][j+1:]

					if cp2 := getMirror(cols, cp1); cp2 > 0 {
						foundReflection = true
						sumP2 += 100 * cp2
						break
					} else if rp2 := getMirror(rows, rp1); rp2 > 0 {
						foundReflection = true
						sumP2 += rp2
						break
					}
					// Flip back
					cols[j] = cols[j][:i] + orig + cols[j][i+1:]
					rows[i] = rows[i][:j] + orig + rows[i][j+1:]
				}
			}
		}
	}

	fmt.Println("Sol 1:", sumP1)
	fmt.Println("Sol 2:", sumP2)
}