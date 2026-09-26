package main

import "core:fmt"
import "core:os"
import "core:strings"

main :: proc() {
	read_file_by_lines_in_whole("input.txt")
}

read_file_by_lines_in_whole :: proc(filepath: string) {
	data, err := os.read_entire_file(filepath, context.allocator)
	if err != nil {
		// could not read file
		return
	}
	defer delete(data, context.allocator)

	it := string(data)
	for line in strings.split_lines_iterator(&it) {
		fmt.println(line)
	}
}

