package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Bank []int
type Input []Bank

func main() {
	input := loadInput("day03/input.txt")
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
		var bank []int
		line := scanner.Text()
		for _, letter := range line {
			number, err := strconv.Atoi(string(letter))
			if err != nil {
				log.Fatalf("Cannot convert string %s to int", string(letter))
			}
			bank = append(bank, number)
		}
		input = append(input, bank)
	}

	return input
}

func part1(input Input) {
	var result int

	for _, bank := range input {
		result += turnOnBatteries(bank)
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func turnOnBatteries(bank Bank) int {
	var max1 int
	var max2 int

	length := len(bank)

	for i, number := range bank {
		if number > max1 && i < length-1 {
			max1 = number
			max2 = 0
		} else if number > max2 {
			max2 = number
		}
	}
	compose := fmt.Sprintf("%d%d", max1, max2)
	power, err := strconv.Atoi(compose)
	if err != nil {
		log.Fatalf("Cannot convert power from string %s to int", compose)
	}
	return power
}

func part2(input Input) {
	var result int

	fmt.Printf("Part 2 = %d\n", result)
}
