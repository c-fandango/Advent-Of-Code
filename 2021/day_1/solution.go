package main

import (
	"fmt"
	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataInt("../data/data_day_1.txt")

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

