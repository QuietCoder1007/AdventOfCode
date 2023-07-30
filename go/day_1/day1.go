package day_1

import (
	fileutils "adventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
)

var file = fileutils.ReadTextFile("data/aoc1_input") // Read local text file

func Parse() {
	scanner := bufio.NewScanner(file)
	// Close f
	//defer f.Close()

	var (
		elves          [250]int
		overallLargest = 0
		localLargest   = 0
		index          = 0
		topThree       = 0
	)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			if localLargest > overallLargest {
				overallLargest = localLargest
			}
			elves[index] = localLargest
			index++
			localLargest = 0
			continue
		}
		var val, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		localLargest += val
	}
	sort.Ints(elves[:])
	for i := 0; i < 3; i++ {
		topThree += elves[(len(elves)-1)-i]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(overallLargest, "\n")
	fmt.Println(topThree)
	return
}
