package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type Input struct {
	ranges []Range
}

func main() {
	input := loadInput("day02/input.txt")
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
		ranges := strings.Split(line, ",")
		for _, tuple := range ranges {
			parts := strings.Split(tuple, "-")
			if len(parts) != 2 {
				log.Fatalf("Invalid range: %s", tuple)
			}
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalf("Invalid range part: %s", parts[0])
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("Invalid range part: %s", parts[1])
			}
			input.ranges = append(input.ranges, Range{start: start, end: end})
		}
	}

	return input
}

func part1(input Input) {
	var result int

	for _, r := range input.ranges {
		for id := r.start; id <= r.end; id++ {
			if isInvalidIdPart1(id) {
				result += id
			}
		}
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func isInvalidIdPart1(id int) bool {
	s := strconv.Itoa(id)
	l := len(s)
	if l%2 != 0 {
		return false
	}
	return s[:l/2] == s[l/2:]
}

func part2(input Input) {
	var result int

	for _, r := range input.ranges {
		for id := r.start; id <= r.end; id++ {
			if isInvalidIdPart2(id) {
				result += id
			}
		}
	}

	fmt.Printf("Part 2 = %d\n", result)
}

func isInvalidIdPart2(id int) bool {
	s := strconv.Itoa(id)
	l := len(s)
	if l == 1 {
		return false
	}

	half := l / 2

out:
	for i := 1; i <= half; i++ {
		if l%i != 0 {
			continue
		}

		pattern := s[0:i]

		nbParts := l / i
		if nbParts < 2 {
			log.Fatal("should not be less than 2")
		}

		for p := 1; p < nbParts; p++ {
			offset := p * i
			if pattern != s[offset:offset+i] {
				continue out
			}
		}

		return true
	}

	return false
}
