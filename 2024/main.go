package main

import (
	"AdventOfCode2024/day/day_01"
	"AdventOfCode2024/day/day_02"
	"AdventOfCode2024/util"
	"flag"
	"log"
	"strconv"
	"time"
)

var days = map[int]func(string){
	1: day_01.Run,
	2: day_02.Run,
}

func main() {
	_, _, currentDay := time.Now().Date()

	dayFlag := flag.String("day", strconv.Itoa(currentDay), "Day of the month")
	inputPathFlag := flag.String("input", "input.txt", "Path to the input.txt file")
	flag.Parse()

	day, err := strconv.Atoi(*dayFlag)
	if err != nil {
		log.Fatal(err)
	}

	days[day](util.ReadInput(*inputPathFlag))
}
