// day seven solution
package main

import (
	"fmt"
	"math"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_7.txt")

	answerOne := partOne(input)
	answerTwo := partTwo(input)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func partOne(input []string) int {
	// I think that there are much better ways to do this e.g utilising the mean

	crabs := utils.CSVToInts(input[0])

	minCrab, maxCrab := utils.MinMax(crabs)

	optimum := math.MaxInt64

	for i := minCrab; i <= maxCrab; i++ {
		totalDist := 0
		for _, crab := range crabs {
			totalDist += int(math.Abs(float64(i - crab)))
		}
		if totalDist < optimum {
			optimum = totalDist
		}
	}

	return optimum
}

func partTwo(input []string) int {

	crabs := utils.CSVToInts(input[0])

	minCrab, maxCrab := utils.MinMax(crabs)

	optimum := math.MaxInt64

	for i := minCrab; i <= maxCrab; i++ {
		totalDist := 0
		for _, crab := range crabs {
			dist := math.Abs(float64(i - crab))
			totalDist += int(0.5 * dist * (dist + 1))
		}
		if totalDist < optimum {
			optimum = totalDist
		}
	}

	return optimum
}
