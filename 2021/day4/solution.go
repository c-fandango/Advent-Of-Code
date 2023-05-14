// day four solution
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

type cell struct {
	number  int
	crossed bool
}

func main() {

	input := utils.ReadDataStr("../data/data_day_4.txt")

	answerOne := partOne(input)
	answerTwo := partTwo(input)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func partOne(input []string) int {
	// some of this might be better with pointers
	cards, balls := extractGame(input)

	for _, ball := range balls {
		for i, card := range cards {
			cards[i] = crossCard(card, ball)
			if checkBingo(cards[i]) || checkBingo(transpose(cards[i])) {
				return ball * calcScore(cards[i])
			}
		}
	}
	return 0
}

func partTwo(input []string) int {
	cards, balls := extractGame(input)

	for _, ball := range balls {
		var remaining [][][]cell
		for i, card := range cards {
			cards[i] = crossCard(card, ball)
			if !checkBingo(cards[i]) && !checkBingo(transpose(cards[i])) {
				remaining = append(remaining, cards[i])
			}
		}
		if len(remaining) == 0 {
			return ball * calcScore(cards[0])
		}
		cards = remaining
	}
	return 0
}

func extractGame(input []string) ([][][]cell, []int) {

	ballsStr := strings.Split(input[0], ",")
	balls := make([]int, len(ballsStr))

	for i, ball := range ballsStr {
		balls[i], _ = strconv.Atoi(ball)

	}
	tables := input[1:]
	size := len(strings.Fields(tables[0]))

	var bingoCards [][][]cell

	table := make([][]cell, size)
	rowTable := make([]cell, size)

	for i, row := range tables {
		rowTable = toCells(strings.Fields(row))
		table[i%size] = rowTable
		if i%size == size-1 {
			bingoCards = append(bingoCards, table)
			table = make([][]cell, size)
		}
	}

	return bingoCards, balls
}

func toCells(input []string) []cell {
	output := make([]cell, len(input))
	for i, num := range input {
		newCell := cell{}
		newCell.number, _ = strconv.Atoi(num)
		newCell.crossed = false
		output[i] = newCell
	}
	return output
}

func newCard(length int) [][]cell {
	output := make([][]cell, length)
	for i := 0; i < length; i++ {
		output[i] = make([]cell, length)
	}
	return output
}

func transpose(input [][]cell) [][]cell {
	output := newCard(len(input))

	for i, row := range input {
		for j, val := range row {
			output[j][i] = val
		}
	}
	return output
}

func checkBingo(card [][]cell) bool {
	size := len(card)

	for _, row := range card {
		for i, val := range row {
			if !val.crossed {
				break
			}
			if i == size-1 {
				return true
			}
		}
	}
	return false
}

func crossCard(card [][]cell, ball int) [][]cell {

	for i, row := range card {
		for j, val := range row {
			if val.number == ball {
				card[i][j].crossed = true
			}
		}
	}
	return card

}

func calcScore(card [][]cell) int {
	sum := 0

	for _, row := range card {
		for _, val := range row {
			if !val.crossed {
				sum += val.number
			}
		}
	}
	return sum
}
