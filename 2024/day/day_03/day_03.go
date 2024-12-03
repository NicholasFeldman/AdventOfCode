package day_03

import (
	"fmt"
	"regexp"
	"strconv"
)

func Run(input string) {
	fmt.Println("Part 1: ", PartOne(input))
	fmt.Println("Part 2: ", PartTwo(input))
}

func PartOne(input string) string {
	return strconv.Itoa(sumValidMulls(input))
}

func PartTwo(input string) string {
	input = regexp.MustCompile(`(?s)don't\(\).*?(?:do\(\)|$)`).ReplaceAllString(input, "")
	return strconv.Itoa(sumValidMulls(input))
}

func sumValidMulls(input string) int {
	matches := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`).FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	return sum
}
