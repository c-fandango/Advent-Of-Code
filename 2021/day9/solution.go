// day two solution
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_9.txt")

	answerOne := partOne(input)
	answerTwo := partTwo(input)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func extractData(input []string) map[[2]int]int {
	output := make(map[[2]int]int)
	for i, row := range input {
		rowSplt := strings.Split(row, "")
		for j, cell := range rowSplt {
			output[[2]int{i, j}], _ = strconv.Atoi(cell)
			// build a wall around cave
			output[[2]int{-1, j}] = 9
			output[[2]int{len(input), j}] = 9
		}
		output[[2]int{i, -1}] = 9
		output[[2]int{i, len(rowSplt)}] = 9
	}
	return output
}

func findLowPoints(vents map[[2]int]int) map[[2]int]int {
	lowPoints := make(map[[2]int]int)

	for coord, height := range vents {
		x, y := coord[0], coord[1]
		left := [2]int{x - 1, y}
		right := [2]int{x + 1, y}
		down := [2]int{x, y - 1}
		up := [2]int{x, y + 1}
		if vents[left] > height && vents[right] > height && vents[down] > height && vents[up] > height {
			lowPoints[coord] = height
		}
	}

	return lowPoints
}

func partOne(input []string) int {
	vents := extractData(input)

	lowPoints := findLowPoints(vents)
	lowPointSum := 0
	for _, value := range lowPoints {
		lowPointSum += value + 1
	}

	return lowPointSum

}

func expandBasin(point [2]int, basin map[[2]int]int, vents map[[2]int]int) map[[2]int]int {
	height := vents[point]
	if height == 9 {
		return basin
	}
	basin[point] = height

	x, y := point[0], point[1]
	left := [2]int{x - 1, y}
	if vents[left] > height {
		basin = expandBasin(left, basin, vents)
	}
	right := [2]int{x + 1, y}
	if vents[right] > height {
		basin = expandBasin(right, basin, vents)
	}
	down := [2]int{x, y - 1}
	if vents[down] > height {
		basin = expandBasin(down, basin, vents)
	}
	up := [2]int{x, y + 1}
	if vents[up] > height {
		basin = expandBasin(up, basin, vents)
	}
	delete(vents, point)

	return basin
}

func partTwo(input []string) int {
	vents := extractData(input)
	counts := []int{0, 0, 0}

	lowPoints := findLowPoints(vents)
	for point, _ := range lowPoints {
		basin := make(map[[2]int]int)
		basin = expandBasin(point, basin, vents)
		if len(basin) > counts[0] {
			counts[0] = len(basin)
			sort.Ints(counts)
		}
	}

	return counts[0] * counts[1] * counts[2]
}
