package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("in.txt")
	disk_map_string := string(input)
	// disk_map_string = "2333133121414131402"

	checksum := 0

	disk_map := make([]int, len(disk_map_string))
	for i := 0; i < len(disk_map_string); i++ {
		disk_map[i], _ = strconv.Atoi(string(disk_map_string[i]))
	}

	latest_file_index := len(disk_map) - 1
	filesystem_index := 0
	for disk_map_index := 0; disk_map_index < len(disk_map); disk_map_index++ {
		if disk_map_index%2 == 0 {
			for j := 0; j < disk_map[disk_map_index]; j++ {
				checksum += filesystem_index * (disk_map_index / 2)
				// fmt.Println(filesystem_index, "*", disk_map_index, "=", filesystem_index*disk_map_index)
				// fmt.Print(disk_map_index / 2)
				filesystem_index++
			}
			disk_map[disk_map_index] = 0
		} else {
			for j := 0; j < disk_map[disk_map_index]; j++ {
				// Move one file to the left
				if disk_map[latest_file_index] == 0 {
					latest_file_index -= 2
				}
				// If this file is already finished, break
				if disk_map[latest_file_index] == 0 {
					break
				}
				checksum += filesystem_index * (latest_file_index / 2)
				// fmt.Println(filesystem_index, "*", latest_file_index/2, "=", filesystem_index*(latest_file_index/2))
				// fmt.Print(latest_file_index / 2)
				filesystem_index++
				disk_map[latest_file_index]--
			}
		}
	}
	fmt.Println("Sol1:", checksum)
}
