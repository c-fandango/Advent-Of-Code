// day two solution
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_5.txt")

	answerOne := partOne(input)
	answerTwo := partTwo(input)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func extractData(input []string) [][][2]int {
	output := make([][][2]int, len(input))

	for i, line := range input {
		lineSplt := strings.Split(line, " -> ")
		moves := make([][2]int, len(lineSplt))

		for j, move := range lineSplt {
			coords := strings.Split(move, ",")
			var coordsNum [2]int
			for k, coord := range coords {
				coordsNum[k], _ = strconv.Atoi(coord)
			}
			moves[j] = coordsNum
		}
		output[i] = moves
	}
	return output
}

func updatePoints(points map[[2]int]int, move [][2]int) map[[2]int]int {

	start, end := move[0], move[1]
	xDist := end[0] - start[0]
	yDist := end[1] - start[1]
	xDir := utils.Sign(xDist)
	yDir := utils.Sign(yDist)

	trueDist := xDist * xDir
	if trueDist == 0 {
		trueDist = yDist * yDir
	}

	for i := 0; i <= trueDist; i++ {
		xNew := start[0] + i*xDir
		yNew := start[1] + i*yDir
		newPoint := [2]int{xNew, yNew}
		_, exists := points[newPoint]
		if !exists {
			points[newPoint] = 0
			continue
		}
		points[newPoint] = 1
	}
	return points
}

func partOne(input []string) int {
	moves := extractData(input)

	pointsVisited := make(map[[2]int]int)

	for _, move := range moves {
		if move[0][0] == move[1][0] || move[0][1] == move[1][1] {
			pointsVisited = updatePoints(pointsVisited, move)
		}
	}

	score := 0
	for _, num := range pointsVisited {
		score += num
	}

	return score
}

func partTwo(input []string) int {
	moves := extractData(input)

	pointsVisited := make(map[[2]int]int)

	for _, move := range moves {
		pointsVisited = updatePoints(pointsVisited, move)
	}

	score := 0
	for _, num := range pointsVisited {
		score += num
	}

	return score
}
