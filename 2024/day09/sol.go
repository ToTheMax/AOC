package main

import (
	"fmt"
	"os"
	"strconv"
)

func solve(disk_map_string string, part int) int {

	disk_map := make([]int, len(disk_map_string))
	for i := 0; i < len(disk_map_string); i++ {
		disk_map[i], _ = strconv.Atoi(string(disk_map_string[i]))
	}

	checksum := 0
	placed_files := make(map[int]bool)
	latest_file_index := len(disk_map) - 1
	filesystem_index := 0
	for disk_map_index := 0; disk_map_index < len(disk_map); disk_map_index++ {
		size := disk_map[disk_map_index]
		if disk_map_index%2 == 0 && !placed_files[disk_map_index] {
			for j := 0; j < size; j++ {
				checksum += filesystem_index * (disk_map_index / 2)
				filesystem_index++
			}
			disk_map[disk_map_index] = 0
		} else {
			// Part 1: Place latest file in the free space
			if part == 1 {
				for j := 0; j < size; j++ {
					if disk_map[latest_file_index] == 0 {
						latest_file_index -= 2
						if disk_map[latest_file_index] == 0 {
							break
						}
					}
					checksum += filesystem_index * (latest_file_index / 2)
					filesystem_index++
					disk_map[latest_file_index]--
				}
			}
			// Part 2: Place latest file that fits and not already placed in the free space
			if part == 2 {
				finished := false
				for !finished {
					finished = true
					for file_index := len(disk_map) - 1; file_index >= disk_map_index; file_index -= 2 {
						if disk_map[file_index] <= size && !placed_files[file_index] {
							placed_files[file_index] = true
							for i := 0; i < disk_map[file_index]; i++ {
								checksum += filesystem_index * (file_index / 2)
								filesystem_index++
								size--
							}
							finished = false
							break
						}
					}
				}
				filesystem_index += size
			}
		}
	}
	return checksum
}

func main() {
	input, _ := os.ReadFile("in.txt")
	fmt.Println("Sol1:", solve(string(input), 1))
	fmt.Println("Sol2:", solve(string(input), 2))
}
