package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Move struct {
	dir   string
	value int
}

func main() {
	moves := loadInput("day01/input.txt")
	//part1(moves)
	part2(moves)
}

func loadInput(filepath string) []Move {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var move Move
	var moves []Move

	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Cannot extract value from %s: %s", line, err)
		}
		move = Move{
			dir:   line[0:1],
			value: value,
		}

		moves = append(moves, move)
	}

	return moves
}

func part1(moves []Move) {
	count := 0
	numbers := 100
	value := 50
	var change int
	var move Move

	for _, move = range moves {
		if move.dir == "L" {
			change = -move.value
		} else {
			change = +move.value
		}

		value = (value + change) % numbers
		if value < 0 {
			value += numbers
		}

		if value == 0 {
			count++
		}
	}

	fmt.Printf("Part 1 = %d\n", count)
}

func part2(moves []Move) {
	count := 0
	numbers := 100
	value := 50
	var change int
	var loops int
	var move Move

	for _, move = range moves {
		change = move.value
		loops = change / numbers
		count += loops
		change = change - (loops * numbers)

		if move.dir == "L" {
			change = -change
		}

		if change == 0 {
			continue
		} else if (value != 0) && (change > 0) && (value+change > numbers) {
			count++
		} else if (value != 0) && (change < 0) && (value+change < 0) {
			count++
		}

		value = (value + change) % numbers
		if value < 0 {
			value += numbers
		}

		if value == 0 {
			count++
		}
	}

	fmt.Printf("Part 2 = %d\n", count)
}
