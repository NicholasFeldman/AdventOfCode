package day_02

import (
	"AdventOfCode2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run(input string) {
	fmt.Println("Part 1: ", PartOne(util.SplitInput(input)))
	fmt.Println("Part 2: ", PartTwo(util.SplitInput(input)))
}

func PartOne(inputLines []string) string {
	reports := createReports(inputLines)

	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

func PartTwo(inputLines []string) string {
	reports := createReports(inputLines)

	safeCount := 0
	for _, report := range reports {
		if isSafeDamper(report) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

func createReports(inputLines []string) [][]int {
	var reports [][]int

	for _, line := range inputLines {
		parts := strings.Fields(line)
		var report []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	return reports
}

func isSafe(report []int) bool {
	increasing := false
	decreasing := false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if util.Abs(diff) > 3 || util.Abs(diff) < 1 {
			return false
		}
		if diff < 0 {
			decreasing = true
		}
		if diff > 0 {
			increasing = true
		}
	}

	return (increasing || decreasing) && !(increasing && decreasing)
}

// Allow one entry to be incorrect, remove it and check again.
func isSafeDamper(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := range report {
		if isSafe(slices.Delete(slices.Clone(report), i, i+1)) {
			return true
		}
	}

	return false
}
