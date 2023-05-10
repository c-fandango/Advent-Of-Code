package main

import (
	"fmt"
	"github.com/c-fandango/advent_of_code/2021/utils"
	"math"
	"strconv"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_3.txt")

	part_one := PartOne(input)
	part_two := PartTwo(input)

	fmt.Println(part_one)
	fmt.Println(part_two)

}

func modeBits(input []string) (string, string) {
	length := len(input[0])
	common := ""
	least_common := ""
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
			least_common += "0"
		} else {
			common += "0"
			least_common += "1"
		}
	}
	return common, least_common
}

func PartOne(input []string) int64 {

	length := len(input[0])
	gamma, _ := modeBits(input)

	gamma_val, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_val := int64(math.Pow(2, float64(length))) - gamma_val - 1

	return gamma_val * epsilon_val
}

func PartTwo(input []string) int64 {
	//not very nice - maybe make nicer later
	var new_remaining []string

	length := len(input[0])

	remaining := input

	//oxygen
	for i := 0; i < length; i++ {
		new_remaining = []string{}
		if len(remaining) == 1 {
			break
		}
		common, _ := modeBits(remaining)
		for _, code := range remaining {
			if code[i] == common[i] {
				new_remaining = append(new_remaining, code)
			}
		}
		remaining = new_remaining
	}

	oxygen := remaining[0]

	//carbon
	remaining = input
	for i := 0; i < length; i++ {
		new_remaining = []string{}
		if len(remaining) == 1 {
			break
		}
		_, least_common := modeBits(remaining)
		for _, code := range remaining {
			if code[i] == least_common[i] {
				new_remaining = append(new_remaining, code)
			}
		}
		remaining = new_remaining
	}

	carbon := remaining[0]
	oxygen_val, _ := strconv.ParseInt(oxygen, 2, 64)
	carbon_val, _ := strconv.ParseInt(carbon, 2, 64)

	return oxygen_val * carbon_val
}
