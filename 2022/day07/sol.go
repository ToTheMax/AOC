package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	directory bool
	size      int
	parent    *File
	children  map[string]*File
}

func addDir(curDir *File, line string) {
	dir := line[4:]
	curDir.children[dir] = &File{
		directory: true,
		size:      0,
		parent:    curDir,
		children:  make(map[string]*File),
	}
}

func addFile(curDir *File, line string) {
	splitted_file := strings.Split(line, " ")
	size, _ := strconv.Atoi(splitted_file[0])
	name := splitted_file[1]
	curDir.children[name] = &File{
		directory: false,
		size:      size,
		parent:    curDir,
		children:  nil,
	}
	parent := curDir
	for parent != nil {
		parent.size += size
		parent = parent.parent
	}
}

func findDirs(root *File, minSize int, maxSize int) []*File {
	var files []*File
	if root.directory {
		for _, child := range root.children {
			if child.directory {
				if child.size >= minSize && child.size <= maxSize {
					files = append(files, child)
				}
				files = append(files, findDirs(child, minSize, maxSize)...)
			}
		}
	}
	return files
}

func main() {
	input, _ := os.ReadFile("in.txt")
	commands := strings.Split(string(input), "$ ")

	root := &File{
		directory: true,
		size:      0,
		parent:    nil,
		children:  make(map[string]*File),
	}
	curDir := root

	// Parse input into a tree
	for _, command := range commands[2:] {
		lines := strings.Split(string(strings.Trim(command, "\n")), "\n")
		if strings.HasPrefix(lines[0], "cd ") {
			// Change directory
			newDir := lines[0][3:]
			if newDir == ".." {
				curDir = curDir.parent
			} else {
				curDir = curDir.children[lines[0][3:]]
			}
		} else if strings.HasPrefix(lines[0], "ls") {
			// List directory
			for _, line := range lines[1:] {
				if strings.HasPrefix(line, "dir ") {
					addDir(curDir, line)
				} else {
					addFile(curDir, line)
				}
			}
		}
	}

	// Part 1
	sum := 0
	for _, dir := range findDirs(root, 0, 100000) {
		sum += dir.size
	}
	fmt.Println("Sol1:", sum)

	// Part 2
	minimalSize := 30000000 - (70000000 - root.size)
	smallestDirSize := math.MaxInt
	for _, dir := range findDirs(root, minimalSize, math.MaxInt) {
		if dir.size < smallestDirSize {
			smallestDirSize = dir.size
		}
	}
	fmt.Println("Sol2:", smallestDirSize)
}
