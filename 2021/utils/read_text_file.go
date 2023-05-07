package utils

import (
	"ioutil"
	"strings"
	"runtime"
	"path"
)

func ReadDataStr(path string) []int {

        file_bytes, _ := ioutil.ReadFile(toAbsPath(path))
        input := string(file_bytes)

        input_strings := strings.Split(input, "\n")

        return input_strings
}

func ReadDataInt(path string) []int {

	input_strings := ReadDataStr(path)

        input_ints := make([]int, len(input_strings))

        for i, str := range input_strings {
                input_ints[i], _ = strconv.Atoi(str)
        }

        return input_ints
}

func toAbsPath(rel_path string) string {
        _, filename, _, _ := runtime.Caller(1)
        return path.Join(path.Dir(filename), rel_path)
}

