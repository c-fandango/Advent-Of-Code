// day six solution
package main

import (
	"fmt"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_6.txt")

	answerOne := partOne(input, 80)
	answerTwo := partOne(input, 256)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func initMap(fishes []int, count int) map[int]int {
	ages := map[int]int{-1: 0}
	for i := 0; i <= count; i++ {
		_, exists := ages[i]
		if !exists {
			ages[i] = 0
		}
	}

	for _, fish := range fishes {
		ages[fish]++
	}

	return ages
}

func partOne(input []string, days int) int {
	const childhood, cycleLength = 8, 6

	fishes := utils.CSVToInts(input[0])
	ages := initMap(fishes, childhood)

	for days > 0 {
		for i := 0; i <= childhood; i++ {
			ages[i-1] = ages[i]
		}
		ages[childhood] = ages[-1]
		ages[cycleLength] += ages[-1]
		ages[-1] = 0

		days--
	}

	endFishes := 0
	for _, num := range ages {
		endFishes += num
	}

	return endFishes
}
