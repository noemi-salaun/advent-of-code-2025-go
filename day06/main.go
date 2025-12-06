package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	operator string
	elements []int
}

func (p *Problem) calculate() int {
	var total int
	if p.operator == "*" {
		total = 1
		for _, el := range p.elements {
			total *= el
		}
	} else if p.operator == "+" {
		for _, el := range p.elements {
			total += el
		}
	} else {
		log.Fatalf("Unknown operator %s", p.operator)
	}

	return total
}

type Input struct {
	problems []Problem
}

func (i *Input) calculateGrandTotal() int {
	var total int

	for _, p := range i.problems {
		total += p.calculate()
	}

	return total
}

func main() {
	input := loadInputPart1("day06/input.txt")
	result := input.calculateGrandTotal()
	fmt.Printf("Result = %d\n", result)
}

func loadInputPart1(filepath string) Input {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input Input

	var lines [][]string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Fields(line))
	}

	nbProblems := len(lines[0])
	for i := 0; i < nbProblems; i++ {
		var problem = Problem{}
		for index, line := range lines {
			isLastLine := index == len(lines)-1

			if isLastLine {
				problem.operator = line[i]
			} else {
				num, err := strconv.Atoi(line[i])
				if err != nil {
					log.Fatalf("cannot convert number %s", line[i])
				}
				problem.elements = append(problem.elements, num)
			}
		}
		input.problems = append(input.problems, problem)
	}

	return input
}
