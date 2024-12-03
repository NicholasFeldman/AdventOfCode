package day_03

import (
	"AdventOfCode2024/util"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := util.ReadInput("input.txt")
	solution := util.ReadInput("part_1_solution.txt")

	result := PartOne(input)
	if result != solution {
		t.Errorf("Got %s, expected %s", result, solution)
	}
}

func TestPartTwo(t *testing.T) {
	input := util.ReadInput("input_2.txt")
	solution := util.ReadInput("part_2_solution.txt")

	result := PartTwo(input)
	if result != solution {
		t.Errorf("Got %s, expected %s", result, solution)
	}
}
