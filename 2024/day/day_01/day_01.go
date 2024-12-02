package day_01

import (
	"AdventOfCode2024/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Run(input string) {
	fmt.Println("Part 1: ", PartOne(util.SplitInput(input)))
	fmt.Println("Part 2: ", PartTwo(util.SplitInput(input)))
}

func PartOne(inputLines []string) string {
	leftList, rightList := splitLists(inputLines)

	sort.Ints(leftList)
	sort.Ints(rightList)

	distance := 0
	for i := range leftList {
		distance += util.Abs(leftList[i] - rightList[i])
	}

	return strconv.Itoa(distance)
}

func PartTwo(inputLines []string) string {
	leftList, rightList := splitLists(inputLines)

	// Count the number of times every number appears in the right list
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	// Multiply every number in the left list by the number of times it appears in the right list
	score := 0
	for _, entry := range leftList {
		score += entry * rightCount[entry]
	}

	return strconv.Itoa(score)
}

// Split our input into two lists, as columns.
// For example:
//
//	1 2
//	3 4
//	5 6
//
// Would return [1, 3, 5] and [2, 4, 6]
func splitLists(inputLines []string) ([]int, []int) {
	var leftList []int
	var rightList []int

	for _, line := range inputLines {
		split := strings.Fields(line)
		leftString, _ := strconv.Atoi(split[0])
		rightString, _ := strconv.Atoi(split[1])
		leftList = append(leftList, leftString)
		rightList = append(rightList, rightString)
	}

	return leftList, rightList
}
