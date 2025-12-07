package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Line []string

type Input struct {
	start int
	grid  []Line
}

func main() {
	input := loadInput("day07/input.txt")
	//part1(input)
	part2(input)
}

func loadInput(filepath string) Input {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input Input
	var first = true

	for scanner.Scan() {
		line := scanner.Text()

		if first {
			input.start = strings.Index(line, "S")
			first = false
		} else {
			input.grid = append(input.grid, strings.Split(line, ""))
		}
	}

	return input
}

func part1(input Input) {
	var result int

	beams := map[int]bool{
		input.start: true,
	}

	for _, line := range input.grid {

		nextBeams := map[int]bool{}

		for b := range beams {
			if line[b] == "." {
				nextBeams[b] = true
			} else if line[b] == "^" {
				result++
				nextBeams[b-1] = true
				nextBeams[b+1] = true
			}
		}

		beams = nextBeams
	}

	fmt.Printf("Part 1 = %d\n", result)
}

var theGrid []Line

func part2(input Input) {
	var result int

	theGrid = input.grid
	result = countTimelines(0, input.start)

	fmt.Printf("Part 2 = %d\n", result)
}

var cache = map[string]int{}

func countTimelines(depth int, beam int) int {
	key := fmt.Sprintf("%d-%d", depth, beam)

	var count int
	var ok bool

	count, ok = cache[key]
	if !ok {
		if depth == len(theGrid)-1 {
			count = 1
		} else if theGrid[depth][beam] == "^" {
			count = countTimelines(depth+1, beam-1) + countTimelines(depth+1, beam+1)
		} else {
			count = countTimelines(depth+1, beam)
		}

		cache[key] = count
	}

	return count
}
