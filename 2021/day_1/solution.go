// day one solution
package main

import (
	"fmt"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataInt("../data/data_day_1.txt")

	answerOne := partOne(input)
	answerTwo := partTwo(input)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func partOne(input []int) int {
	output := 0
	for i, depth := range input {
		if i != 0 && depth > input[i-1] {
			output++
		}
	}
	return output
}

func partTwo(input []int) int {
	prevDepth := 0
	output := -1
	for i := 3; i <= len(input); i++ {
		window := input[i-3 : i]
		depth := sum(window)
		if depth > prevDepth {
			output++
		}
		prevDepth = depth
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
