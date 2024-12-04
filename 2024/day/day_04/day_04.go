package day_04

import (
	"AdventOfCode2024/util"
	"fmt"
	"strconv"
)

func Run(input string) {
	fmt.Println("Part 1: ", PartOne(util.SplitInput(input)))
	fmt.Println("Part 2: ", PartTwo(util.SplitInput(input)))
}

func PartOne(inputLines []string) string {
	return strconv.Itoa(countXMAS(inputLines))
}

func PartTwo(inputLines []string) string {
	return strconv.Itoa(countCrossMas(inputLines))
}

type SearchOffset struct {
	row int
	col int
}

func countXMAS(inputLines []string) int {
	searchOffsets := []SearchOffset{
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 0},
		{-1, 1},
		{-1, -1},
		{0, 1},
		{0, -1},
	}
	count := 0
	for row, line := range inputLines {
		for col := range line {
			for _, offset := range searchOffsets {
				if searchXMAS(inputLines, row, col, offset) {
					count++
				}
			}
		}
	}

	return count
}

// Find a single XMAS in a given direction
func searchXMAS(inputLines []string, row, col int, offset SearchOffset) bool {
	maxRow := len(inputLines)
	maxCol := len(inputLines[row])

	for i := 0; i < len("XMAS"); i++ {
		curRow := row + i*offset.row
		curCol := col + i*offset.col

		// Check boundaries
		if curRow < 0 || curRow >= maxRow || curCol < 0 || curCol >= maxCol {
			return false
		}

		if inputLines[curRow][curCol] != "XMAS"[i] {
			return false
		}
	}

	return true
}

func countCrossMas(inputLines []string) int {
	count := 0
	for row, line := range inputLines {
		for col := range line {
			if searchCross(inputLines, row, col) {
				count++
			}
		}
	}

	return count
}

// Count zero, one, or two XMAS in a cross pattern
func searchCross(inputLines []string, row, col int) bool {
	maxRow := len(inputLines)
	maxCol := len(inputLines[row])

	// Check if we have a starting point
	if inputLines[row][col] != 'A' {
		return false
	}

	// Check boundaries
	if row+1 >= maxRow || col+1 >= maxCol || row-1 < 0 || col-1 < 0 {
		return false
	}

	topLeft := inputLines[row-1][col-1]
	topRight := inputLines[row-1][col+1]
	bottomLeft := inputLines[row+1][col-1]
	bottomRight := inputLines[row+1][col+1]

	// Check down slope
	if !((topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')) {
		return false
	}

	// Check up slope
	if !((bottomLeft == 'M' && topRight == 'S') || (bottomLeft == 'S' && topRight == 'M')) {
		return false
	}

	return true
}
