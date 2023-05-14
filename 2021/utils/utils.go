// Package utils for advent of code solutions.
package utils

import (
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// IntInSlice returns true if an integer is in a slice of integers.
func IntInSlice(inputInt int, inputSlice []int) bool {
	for _, num := range inputSlice {
		if num == inputInt {
			return true
		}

	}
	return false
}

// StringInSlice returns true if an string is in a slice of strings.
func StringInSlice(inputString string, inputSlice []string) bool {
	for _, num := range inputSlice {
		if num == inputString {
			return true
		}

	}
	return false
}

// MinMax function to return the minimum and maximumof a slice of ints.
func MinMax(input []int) (int, int) {
	min := input[0]
	max := input[0]
	for _, num := range input {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	return min, max
}

// CSVToInts converts a comma separated string of numbers to slice of ints.
func CSVToInts(input string) []int {
	inputSplt := strings.Split(input, ",")
	output := make([]int, len(inputSplt))

	for i, elem := range inputSplt {
		output[i], _ = strconv.Atoi(elem)
	}
	return output
}

// Sign gives sign of integer or zero.
func Sign(input int) int {
	if input > 0 {
		return 1
	} else if input < 0 {
		return -1
	}
	return 0
}

// SwapInt for swapping two integers.
func SwapInt(x int, y int) (int, int) {
	return y, x
}

// SwapInt for swapping two strings.
func SwapStr(x string, y string) (string, string) {
	return y, x
}

// ReadDataStr converts text file into array of strings.
func ReadDataStr(path string) []string {

	fileBytes, _ := ioutil.ReadFile(toAbsPath(path))
	input := string(fileBytes)
	var output []string

	inputStrings := strings.Split(input, "\n")

	for _, word := range inputStrings {
		if len(word) > 0 {
			output = append(output, word)

		}
	}

	return output
}

// ReadDataInt converts text file into array of integers.
func ReadDataInt(path string) []int {

	inputStrings := ReadDataStr(path)

	inputInts := make([]int, len(inputStrings))

	for i, str := range inputStrings {
		inputInts[i], _ = strconv.Atoi(str)
	}

	return inputInts
}

func toAbsPath(relPath string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), relPath)
}
