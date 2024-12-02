package day_01

import (
	"AdventOfCode2024/util"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputLines := util.SplitInput(util.ReadInput("input.txt"))
	solution := util.ReadInput("part_1_solution.txt")

	result := PartOne(inputLines)
	if result != solution {
		t.Errorf("Got %s, expected %s", result, solution)
	}
}

func TestPartTwo(t *testing.T) {
	inputLines := util.SplitInput(util.ReadInput("input.txt"))
	solution := util.ReadInput("part_2_solution.txt")

	result := PartTwo(inputLines)
	if result != solution {
		t.Errorf("Got %s, expected %s", result, solution)
	}
}
