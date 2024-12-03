package day2

import (
	"AdventOfCode-2024/util"
	"strconv"
	"strings"
)

const filename = "day2/input.txt"

func PartOne() int {
	content := util.ReadFile(filename)
	if content == "" {
		return 0
	}

	lines := strings.Split(content, "\r\n")
	result := 0
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		decreasing := false
		for i := range numbers {
			number, _ := strconv.Atoi(numbers[i])
			if i == len(numbers)-1 {
				result++
				break
			}

			nextNumber, _ := strconv.Atoi(numbers[i+1])
			if i == 0 && nextNumber < number {
				decreasing = true
			}

			isWronglyDecreasing := decreasing && nextNumber >= number
			isWronglyIncreasing := !decreasing && nextNumber <= number
			isWronglyDifferent := util.AbsoluteValue(number-nextNumber) > 3
			if isWronglyDecreasing || isWronglyIncreasing || isWronglyDifferent {
				break
			}
		}

	}
	return result
}

func PartTwo() int {
	content := util.ReadFile(filename)
	if content == "" {
		return 0
	}

	lines := strings.Split(content, "\r\n")
	result := 0
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		if isSafe(numbers) {
			result++
		} else {
			for i := range numbers {
				subNumbers := make([]string, len(numbers))
				copy(subNumbers, numbers)
				subNumbers = append(subNumbers[:i], subNumbers[i+1:]...)
				if isSafe(subNumbers) {
					result++
					break
				}
			}
		}

	}
	return result
}

func isSafe(numbers []string) bool {
	decreasing := false
	for i := range numbers {
		number, _ := strconv.Atoi(numbers[i])
		if i == len(numbers)-1 {
			break
		}

		nextNumber, _ := strconv.Atoi(numbers[i+1])
		if i == 0 && nextNumber < number {
			decreasing = true
		}

		isWronglyDecreasing := decreasing && nextNumber >= number
		isWronglyIncreasing := !decreasing && nextNumber <= number
		isWronglyDifferent := util.AbsoluteValue(number-nextNumber) > 3

		if isWronglyDecreasing || isWronglyIncreasing || isWronglyDifferent {
			return false
		}
	}

	return true
}
