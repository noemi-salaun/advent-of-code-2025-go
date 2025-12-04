package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Input [][]string

func main() {
	input := loadInput("day04/input.txt")
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

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
	}

	return input
}

func part1(input Input) {
	var result int
	var width int

	height := len(input)

	for i, line := range input {
		width = len(line)
		for j, slot := range line {
			if slot == "@" && hasFewerThanFourAdjacent(input, height, width, i, j) {
				result++
			}
		}
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func part2(input Input) {
	var result int
	var width int

	height := len(input)

	for {
		removedThisTurn := 0
		for i, line := range input {
			width = len(line)
			for j, slot := range line {
				if slot == "@" && hasFewerThanFourAdjacent(input, height, width, i, j) {
					input[i][j] = "."
					removedThisTurn++
				}
			}
		}
		result += removedThisTurn
		if removedThisTurn == 0 {
			break
		}
	}

	fmt.Printf("Part 2 = %d\n", result)
}

func hasFewerThanFourAdjacent(input Input, height int, width int, i int, j int) bool {
	var count int
	var theRange = [3]int{-1, 0, 1}

	for _, y := range theRange {
		for _, x := range theRange {
			if x == 0 && y == 0 {
				continue
			}
			if i+y < 0 || i+y >= height || j+x < 0 || j+x >= width {
				continue
			}

			if input[i+y][j+x] == "@" {
				count++
				if count >= 4 {
					return false
				}
			}
		}
	}

	return true
}
