package day_2

import (
	fileutils "adventOfCode/utils"
	"bufio"
	"fmt"
	"io"
)

var file io.Reader = fileutils.ReadTextFile("data/aoc2_input") // Read local text file

func Parse1() {
	scanner := bufio.NewScanner(file)

	var (
		options = map[string]int{
			"X": 1,
			"Y": 2,
			"Z": 3,
		}
		outcomes = map[string]int{
			// WIN
			"A Y": 6 + options["Y"],
			"B Z": 6 + options["Z"],
			"C X": 6 + options["X"],
			// DRAW
			"A X": 4 + options["X"],
			"B Y": 5 + options["Y"],
			"C Z": 6 + options["Z"],
			// LOSE
			"A Z": 3 + options["Z"],
			"B X": 1 + options["X"],
			"C Y": 2 + options["Y"],
		}
		myScore = 0
	)
	for scanner.Scan() {
		round := scanner.Text()

		if len(round) > 0 {
			myScore += outcomes[round]
		}
	}
	fmt.Println("My Score: ", myScore, "\n")
	return
}
func Parse2() {
	scanner := bufio.NewScanner(file)

	var options = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	var (
		outcomes = map[string]int{
			// WIN
			"A Z": options["Y"] + 3,
			"B Z": options["Z"] + 6,
			"C Z": options["X"] + 6,
			// DRAW
			"A Y": options["X"] + 3,
			"B Y": options["Y"] + 3,
			"C Y": options["Z"] + 3,
			// LOSE
			"A X": options["Z"] + 0,
			"B X": options["X"] + 0,
			"C X": options["Y"] + 0,
		}
		myScore = 0
		lines   = 0
	)
	for scanner.Scan() {
		lines++
		myScore += outcomes[scanner.Text()]
	}

	fmt.Println("Num of lines: ", lines, "\n")
	fmt.Println("My Score: ", myScore, "\n")
	return
}
