// day eight solution
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c-fandango/advent_of_code/2021/utils"
)

func main() {

	input := utils.ReadDataStr("../data/data_day_8.txt")

	answerOne := partOne(input)
	answerTwo := partTwo(input)

	fmt.Println(answerOne)
	fmt.Println(answerTwo)
}

func partOne(input []string) int {

	_, outputSignals := extractData(input)

	uniqSignalLengths := []int{2, 3, 4, 7}

	count := 0

	for _, line := range outputSignals {
		for _, signal := range line {
			if utils.IntInSlice(len(signal), uniqSignalLengths) {
				count++
			}
		}
	}
	return count

}

func partTwo(input []string) int {

	count := 0
	inputSignals, outputSignals := extractData(input)

	for i, line := range inputSignals {
		encodings := deduceEncoding(line)
		outputDigits := ""

		for _, code := range outputSignals[i] {
			for key, value := range encodings {
				if len(code) == len(key) && charsInCommon(key, code) == len(code) {
					outputDigits += value
				}
			}
		}
		outputInt, _ := strconv.Atoi(outputDigits)
		count += outputInt
	}

	return count
}

func extractData(input []string) ([][]string, [][]string) {
	inputSignals := make([][]string, len(input))
	outputSignals := make([][]string, len(input))

	for i, line := range input {
		lineSplit := strings.Split(line, " | ")
		inputSignals[i] = strings.Fields(lineSplit[0])
		outputSignals[i] = strings.Fields(lineSplit[1])
	}
	return inputSignals, outputSignals
}

func charsInCommon(stringOne string, stringTwo string) int {
	output := 0

	for _, char := range stringOne {
		if strings.Contains(stringTwo, string(char)) {
			output++
		}
	}
	return output

}

func deduceEncoding(codes []string) map[string]string {
	lenToDigit := map[int]string{2: "1", 3: "7", 4: "4", 7: "8"}
	encodings := make(map[string]string)
	revEncodings := make(map[string]string)

	for _, code := range codes {
		digit, exists := lenToDigit[len(code)]
		if exists {
			encodings[code] = digit
			revEncodings[digit] = code
		}
	}

	for _, code := range codes {
		if len(code) == 5 {
			if charsInCommon(code, revEncodings["1"]) == 2 {
				encodings[code] = "3"
			} else if charsInCommon(code, revEncodings["4"]) == 3 {
				encodings[code] = "5"
			} else {
				encodings[code] = "2"
			}
		} else if len(code) == 6 {
			if charsInCommon(code, revEncodings["1"]) == 1 {
				encodings[code] = "6"
			} else if charsInCommon(code, revEncodings["4"]) == 4 {
				encodings[code] = "9"

			} else {
				encodings[code] = "0"
			}
		}
	}
	return encodings
}
