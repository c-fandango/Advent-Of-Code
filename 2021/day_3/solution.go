// day three
package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_3.txt")

	answerOne := partOne(input)
	answerTwo := partTwo(input)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func modeBits(input []string) (string, string) {
	length := len(input[0])
	common := ""
	leastCommon := ""
	counts := make([]int, length)

	for i := 0; i < length; i++ {
		counts[i] = 0
	}

	for _, code := range input {
		for i := 0; i < length; i++ {
			num, _ := strconv.Atoi(string(code[i]))
			counts[i] += num
		}
	}

	for _, count := range counts {
		if 2*count >= len(input) {
			common += "1"
			leastCommon += "0"
		} else {
			common += "0"
			leastCommon += "1"
		}
	}
	return common, leastCommon
}

func partOne(input []string) int64 {

	length := len(input[0])
	gamma, _ := modeBits(input)

	gammaVal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonVal := int64(math.Pow(2, float64(length))) - gammaVal - 1

	return gammaVal * epsilonVal
}

func partTwo(input []string) int64 {
	//not very nice - maybe make nicer later
	var newRemaining []string

	length := len(input[0])
	remaining := input

	//oxygen
	for i := 0; i < length; i++ {
		newRemaining = []string{}
		if len(remaining) == 1 {
			break
		}
		common, _ := modeBits(remaining)
		for _, code := range remaining {
			if code[i] == common[i] {
				newRemaining = append(newRemaining, code)
			}
		}
		remaining = newRemaining
	}

	oxygen := remaining[0]

	//carbon
	remaining = input
	for i := 0; i < length; i++ {
		newRemaining = []string{}
		if len(remaining) == 1 {
			break
		}
		_, leastCommon := modeBits(remaining)
		for _, code := range remaining {
			if code[i] == leastCommon[i] {
				newRemaining = append(newRemaining, code)
			}
		}
		remaining = newRemaining
	}

	carbon := remaining[0]
	oxygenVal, _ := strconv.ParseInt(oxygen, 2, 64)
	carbonVal, _ := strconv.ParseInt(carbon, 2, 64)

	return oxygenVal * carbonVal
}
