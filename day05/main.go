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
	fresh     []Range
	available []int
}

func main() {
	input := loadInput("day05/input.txt")
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
	var fresh = true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fresh = false
		} else {
			if fresh {
				parts := strings.Split(line, "-")
				start, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("Cannot convert start %s to int", parts[0])
				}
				end, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatalf("Cannot convert end %s to int", parts[1])
				}
				input.fresh = append(input.fresh, Range{start, end})
			} else {
				id, err := strconv.Atoi(line)
				if err != nil {
					log.Fatalf("Cannot convert ID %s to int", line)
				}
				input.available = append(input.available, id)
			}
		}
	}

	return input
}

func part1(input Input) {
	var result int

out:
	for _, id := range input.available {
		for _, rg := range input.fresh {
			if id >= rg.start && id <= rg.end {
				result++
				continue out
			}
		}
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func part2(input Input) {
	var result int

	var ranges = input.fresh

	var safeRanges []Range
	var nextRanges []Range
	var atLeastOneNewMerge bool

	for {
		atLeastOneNewMerge = false

		for {
			if len(ranges) == 0 {
				break
			}

			var current = ranges[0]
			var others = ranges[1:]

			for _, o := range others {
				merge, merged := mergeRanges(current, o)
				if merged {
					current = merge
					atLeastOneNewMerge = true
				} else {
					nextRanges = append(nextRanges, o)
				}
			}

			safeRanges = append(safeRanges, current)
			ranges = nextRanges
			nextRanges = []Range{}
		}

		if !atLeastOneNewMerge {
			break
		}
		ranges = safeRanges
		safeRanges = []Range{}
		nextRanges = []Range{}
	}

	for _, r := range safeRanges {
		result += (r.end - r.start) + 1
	}

	fmt.Printf("Part 2 = %d\n", result)
}

func mergeRanges(r1 Range, r2 Range) (Range, bool) {
	if r1.end < r2.start {
		return r1, false
	}

	if r1.start <= r2.start && r1.end >= r2.start && r1.end <= r2.end {
		return Range{r1.start, r2.end}, true
	}

	if r1.start <= r2.start && r1.end >= r2.end {
		return r1, true
	}

	if r1.start >= r2.start && r1.start <= r2.end && r1.end >= r2.end {
		return Range{r2.start, r1.end}, true
	}

	if r1.start > r2.end {
		return r1, false
	}

	if r1.start >= r2.start && r1.end <= r2.end {
		return r2, true
	}

	panic("should not happened")
}
