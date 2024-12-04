package day3

import (
	"AdventOfCode-2024/util"
	"strconv"
	"strings"
)

const filename = "day3/input.txt"

func PartOne() int {
	content := util.ReadFile(filename)
	if content == "" {
		return 0
	}

	lines := strings.Split(content, "mul(")
	result := 0
	for _, line := range lines {
		values := strings.Split(line, ",")
		if len(values) < 2 {
			continue
		}

		firstValue, err := strconv.Atoi(values[0])
		if err != nil {
			continue
		}

		secondPart := strings.Split(values[1], ")")
		secondValue, err := strconv.Atoi(secondPart[0])
		if err != nil {
			continue
		}

		result += firstValue * secondValue
	}

	return result
}

func PartTwo() int {
	content := util.ReadFile(filename)
	if content == "" {
		return 0
	}

	lines := strings.Split(content, "mul(")
	result := 0
	enabled := true
	for _, line := range lines {
		newEnabledValue := enabled
		if strings.Contains(line, "don't()") {
			newEnabledValue = false
		}

		if strings.Contains(line, "do()") {
			newEnabledValue = true
		}

		if !enabled {
			enabled = newEnabledValue
			continue
		}

		enabled = newEnabledValue
		values := strings.Split(line, ",")
		if len(values) < 2 {
			continue
		}

		firstValue, err := strconv.Atoi(values[0])
		if err != nil {
			continue
		}

		secondPart := strings.Split(values[1], ")")
		secondValue, err := strconv.Atoi(secondPart[0])
		if err != nil {
			continue
		}

		result += firstValue * secondValue
	}

	return result
}
