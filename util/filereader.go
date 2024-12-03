package util

import "os"

func ReadFile(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(result)
}

func AbsoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
