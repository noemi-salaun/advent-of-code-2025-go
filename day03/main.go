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
		result += turnOn2Batteries(bank)
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func turnOn2Batteries(bank Bank) int {
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

	for _, bank := range input {
		result += turnOn12Batteries(bank)
	}

	fmt.Printf("Part 2 = %d\n", result)
}

func turnOn12Batteries(bank Bank) int {
	var batteries [12]int
	var remaining int
	var needs int
	length := len(bank)

out:
	for i, number := range bank {
		for j, b := range batteries {
			remaining = length - i
			needs = 12 - j
			if number > b && remaining >= needs {
				batteries[j] = number
				clearSliceOfArray(batteries[j+1:])
				continue out
			}
		}
	}

	var compose string
	for _, b := range batteries {
		compose += strconv.Itoa(b)
	}

	power, err := strconv.Atoi(compose)
	if err != nil {
		log.Fatalf("Cannot convert power from string %s to int", compose)
	}
	return power
}

func clearSliceOfArray(slice []int) {
	for i := range slice {
		slice[i] = 0
	}
}
