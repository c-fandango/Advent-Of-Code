package main

import (
	"fmt"
	"github.com/c-fandango/advent_of_code/2021/utils"
	"strconv"
	"strings"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_2.txt")

	part_one := PartOne(input)
	part_two := PartTwo(input)

	fmt.Println(part_one)
	fmt.Println(part_two)

}

func PartOne(input []string) int {
	xpos := 0
	ypos := 0

	for _, move := range input {
		if len(move) == 0 {
			continue
		}
		command := strings.Split(move, " ")

		dir := command[0]
		dist, _ := strconv.Atoi(command[1])
		if dir == "up" {
			ypos -= dist
		} else if dir == "down" {
			ypos += dist
		} else if dir == "forward" {
			xpos += dist
		}
	}
	return xpos * ypos
}

func PartTwo(input []string) int {
	xpos := 0
	ypos := 0
	aim := 0

	for _, move := range input {
		if len(move) == 0 {
			continue
		}
		command := strings.Split(move, " ")

		dir := command[0]
		dist, _ := strconv.Atoi(command[1])
		if dir == "up" {
			aim -= dist
		} else if dir == "down" {
			aim += dist
		} else if dir == "forward" {
			xpos += dist
			ypos += dist * aim
		}
	}
	return xpos * ypos
}
