package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	input := readData("../data/data_day_1.txt")

	part_one := PartOne(input)
	part_two := PartTwo(input)

	fmt.Println(part_one)
	fmt.Println(part_two)

}

func PartOne(input []int) int {
	output := 0
	for i, depth := range input {
		if i!=0 && depth > input[i-1] {
			output++
		}
	}
	return output
}

func PartTwo(input []int) int {
	prev_depth := 0
	output := -1
	for i := 3; i <= len(input); i++ {
		window := input[i-3 : i]
		depth := sum(window)
		if depth > prev_depth {
			output++
		}
		prev_depth = depth
	}
	return output
}

func sum(input []int) int {
	output := 0
	for _, num := range input {
		output += num
	}
	return output
}

func readData(path string) []int {

	file_bytes, _ := ioutil.ReadFile(toAbsPath(path))
	input := string(file_bytes)

	input_strings := strings.Split(input, "\n")

	input_ints := make([]int, len(input_strings))

	for i, str := range input_strings {
		input_ints[i], _ = strconv.Atoi(str)
	}

	return input_ints
}

func toAbsPath(rel_path string) string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), rel_path)
}
