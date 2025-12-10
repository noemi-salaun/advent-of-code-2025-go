package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Input []Coord

func main() {
	input := loadInput("day09/input.txt")
	part1(input)
	//part2(input)
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

		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Cannot convert value %s to int", parts[0])
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Cannot convert value %s to int", parts[1])
		}
		input = append(input, Coord{x, y})
	}

	return input
}

func part1(input Input) {
	var result int

	for i, a := range input {
		for _, b := range input[i+1:] {
			area := calcArea(a, b)
			if area > result {
				result = area
			}
		}
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func part2(input Input) {
	var result int

	fmt.Printf("Part 2 = %d\n", result)
}

func calcArea(a Coord, b Coord) int {
	return (absInt(a.x-b.x) + 1) * (absInt(a.y-b.y) + 1)
}

func absInt(n int) int {
	return int(math.Abs(float64(n)))
}
