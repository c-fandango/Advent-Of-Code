// Package utils for advent of code solutions
package utils

import (
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// ReadDataStr converts text file into array of strings
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

// ReadDataInt converts text file into array of integers
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
