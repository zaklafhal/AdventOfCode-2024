package day4

import (
	"AdventOfCode-2024/util"
	"strings"
)

const filename = "day4/input.txt"

func PartOne() int {
	content := util.ReadFile(filename)
	if content == "" {
		return 0
	}
	keyword := "XMAS"
	matrix := buildMatrix(content)
	occurrences := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j+3 < len(matrix[i]) {
				forward := matrix[i][j] + matrix[i][j+1] + matrix[i][j+2] + matrix[i][j+3]
				if forward == keyword {
					occurrences++
				}
			}
			if j-3 >= 0 {
				backward := matrix[i][j] + matrix[i][j-1] + matrix[i][j-2] + matrix[i][j-3]
				if backward == keyword {
					occurrences++
				}
			}
			if i+3 < len(matrix) {
				downward := matrix[i][j] + matrix[i+1][j] + matrix[i+2][j] + matrix[i+3][j]
				if downward == keyword {
					occurrences++
				}
			}

			if i-3 >= 0 {
				upward := matrix[i][j] + matrix[i-1][j] + matrix[i-2][j] + matrix[i-3][j]
				if upward == keyword {
					occurrences++
				}
			}

			if i+3 < len(matrix) && j+3 < len(matrix[i]) {
				diagonalDownwardRight := matrix[i][j] + matrix[i+1][j+1] + matrix[i+2][j+2] + matrix[i+3][j+3]
				if diagonalDownwardRight == keyword {
					occurrences++
				}

			}

			if i-3 >= 0 && j-3 >= 0 {
				diagonalUpwardLeft := matrix[i][j] + matrix[i-1][j-1] + matrix[i-2][j-2] + matrix[i-3][j-3]
				if diagonalUpwardLeft == keyword {
					occurrences++
				}
			}

			if i+3 < len(matrix) && j-3 >= 0 {
				diagonalDownwardLeft := matrix[i][j] + matrix[i+1][j-1] + matrix[i+2][j-2] + matrix[i+3][j-3]
				if diagonalDownwardLeft == keyword {
					occurrences++
				}
			}

			if i-3 >= 0 && j+3 < len(matrix[i]) {
				diagonalUpwardRight := matrix[i][j] + matrix[i-1][j+1] + matrix[i-2][j+2] + matrix[i-3][j+3]
				if diagonalUpwardRight == keyword {
					occurrences++
				}
			}
		}
	}

	return occurrences
}

func PartTwo() int {
	content := util.ReadFile(filename)
	if content == "" {
		return 0
	}
	matrix := buildMatrix(content)
	occurrences := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i+2 < len(matrix) && j+2 < len(matrix[i]) {
				if matrix[i][j] == "M" && matrix[i+1][j+1] == "A" && matrix[i+2][j+2] == "S" && matrix[i+2][j] == "M" && matrix[i+1][j+1] == "A" && matrix[i][j+2] == "S" {
					occurrences++
				}
				if matrix[i][j] == "S" && matrix[i+1][j+1] == "A" && matrix[i+2][j+2] == "M" && matrix[i+2][j] == "S" && matrix[i+1][j+1] == "A" && matrix[i][j+2] == "M" {
					occurrences++
				}
				if matrix[i][j] == "S" && matrix[i+1][j+1] == "A" && matrix[i+2][j+2] == "M" && matrix[i+2][j] == "M" && matrix[i+1][j+1] == "A" && matrix[i][j+2] == "S" {
					occurrences++
				}
				if matrix[i][j] == "M" && matrix[i+1][j+1] == "A" && matrix[i+2][j+2] == "S" && matrix[i+2][j] == "S" && matrix[i+1][j+1] == "A" && matrix[i][j+2] == "M" {
					occurrences++
				}
			}
		}
	}

	return occurrences
}

func buildMatrix(content string) [][]string {
	lines := strings.Split(content, "\r\n")
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = make([]string, len(line))
		for j := 0; j < len(line); j++ {
			matrix[i][j] = string(line[j])
		}
	}

	return matrix
}
