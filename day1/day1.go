package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

const fileName = "day1/input.txt"

func PartOne() int {
	leftList, rightList, done := getLists()
	if done {
		return 0
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	result := 0
	for i := 0; i < len(leftList); i++ {
		left := leftList[i]
		right := rightList[i]
		if left > right {
			result += left - right
			continue
		}

		result += right - left
	}

	return result
}

func PartTwo() int {
	leftList, rightList, done := getLists()
	if done {
		return 0
	}

	result := 0
	for i := 0; i < len(leftList); i++ {
		left := leftList[i]
		for _, r := range rightList {
			if r == left {
				result += left
			}
		}
	}

	return result
}

func getLists() ([]int, []int, bool) {
	content := readFile()
	if content == "" {
		return nil, nil, true
	}

	lines := strings.Split(content, "\r\n")
	var leftList []int
	var rightList []int
	for _, line := range lines {
		splittedLine := strings.Split(line, "   ")
		leftLine, err := strconv.Atoi(splittedLine[0])
		if err != nil {
			return nil, nil, true
		}
		leftList = append(leftList, leftLine)

		rightLine, err := strconv.Atoi(splittedLine[1])
		if err != nil {
			return nil, nil, true
		}
		rightList = append(rightList, rightLine)
	}

	return leftList, rightList, false
}

func readFile() string {
	result, err := os.ReadFile(fileName)
	if err != nil {
		return ""
	}
	return string(result)
}
